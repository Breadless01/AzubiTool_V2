package web

import (
	"AzubiTool/internal/domain"
	"AzubiTool/internal/ports"
	"AzubiTool/internal/web/views"
	"context"
	"encoding/gob"
	"net/http"
)

func init() {
	gob.Register(&domain.User{})
}

type Handler struct {
	AuthService ports.AuthService
	BcsService  ports.BcsService
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

		user, errUser := h.AuthService.Authenticate(username, password)
		oid, errBcs := h.BcsService.GetOid(username, password)
		if errUser == nil && errBcs == nil {
			sess, _ := SessionStore.Get(r, "session")
			user.Oid = oid
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

func (h *Handler) LandingPage(w http.ResponseWriter, r *http.Request) {
	user := GetCurrentUser(r.Context())
	if user == nil {
		views.LandingPage(
			"AzubiTool",
			false,
			"",
			"",
		).Render(r.Context(), w)
	} else {
		ctx := context.WithValue(r.Context(), "toast", "Dein Toast-Text!")
		views.LandingPage(
			"AzubiTool",
			true,
			user.DisplayName,
			user.Role,
		).Render(ctx, w)
	}
}

func (h *Handler) TransferPage(w http.ResponseWriter, r *http.Request) {
	user := GetCurrentUser(r.Context())
	if user == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	views.TransferPage("BCS -> Blok", true, user.Username, user.Role, user.Oid, "").Render(r.Context(), w)
}
