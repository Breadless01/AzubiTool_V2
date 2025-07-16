package web

import (
	"net/http"
)

// NewRouter baut den HTTP-Router und verkabelt die Handler.
func NewRouter(handler *Handler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.LandingPage)

	// Static-File-Serving für CSS/JS/Images
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets/static"))))

	return mux
}
