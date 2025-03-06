package middlewares

import (
	"log"
	"net/http"

	"github.com/go-chi/cors"
)

func CORSMiddleware(next http.Handler) http.Handler {
	log.Println("CORSMiddleware")
	return cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	})(next)
}
