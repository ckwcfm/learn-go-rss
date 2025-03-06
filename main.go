package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ckwcfm/learn-go/rss/db"
	"github.com/ckwcfm/learn-go/rss/routes"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	defer func() {
		err := db.DisconnectMongo(context.Background())
		if err != nil {
			log.Fatal("Error disconnecting from MongoDB", err)
		}
	}()

	LoadEnv()

	port := os.Getenv("PORT")
	fmt.Println(port)

	err := db.ConnectToMongo(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handlerError)
	v1Router.Mount("/users", routes.UserRouter)
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Println("Starting server on port", port)
	log.Fatal(srv.ListenAndServe())
}
