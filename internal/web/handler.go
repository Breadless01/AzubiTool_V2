package web

import (
	"AzubiTool/internal/domain"
	"AzubiTool/internal/ports"
	"AzubiTool/internal/web/views"
	"context"
	"encoding/gob"
	"net/http"
	"strconv"
)

func init() {
	gob.Register(&domain.User{})
}

type Handler struct {
	AuthService ports.AuthService
	BcsService  ports.BcsService
}

func (h *Handler) LoginPage(w http.ResponseWriter, r *http.Request) {
	ctx := GetToastRedirect(w, r)
	if r.Method == http.MethodGet {
		views.LoginPage(
			"Login",
			"",
		).Render(ctx, w)
		return
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		user, errUser := h.AuthService.Authenticate(username, password)
		oid := 1234567890 //h.BcsService.GetOid(username, password)
		sess, _ := SessionStore.Get(r, "session")
		if errUser == nil {
			user.Oid = strconv.Itoa(oid)
			sess.Values["user"] = user
			sess.Values["loggedIn"] = true
			sess.Save(r, w)
			SetToastRedirect(w, r, "Erfolgreich eingeloggt!", "success")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		sess.Values["loggedIn"] = false
		sess.Save(r, w)

		views.LoginPage(
			"Login",
			"Ungültiger Benutzername oder Passwort!").Render(r.Context(), w)
		return
	}
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	sess, _ := SessionStore.Get(r, "session")
	delete(sess.Values, "user")
	sess.Save(r, w)
	SetToastRedirect(w, r, "Erfolgreich ausgeloggt!", "success")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handler) LandingPage(w http.ResponseWriter, r *http.Request) {
	user := GetCurrentUser(r.Context())
	ctx := context.WithValue(r.Context(), "user", user)

	if user == nil {
		SetToastRedirect(w, r, "Bitte zuerst einloggen!", "info")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	} else {
		// ctx := GetToastRedirect(w, ctx)
		views.LandingPage(
			"AzubiTool",
		).Render(ctx, w)
		return
	}
}

func (h *Handler) TransferPage(w http.ResponseWriter, r *http.Request) {
	user := GetCurrentUser(r.Context())
	ctx := context.WithValue(r.Context(), "user", user)
	if user == nil {
		SetToastRedirect(w, r, "Bitte zuerst einloggen!", "info")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	views.TransferPage(
		"BCS -> Blok",
		"").Render(ctx, w)
}
