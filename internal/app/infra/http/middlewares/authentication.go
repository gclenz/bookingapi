package middlewares

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gclenz/tinybookingapi/internal/app/infra/http/utils"
)

type ContextKey string

const ContextUserID ContextKey = "userID"

type Authentication struct {
	jh utils.IJWTHandler
}

func (am *Authentication) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		at, err := r.Cookie("access_token")
		if err != nil {
			slog.Error("Middleware(Authentication)", "info", "error reading cookie access_token", "error", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		userID, err := am.jh.VerifyJWT(strings.Split(at.Value, "Bearer ")[1])
		if err != nil {
			slog.Error("Middleware(Authentication)", "info", "error verifying access_token", "error", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), ContextUserID, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func NewAuthentication(jh utils.IJWTHandler) *Authentication {
	return &Authentication{
		jh: jh,
	}
}
