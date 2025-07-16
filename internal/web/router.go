package web

import (
	"net/http"
)

func NewRouter(handler *Handler) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", UserSessionMiddleware(http.HandlerFunc(handler.LandingPage)))
	mux.HandleFunc("/login", handler.LoginPage)

	// Static-File-Serving für CSS/JS/Images
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets/static"))))

	return mux
}
