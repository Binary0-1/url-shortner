package services

import (
	"encoding/json"
	"math/rand"
	"net/http"
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
	p, ok := validate_url(w, r)
	if !ok {
		return
	}

	err := shorten_url(p)
	if err != nil {
		return
	}
}

func validate_url(w http.ResponseWriter, r *http.Request) (Payload, bool) {
	var p Payload
	if r.Method != "POST" {
		utils.WriteError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return p, false
	}

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		utils.WriteError(w, "Parsing Error ", http.StatusBadRequest)
		return p, false
	}

	if p.URL == "" {
		utils.WriteError(w, "URL not found Please check the key ", http.StatusBadRequest)
		return p, false
	}
	return p, true

}

func shorten_url(p Payload) error {
	url := models.URL{
		Url:       p.URL,
		Shortcode: generateShortCode(),
	}

	db := db.GetDatabaseConnection()

	err := db.Create(&url).Error
	if err != nil {
		return err
	}

	return nil
}

func generateShortCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length := 6
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)

}
