package web

import (
	"net/http"
)

func NewRouter(handler *Handler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", http.HandlerFunc(handler.Login))
	mux.Handle("/", RequireAuth(http.HandlerFunc(handler.LandingPage)))
	mux.Handle("/logout", RequireAuth(http.HandlerFunc(handler.Logout)))
	// mux.Handle("/transfer", UserSessionMiddleware(http.HandlerFunc(handler.TransferPage)))

	// Static-File-Serving für CSS/JS/Images
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets/static"))))

	return mux
}
