package middleware

import (
    "github.com/rs/cors"
    "net/http"
)

// SetupCORS configures CORS middleware for the application.
func SetupCORS(next http.Handler) http.Handler {
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
        AllowedHeaders:   []string{"Content-Type", "Authorization"},
        AllowCredentials: true,
    })

    return c.Handler(next)
}
