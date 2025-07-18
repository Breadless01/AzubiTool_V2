package web

import (
	"AzubiTool/internal/domain"
	"AzubiTool/internal/ports"
	"AzubiTool/internal/web/views"
	"context"
	"net/http"
	"strconv"
)

type Handler struct {
	AuthService ports.AuthService
	BcsService  ports.BcsService
	DbService   ports.DbService
}

func GetCurrentUser(ctx context.Context) *domain.User {
	user, _ := ctx.Value(userContextKey).(*domain.User)
	return user
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := GetToastRedirect(w, r)
	if r.Method == http.MethodGet {
		views.Login(
			"Login",
			map[string]any{},
			"",
		).Render(ctx, w)
		return
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		user, err := h.AuthService.Authenticate(username, password)

		if err == nil {
			oid := 1234567890 //h.BcsService.GetOid(username, password)
			user.Oid = strconv.Itoa(oid)
			setSession(*user, w, r)
			SetToastRedirect(w, r, "Erfolgreich eingeloggt!", "success")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		views.Login(
			"Login",
			map[string]any{},
			"Ungültiger Benutzername oder Passwort!").Render(r.Context(), w)
		return
	}
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	clearSession(w, r)
	SetToastRedirect(w, r, "Erfolgreich ausgeloggt!", "success")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handler) LandingPage(w http.ResponseWriter, r *http.Request) {
	user, _ := getUser(r)

	views.LandingPage(
		"AzubiTool",
		map[string]any{
			"user": user,
		},
	).Render(r.Context(), w)
}

// func (h *Handler) TransferPage(w http.ResponseWriter, r *http.Request) {
// 	user := GetCurrentUser(r.Context())
// 	ctx := context.WithValue(r.Context(), "user", user)
// 	if user == nil {
// 		SetToastRedirect(w, r, "Bitte zuerst einloggen!", "info")
// 		http.Redirect(w, r, "/login", http.StatusSeeOther)
// 		return
// 	}
// 	views.TransferPage(
// 		"BCS -> Blok",
// 		"").Render(ctx, w)
// }
