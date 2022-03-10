package Middleware

import (
	"context"
	"github.com/go-chi/cors"
	"net/http"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		corsOptionsHandler := cors.Handler(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			AllowCredentials: true,
		})
		ctx := context.WithValue(r.Context(), "cors", corsOptionsHandler)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
