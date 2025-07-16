// internal/web/middleware.go
package web

import (
	"context"
	"net/http"

	"AzubiTool/internal/domain"
)

type ctxKey string

const userCtxKey ctxKey = "currentUser"

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
		if u, ok := sess.Values["user"].(*domain.User); ok {
			user = u
		}
		ctx := SetCurrentUser(r.Context(), user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
