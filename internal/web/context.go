// internal/web/middleware.go
package web

import (
	"context"
	"net/http"

	"AzubiTool/internal/domain"
)

type ctxKey string

const userCtxKey ctxKey = "user"

func SetCurrentUser(ctx context.Context, user *domain.User) context.Context {
	return context.WithValue(ctx, userCtxKey, user)
}

func GetCurrentUser(ctx context.Context) *domain.User {
	val := ctx.Value(userCtxKey)
	if user, ok := val.(*domain.User); ok {
		return user
	}
	return nil
}

func UserSessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess, _ := SessionStore.Get(r, "session")
		var user *domain.User
		var loggedIn bool
		if u, ok := sess.Values["user"].(*domain.User); ok {
			user = u
			loggedIn = sess.Values["loggedIn"].(bool)
		}
		ctx := SetCurrentUser(r.Context(), user)
		ctx = context.WithValue(ctx, "loggedIn", loggedIn)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
