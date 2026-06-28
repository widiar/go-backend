package services

import (
	"backendmaw/dto"

	"gorm.io/gorm"
)

type BannerService struct {
	DB *gorm.DB
}

func NewBannerService(db *gorm.DB) *BannerService {
	return &BannerService{db}
}

func (s *BannerService) ListBanner() (*dto.ResponseDto, error) {
	return new(dto.SuccessResponse(nil)), nil
}
