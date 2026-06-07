package services

import "backendmaw/dto"

func ListBannersService() (*dto.ResponseDto, error) {
	return new(dto.SuccessResponse(nil)), nil
}
