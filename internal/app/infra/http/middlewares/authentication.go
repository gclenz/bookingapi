package middlewares

import (
	"context"
	"log/slog"
	"net/http"
)

type ContextKey string

const ContextUserID ContextKey = "userID"

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		at, err := r.Cookie("access_token")
		if err != nil {
			slog.Error("error reading cookie access_token", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), ContextUserID, at.Value)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
