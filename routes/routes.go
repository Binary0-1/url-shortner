package routes

import (
	"net/http"
	"urlshort/routes/services"
)

func Shortener(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		services.URLShortener(w, r)
	}
	if r.Method == http.MethodPut {
		services.UpdateURL(w, r)
	}
}

func RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/shorten/", Shortener)
	return mux
}
