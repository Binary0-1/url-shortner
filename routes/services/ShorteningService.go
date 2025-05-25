package services

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"urlshort/db"
	"urlshort/models"
	"urlshort/utils"
)

// Why not just use simple sequential numbers?
// Predictability & Scraping
// Anyone can guess your URLs just by incrementing numbers:
// https://short.ly/1, https://short.ly/2, https://short.ly/3, etc.
// This makes it easy to scrape all your shortened URLs if someone wants to crawl or abuse your system.

// Privacy
// If your URLs are linked to sensitive or private content, guessable shortcodes expose them easily.

// Brute force / Enumeration attacks
// Malicious actors could try millions of numbers to find active URLs, which could lead to:

// Traffic spikes

// Privacy leaks

// Unwanted data harvesting

// Professionalism & Branding
// Random or nicely encoded shortcodes look more professional and less like a test/demo.

type Payload struct {
	URL string `json:"url"`
}

func URLShortener(w http.ResponseWriter, r *http.Request) {
	p, ok := validateURL(w, r)
	if !ok {
		return
	}

	url, err := shortenURL(p)
	if err != nil {

		utils.WriteError(w, "Failed to shorten URL", http.StatusInternalServerError)
		return
	}

	response := map[string]any{
		"id":        url.ID,
		"shortCode": url.Shortcode,
		"url":       url.Url,
		"createdAt": url.CreatedAt,
		"updatedAt": url.UpdatedAt,
	}

	utils.JSONResponse(w, response, http.StatusOK)
}

func validateURL(w http.ResponseWriter, r *http.Request) (Payload, bool) {
	var p Payload
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		utils.WriteError(w, "Parsing Error ", http.StatusBadRequest)
		return p, false
	}
	return p, true
}

func shortenURL(p Payload) (models.URL, error) {

	shortcode, error := generateUniqueShortCode()
	if error != nil {
		return models.URL{}, error
	}
	url := models.URL{
		Url:       p.URL,
		Shortcode: shortcode,
	}

	database := db.GetDatabaseConnection()
	err := database.Create(&url).Error
	if err != nil {
		return models.URL{}, err
	}

	return url, nil
}

func generateShortCode() (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length := 6
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	for i := range b {
		b[i] = charset[int(b[i])%len(charset)]
	}
	return string(b), nil

}

func generateUniqueShortCode() (string, error) {
	database := db.GetDatabaseConnection()
	const maxAttempts = 5
	for i := 0; i < maxAttempts; i++ {
		code, err := generateShortCode()
		if err != nil {
			return "", err
		}
		var count int64
		database.Model(&models.URL{}).Where("shortcode = ?", code).Count(&count)
		if count == 0 {
			return code, nil
		}
	}
	return "", fmt.Errorf("failed to generate unique shortcode after %d attempts", maxAttempts)
}

func UpdateURL(w http.ResponseWriter, r *http.Request) {
	p, ok := validateURL(w, r)
	if !ok {
		return
	}

	params, err := GetUrlParams(r, "/shorten/", 1)
	if err != nil {
		utils.WriteError(w, "URL unknown ", http.StatusBadRequest)
		return
	}

	shorturl := params[0]
	db := db.GetDatabaseConnection()
	var existingUrl models.URL
	result := db.Where("shortcode = ?", shorturl).First(&existingUrl)
	if result.Error != nil {
		utils.WriteError(w, "URL unknown ", http.StatusBadRequest)
		return
	}

	existingUrl.Url = p.URL

	saveResult := db.Save(&existingUrl)
	if saveResult.Error != nil {
		utils.WriteError(w, "Failed to update URL: "+saveResult.Error.Error(), http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, existingUrl, http.StatusOK)
}

func GetUrlParams(r *http.Request, basePath string, expectedCount int) ([]string, error) {
	path := strings.TrimPrefix(r.URL.Path, basePath)
	path = strings.Trim(path, "/")
	parts := strings.Split(path, "/")
	if path == "" && expectedCount == 0 {
		return nil, nil
	}
	if len(parts) != expectedCount {
		return nil, fmt.Errorf("expected %d path segment(s), got %d", expectedCount, len(parts))
	}
	return parts, nil
}
