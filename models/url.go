package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type URL struct {
	BaseModel
	Url         string `gorm:"type:varchar(2048);not null" json:"url"`
	Shortcode   string `gorm:"type:varchar(100);not null;uniqueIndex" json:"shortcode"`
	AccessCount int    `gorm:"type:int;default:0" json:"access_count"`
}
