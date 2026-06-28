package services

import (
	"backendmaw/dto"
	"backendmaw/models"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type MerchantService struct {
	DB *gorm.DB
}

func NewMerchantService(db *gorm.DB) *MerchantService {
	return &MerchantService{DB: db}
}

func (s *MerchantService) List() (*dto.ResponseDto, error) {
	data, err := ListAndMap(s.DB, func(m models.Merchant) dto.MerchantResponse {
		return dto.MerchantResponse{
			Id:          m.Id,
			Name:        m.Name,
			Description: m.Description,
		}
	})
	if err != nil {
		return new(dto.ErrorResponse()), err
	}
	return new(dto.SuccessResponse(data)), nil
}

func (s *MerchantService) Create(merchant *dto.MerchantRequest) (*dto.ResponseDto, error) {
	switch err := CreateAndValidate(s.DB, merchant, "name = ?",
		func(r *dto.MerchantRequest) string { return r.Name },
		func(r *dto.MerchantRequest, id string) models.Merchant {
			return models.Merchant{
				Id:          id,
				Name:        r.Name,
				Description: r.Description,
			}
		}); {
	case err != nil:
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return new(dto.FailedResponse("Data already exists", http.StatusConflict)), err
		}
		return new(dto.ErrorResponse()), err
	}
	return new(dto.SuccessResponse(merchant)), nil
}

func (s *MerchantService) Update(id string, merchant *dto.MerchantRequest) (*dto.ResponseDto, error) {
	//check id exist or not
	switch err := UpdateAndValidate(s.DB, merchant, "name", &id, func(r *dto.MerchantRequest) string { return r.Name }, models.Merchant{}); {
	case err != nil:
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return new(dto.FailedResponse("Merchant not found", http.StatusOK)), err
		} else if errors.Is(err, gorm.ErrDuplicatedKey) {
			return new(dto.FailedResponse("Data already exists", http.StatusConflict)), err
		}
		return new(dto.ErrorResponse()), err
	}
	return new(dto.SuccessResponse(merchant)), nil
}

func (s *MerchantService) Delete(id string) (*dto.ResponseDto, error) {
	if err := s.DB.Delete(&models.Merchant{}, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return new(dto.FailedResponse("Merchant not found", http.StatusOK)), err
		}
		return new(dto.ErrorResponse()), err
	}
	return new(dto.SuccessResponse("Merchant deleted")), nil
}

func (s *MerchantService) RelateFeature(request *dto.MerchantFeatureRequest) (*dto.ResponseDto, error) {
	//extract merchant and feature id
	merchantIdMap := make(map[string]bool)
	featureIdMap := make(map[string]bool)
	var payloadMerchantId []string
	var payloadFeatureId []string
	for _, item := range request.Items {
		if !merchantIdMap[item.MerchantId] {
			merchantIdMap[item.MerchantId] = true
			payloadMerchantId = append(payloadMerchantId, item.MerchantId)
		}
		for _, featureId := range item.FeatureId {
			if !featureIdMap[featureId] {
				featureIdMap[featureId] = true
				payloadFeatureId = append(payloadFeatureId, featureId)
			}
		}
	}

	// validate merchant id and feature id
	var merchantValidCount int64
	if err := s.DB.Model(&models.Merchant{}).Where("id IN ?", payloadMerchantId).Count(&merchantValidCount).Error; err != nil {
		return new(dto.ErrorResponse()), err
	}
	if int(merchantValidCount) != len(payloadMerchantId) {
		return new(dto.FailedResponse("Merchant ID not valid", http.StatusBadRequest)), nil
	}
	if len(payloadFeatureId) > 0 {
		var featureValidCount int64
		if err := s.DB.Model(&models.Feature{}).Where("id IN ?", payloadFeatureId).Count(&featureValidCount).Error; err != nil {
			return new(dto.ErrorResponse()), err
		}
		if int(featureValidCount) != len(payloadFeatureId) {
			return new(dto.FailedResponse("Feature ID not valid", http.StatusBadRequest)), nil
		}
	}

	//put on db
	err := s.DB.Transaction(func(tx *gorm.DB) error {
		for _, update := range request.Items {
			var newFeatures []*models.Feature
			for _, featureId := range update.FeatureId {
				newFeatures = append(newFeatures, &models.Feature{Id: featureId})
			}
			merchant := &models.Merchant{Id: update.MerchantId}
			if err := tx.Model(&merchant).Association("Features").Replace(newFeatures); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return new(dto.ErrorResponse()), err
	}
	return new(dto.SuccessResponse(request)), nil
}

func (s *MerchantService) ListMerchantFeature() (*dto.ResponseDto, error) {
	var merchants []models.Merchant
	if err := s.DB.Model(&models.Merchant{}).Preload("Features").Find(&merchants).Error; err != nil {
		return new(dto.ErrorResponse()), err
	}
	response := make([]dto.MerchantFeatureResponse, 0, len(merchants))
	for _, merchant := range merchants {
		response = append(response, dto.MerchantFeatureResponse{
			MerchantResponse: dto.MerchantResponse{
				Id:          merchant.Id,
				Name:        merchant.Name,
				Description: merchant.Description,
			},
			Features: mapFeatureResponse(merchant.Features),
		})
	}

	return new(dto.SuccessResponse(response)), nil
}

func mapFeatureResponse(features []*models.Feature) []dto.FeatureResponse {
	res := make([]dto.FeatureResponse, 0, len(features))
	for _, feature := range features {
		res = append(res, dto.FeatureResponse{
			Id:    feature.Id,
			Name:  feature.Name,
			Label: feature.Label,
		})
	}
	return res
}
