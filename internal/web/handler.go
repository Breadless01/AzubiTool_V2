package web

import (
	"AzubiTool/internal/web/views"
	"net/http"
)

type Handler struct {
}

func (h *Handler) LandingPage(w http.ResponseWriter, r *http.Request) {
	// Später echte Authentifizierung, jetzt Dummy-User
	loggedIn := true
	userName := "Julius"
	userRole := "Software Entwickler"
	views.LandingPage(
		"AzubiTool – Landingpage",
		loggedIn,
		userName,
		userRole,
	).Render(r.Context(), w)
}
