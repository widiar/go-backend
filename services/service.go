package services

import (
	"backendmaw/config"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func ListAndMap[M any, R any](mapper func(M) R) ([]R, error) {
	var model []M
	if err := config.DB.Find(&model).Error; err != nil {
		return nil, err
	}
	data := make([]R, 0, len(model))
	for _, m := range model {
		data = append(data, mapper(m))
	}
	return data, nil
}

func CreateAndValidate[REQ any, M any](req *REQ, column string, getKey func(*REQ) string, mapper func(*REQ, string) M) error {
	var count int64
	var model M
	key := getKey(req)
	if err := config.DB.Model(&model).Where(column, key).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return gorm.ErrDuplicatedKey
	}
	newID := uuid.NewString()
	if err := config.DB.Create(new(mapper(req, newID))).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAndValidate[REQ any, M any](
	req *REQ, columnUnique string, id *string, getKey func(*REQ) string, model M) error {
	if err := config.DB.First(&model, "id = ?", id).Error; err != nil {
		return err
	}
	key := getKey(req)
	var count int64
	if err := config.DB.Model(&model).Where(columnUnique+" = ? AND id <> ?", key, id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return gorm.ErrDuplicatedKey
	}
	if err := config.DB.Model(&model).Updates(req).Error; err != nil {
		return err
	}
	return nil
}
