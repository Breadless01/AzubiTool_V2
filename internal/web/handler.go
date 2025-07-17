package web

import (
	"AzubiTool/internal/domain"
	"AzubiTool/internal/ports"
	"AzubiTool/internal/web/views"
	"encoding/gob"
	"net/http"
)

func init() {
	gob.Register(&domain.User{})
}

type Handler struct {
	AuthService ports.AuthService
}

func (h *Handler) LandingPage(w http.ResponseWriter, r *http.Request) {
	// Später echte Authentifizierung, jetzt Dummy-User
	user := GetCurrentUser(r.Context())
	if user == nil {
		views.LandingPage(
			"AzubiTool",
			false,
			"",
			"",
		).Render(r.Context(), w)
	} else {
		views.LandingPage(
			"AzubiTool",
			true,
			user.DisplayName,
			user.Role,
		).Render(r.Context(), w)
	}
}

func (h *Handler) LoginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		views.LoginPage(
			"Login",
			false,
			"",
			"",
			"",
		).Render(r.Context(), w)
		return
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		user, err := h.AuthService.Authenticate(username, password)
		if err == nil {
			sess, _ := SessionStore.Get(r, "session")
			sess.Values["user"] = user
			sess.Save(r, w)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		views.LoginPage(
			"Login",
			false,
			"",
			"",
			"Ungültiger Benutzername oder Passwort!").Render(r.Context(), w)
		return
	}
}
