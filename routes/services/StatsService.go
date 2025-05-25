package services

import (
	"time"
)

type UrlResponse struct {
	Shortcode   string    `json:"shortCode"`
	Original    string    `json:"url"`
	Id          int       `json:"id"`
	AccessCount *int      `json:"accessCount,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
