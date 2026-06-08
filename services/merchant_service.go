package services

import (
	"backendmaw/config"
	"backendmaw/dto"
	"backendmaw/models"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func ListMerchantService() (*dto.ResponseDto, error) {
	var merchants []models.Merchant
	if err := config.DB.Find(&merchants).Error; err != nil {
		return new(dto.ErrorResponse()), err
	}
	data := make([]dto.MerchantResponse, 0, len(merchants))
	for _, m := range merchants {
		data = append(data, dto.MerchantResponse{
			Id:          m.Id,
			Name:        m.Name,
			Description: m.Description,
		})
	}
	return new(dto.SuccessResponse(data)), nil
}

func CreateMerchantService(merchant *dto.MerchantRequest) (*dto.ResponseDto, error) {
	var count int64
	if err := config.DB.Find(&models.Merchant{}, "name = ?", merchant.Name).Count(&count).Error; err != nil {
		return new(dto.ErrorResponse()), err
	}
	if count > 0 {
		return new(dto.FailedResponse("Name already exists", http.StatusConflict)), nil
	}
	data := models.Merchant{
		Id:          uuid.NewString(),
		Name:        merchant.Name,
		Description: merchant.Description,
	}
	if err := config.DB.Create(&data).Error; err != nil {
		return new(dto.ErrorResponse()), err
	}
	return new(dto.SuccessResponse(merchant)), nil
}

func UpdateMerchantService(id string, merchant *dto.MerchantRequest) (*dto.ResponseDto, error) {
	//check id exist or not
	var data models.Merchant
	if err := config.DB.First(&data, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return new(dto.FailedResponse("Merchant not found", http.StatusOK)), nil
		}
		return new(dto.ErrorResponse()), err
	}
	// check name unique
	var count int64
	if err := config.DB.Find(&models.Merchant{}, "name = ? AND id <> ?", merchant.Name, id).Count(&count).Error; err != nil {
		return new(dto.ErrorResponse()), err
	}
	if count > 0 {
		return new(dto.FailedResponse("Name already exists", http.StatusConflict)), nil
	}
	if err := config.DB.Model(&data).Updates(merchant).Error; err != nil {
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
