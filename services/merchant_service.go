package services

import (
	"backendmaw/config"
	"backendmaw/dto"
	"backendmaw/models"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

func ListMerchantService() (*dto.ResponseDto, error) {
	data, err := ListAndMap(func(m models.Merchant) dto.MerchantResponse {
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

func CreateMerchantService(merchant *dto.MerchantRequest) (*dto.ResponseDto, error) {
	switch err := CreateAndValidate(merchant, "name = ?",
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

func UpdateMerchantService(id string, merchant *dto.MerchantRequest) (*dto.ResponseDto, error) {
	//check id exist or not
	switch err := UpdateAndValidate(merchant, "name", &id, func(r *dto.MerchantRequest) string { return r.Name }, models.Merchant{}); {
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

func DeleteMerchantService(id string) (*dto.ResponseDto, error) {
	if err := config.DB.Delete(&models.Merchant{}, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return new(dto.FailedResponse("Merchant not found", http.StatusOK)), err
		}
		return new(dto.ErrorResponse()), err
	}
	return new(dto.SuccessResponse("Merchant deleted")), nil
}
