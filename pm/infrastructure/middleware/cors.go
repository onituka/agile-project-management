package middleware

import (
	"net/http"

	"github.com/rs/cors"

	"github.com/onituka/agile-project-management/project-management/config"
)

func CorsMiddlewareFunc() func(http.Handler) http.Handler {
	corsWrapper := cors.New(cors.Options{
		AllowedOrigins: config.Env.Cors.AllowedOrigins,
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders: []string{"Origin", "Content-Type", "Accept", "Accept-Language"},
	})

	return corsWrapper.Handler
}
