package services

import (
	"net/http"
	"time"
	"urlshort/db"
	"urlshort/models"
	"urlshort/utils"
)

type UrlResponse struct {
	Shortcode   string    `json:"shortCode"`
	Original    string    `json:"url"`
	Id          int       `json:"id"`
	AccessCount int       `json:"accessCount"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func UrlStats(w http.ResponseWriter, r *http.Request) {
	params, err := Get_url_params(r, "/shorten/", 2)
	if err != nil {
		return
	}
	shorturl := params[0]

	var existingUrl models.URL
	database := db.GetDatabaseConnection()
	result := database.Where("shortcode = ?", shorturl).First(&existingUrl)

	if result.Error != nil {
		utils.WriteError(w, "URL unknown ", http.StatusBadRequest)
		return
	}
	response := UrlResponse{
		Shortcode:   existingUrl.Shortcode,
		Original:    existingUrl.Url,
		Id:          int(existingUrl.ID),
		AccessCount: existingUrl.AccessCount,
		CreatedAt:   existingUrl.CreatedAt,
		UpdatedAt:   existingUrl.UpdatedAt,
	}

	utils.JSONResponse(w, response, http.StatusOK)
}
