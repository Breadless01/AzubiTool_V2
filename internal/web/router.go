package web

import (
	"net/http"
)

func NewRouter(handler *Handler) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/login", UserSessionMiddleware(http.HandlerFunc(handler.LoginPage)))
	mux.Handle("/logout", UserSessionMiddleware(http.HandlerFunc(handler.Logout)))
	mux.Handle("/", UserSessionMiddleware(http.HandlerFunc(handler.LandingPage)))
	mux.Handle("/transfer", UserSessionMiddleware(http.HandlerFunc(handler.TransferPage)))

	// Static-File-Serving für CSS/JS/Images
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets/static"))))

	return mux
}
