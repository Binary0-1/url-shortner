package models

import (
	"gorm.io/gorm"
)


type URL struct {
	gorm.Model
	Url      string    `gorm:"type:varchar(2048);not null"`
	Shortcode string    `gorm:"type:varchar(100);not null;uniqueIndex"`
}