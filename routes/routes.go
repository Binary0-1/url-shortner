package routes

import (
	"net/http"
	"urlshort/routes/services"
)

func Shortener(w http.ResponseWriter, r *http.Request) {
	services.URLShortener(w, r);
}


func RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/shorten",Shortener);
	return mux
}