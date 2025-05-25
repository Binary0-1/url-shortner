package services

import (
	"net/http"
	"urlshort/db"
	"urlshort/models"
	"urlshort/utils"
)

func RedirectUrl(w http.ResponseWriter, r *http.Request) {
	params, err := GetUrlParams(r, "", 1)
	if err != nil {
		return
	}
	shortcode := params[0]
	database := db.GetDatabaseConnection()
	var url models.URL

	result := database.Select("url").Where("shortcode = ?", shortcode).First(&url)
	if result.Error != nil {
		utils.WriteError(w, "Shortcode not found", http.StatusNotFound)
		return
	}
	url.AccessCount++
	database.Save(&url)
	http.Redirect(w, r, url.Url, http.StatusFound)

}
