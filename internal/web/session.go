package web

import (
	"os"

	"github.com/gorilla/sessions"
)

var SessionStore *sessions.CookieStore

func InitSessionStore() {
	key := os.Getenv("SESSION_KEY")
	if key == "" {
		key = "super-unsicheres-dev-secret!!!bitte-in-prod-setzen"
	}
	SessionStore = sessions.NewCookieStore([]byte(key))
	SessionStore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 12,
		HttpOnly: true,
		SameSite: 3,
		// Secure: true, // in Prod setzen, wenn https!
	}
}
