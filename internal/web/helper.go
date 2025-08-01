package web

import (
	"context"
	"net/http"
)

const toastContextKey contextKey = "toast"
const toastTypeContextKey contextKey = "toastType"

func SetToast(r *http.Request, message, toastType string) context.Context {
	ctx := context.WithValue(r.Context(), toastContextKey, message)
	ctx = context.WithValue(ctx, toastTypeContextKey, toastType)
	return ctx
}

func SetToastRedirect(w http.ResponseWriter, r *http.Request, message, toastType string) error {
	sess, _ := SessionStore.Get(r, "session")
	sess.Values["toast"] = message
	sess.Values["toastType"] = toastType
	return sess.Save(r, w)
}

func GetToastRedirect(w http.ResponseWriter, r *http.Request) context.Context {
	sess, _ := SessionStore.Get(r, "session")
	var msg, ttype string
	if v, ok := sess.Values["toast"].(string); ok {
		msg = v
		delete(sess.Values, "toast")
	}
	if v, ok := sess.Values["toastType"].(string); ok {
		ttype = v
		delete(sess.Values, "toastType")
	}
	sess.Save(r, w)
	if msg == "" {
		return r.Context()
	}
	ctx := context.WithValue(r.Context(), toastContextKey, msg)
	ctx = context.WithValue(ctx, toastTypeContextKey, ttype)
	return r.WithContext(ctx).Context()
}
