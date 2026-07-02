package models

import "gorm.io/gorm"

type ConfigCalendar struct {
	gorm.Model
	Type  string `gorm:"not null"`
	Value string `gorm:"not null"`
}
