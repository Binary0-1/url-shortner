package services

import (
	"net/http"
	"urlshort/db"
	"urlshort/models"
	"urlshort/utils"
)

func DeleteUrl(w http.ResponseWriter, r *http.Request) {
	params, ok := Get_url_params(r, "/shorten/", 1)

	if ok != nil {
		return
	}

	shortUrl := params[0]
	database := db.GetDatabaseConnection()

	var existingUrl models.URL
	result := database.Where("shortcode = ?", shortUrl).First(&existingUrl)

	if result.Error != nil {
		utils.WriteError(w, "URL unknown ", http.StatusBadRequest)
		return
	}

	database.Delete(&existingUrl)

	w.WriteHeader(http.StatusNoContent)

}
