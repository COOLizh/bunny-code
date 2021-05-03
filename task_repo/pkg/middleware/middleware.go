// Package middleware implements middlewares for routers
package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/rs/cors"

	"gitlab.com/greenteam1/task_repo/pkg/config"
	"gitlab.com/greenteam1/task_repo/pkg/models"
	"gitlab.com/greenteam1/task_repo/pkg/token"
)

// AuthValidateTokenFunc ...
func AuthValidateTokenFunc(conf *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenBearer := r.Header.Get("Authorization")
			tokenAuth := strings.TrimPrefix(tokenBearer, "Bearer ")
			id, err := token.Parse(tokenAuth, conf.JwtSalt)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			if tokenAuth == "" {
				w.WriteHeader(http.StatusForbidden)
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), models.ContextTokenID, id)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// SetCORSHeaders sets the necessary headers to match the CORS policy
func SetCORSHeaders(conf *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		c := cors.New(cors.Options{
			AllowedOrigins:   []string{conf.ClientAddr},
			AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodOptions},
			AllowedHeaders:   []string{"Content-Type", "Authorization"},
			AllowCredentials: true,
		})
		return c.Handler(next)
	}
}
