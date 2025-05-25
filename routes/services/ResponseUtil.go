package services

import (
	"net/http"
	"strings"
	"urlshort/db"
	"urlshort/models"
	"urlshort/utils"
)

func DecideHander(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimPrefix(r.URL.Path, "/shorten/")
	parts := strings.Split(path, "/")

	if len(parts) == 1 {
		GetUrlInfo(w, r)
	} else if len(parts) == 2 && parts[1] == "stats" {
		UrlStats(w, r)
	} else {
		http.NotFound(w, r)
	}

}

func GetUrlInfo(w http.ResponseWriter, r *http.Request) {
	params, err := GetUrlParams(r, "/shorten/", 1)
	if err != nil {
		utils.WriteError(w, "URL unknown", http.StatusBadRequest)
		return
	}
	url, err := fetchURLByShortcode(params[0])
	if err != nil {
		utils.WriteError(w, "URL unknown", http.StatusBadRequest)
		return
	}
	response := UrlResponse{
		Shortcode: url.Shortcode,
		Original:  url.Url,
		Id:        int(url.ID),
		CreatedAt: url.CreatedAt,
		UpdatedAt: url.UpdatedAt,
	}
	utils.JSONResponse(w, response, http.StatusOK)
}

func UrlStats(w http.ResponseWriter, r *http.Request) {
	params, err := GetUrlParams(r, "/shorten/", 2)
	if err != nil {
		utils.WriteError(w, "URL unknown", http.StatusBadRequest)
		return
	}
	url, err := fetchURLByShortcode(params[0])
	if err != nil {
		utils.WriteError(w, "URL unknown", http.StatusBadRequest)
		return
	}
	response := UrlResponse{
		Shortcode:   url.Shortcode,
		Original:    url.Url,
		Id:          int(url.ID),
		AccessCount: &url.AccessCount,
		CreatedAt:   url.CreatedAt,
		UpdatedAt:   url.UpdatedAt,
	}
	utils.JSONResponse(w, response, http.StatusOK)
}

func fetchURLByShortcode(shortcode string) (models.URL, error) {
	var url models.URL
	db := db.GetDatabaseConnection()
	result := db.Where("shortcode = ?", shortcode).First(&url)
	if result.Error != nil {
		return models.URL{}, result.Error
	}
	return url, nil
}
