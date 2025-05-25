package routes

import (
	"fmt"
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
	fmt.Print("yahan tk to aya req")
	if r.Method == http.MethodDelete {
		services.DeleteUrl(w, r)
	}
}

func RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/shorten/", Shortener)
	return mux
}
