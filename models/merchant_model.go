package models

import (
	"time"

	"gorm.io/gorm"
)

type Merchant struct {
	Id          string `gorm:"primary_key"`
	Name        string `gorm:"unique;not null"`
	Description string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
