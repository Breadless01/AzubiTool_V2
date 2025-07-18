package web

import (
	"AzubiTool/internal/domain"
	"context"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

type contextKey string

const userContextKey contextKey = "user"

var hashKey = securecookie.GenerateRandomKey(32)
var store = sessions.NewCookieStore(hashKey)

func setSession(user domain.User, w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "azubi_session")
	if err != nil {
		print("Error getting session:", err)
	}
	session.Values["user"] = user
	session.Save(r, w)
}

func getUser(r *http.Request) (user domain.User, ok bool) {
	session, err := store.Get(r, "azubi_session")
	if err != nil {
		print("Error getting session:", err)
	}
	u, found := session.Values["user"]
	println(found)
	if found {
		user, ok = u.(domain.User)
		println("User found in session:", user.DisplayName)
	}
	return
}

func clearSession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "azubi_session")
	session.Options.MaxAge = -1 // Session invalidieren
	session.Save(r, w)
}

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := getUser(r)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		ctx := context.WithValue(r.Context(), userContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
