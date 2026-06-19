package models

import (
	"time"

	"gorm.io/gorm"
)

type Merchant struct {
	gorm.Model
	Id          string `gorm:"primary_key"`
	Name        string `gorm:"unique;not null"`
	Description string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Features    []*Feature     `gorm:"many2many:merchant_features;"`
}
