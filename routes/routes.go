package routes

import (
	"net/http"
	"urlshort/routes/services"
)

func HandleShortenWithId(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		services.UpdateURL(w, r)
	case http.MethodDelete:
		services.DeleteUrl(w, r)
	case http.MethodGet:
		services.UrlStats(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func HandleShortenRoot(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		services.URLShortener(w, r)
	case http.MethodGet:
		services.RedirectUrl(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

}

func RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/shorten/", HandleShortenWithId)
	mux.HandleFunc("/shorten", HandleShortenRoot)
	return mux
}
