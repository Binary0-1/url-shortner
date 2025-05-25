package routes

import (
	"net/http"
	"urlshort/routes/services"
)

func Shortener(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPut {
		services.UpdateURL(w, r)
	}
	if r.Method == http.MethodDelete {
		services.DeleteUrl(w, r)
	}

	if r.Method == http.MethodGet {
		services.UrlStats(w, r)
	}
}

func Shorten_URl(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		services.URLShortener(w, r)
	}
}

func RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/shorten/", Shortener)
	mux.HandleFunc("/shorten", Shorten_URl)
	return mux
}
