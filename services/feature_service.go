package services

import (
	"backendmaw/dto"
	"backendmaw/models"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type FeatureService struct {
	DB *gorm.DB
}

func NewFeatureService(db *gorm.DB) *FeatureService {
	return &FeatureService{db}
}

func (s *FeatureService) List() (*dto.ResponseDto, error) {
	data, err := ListAndMap(s.DB, func(m models.Feature) dto.FeatureResponse {
		return dto.FeatureResponse{
			Id:    m.Id,
			Name:  m.Name,
			Label: m.Label,
		}
	})
	if err != nil {
		return new(dto.ErrorResponse()), err
	}
	return new(dto.SuccessResponse(data)), nil
}

func (s *FeatureService) Create(feature *dto.FeatureRequest) (*dto.ResponseDto, error) {
	switch err := CreateAndValidate(s.DB, feature, "name = ?", func(r *dto.FeatureRequest) string { return r.Name }, func(r *dto.FeatureRequest, id string) models.Feature {
		return models.Feature{
			Id:    id,
			Name:  r.Name,
			Label: r.Label,
		}
	}); {
	case err != nil:
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return new(dto.FailedResponse("Data already exists", http.StatusConflict)), err
		}
		return new(dto.ErrorResponse()), err
	}
	return new(dto.SuccessResponse(feature)), nil
}

func (s *FeatureService) Update(id string, feature *dto.FeatureRequest) (*dto.ResponseDto, error) {
	switch err := UpdateAndValidate(s.DB, feature, "name", &id, func(r *dto.FeatureRequest) string { return r.Name }, &models.Feature{}); {
	case err != nil:
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return new(dto.FailedResponse("Data already exists", http.StatusConflict)), err
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			return new(dto.FailedResponse("Data does not exist", http.StatusOK)), err
		}
		return new(dto.ErrorResponse()), err
	}
	return new(dto.SuccessResponse(feature)), nil
}

func (s *FeatureService) Delete(id string) (*dto.ResponseDto, error) {
	var data models.Feature
	if err := s.DB.First(&data, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return new(dto.FailedResponse("Feature not found", http.StatusOK)), nil
		}
		return new(dto.ErrorResponse()), err
	}
	if err := s.DB.Delete(&data).Error; err != nil {
		return new(dto.ErrorResponse()), err
	}
	return new(dto.SuccessResponse("Feature deleted")), nil
}
