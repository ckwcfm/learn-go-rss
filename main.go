package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ckwcfm/learn-go/rss/db"
	"github.com/ckwcfm/learn-go/rss/middlewares"
	"github.com/ckwcfm/learn-go/rss/routes"

	"github.com/go-chi/chi"
)

func main() {
	defer func() {
		err := db.DisconnectMongo(context.Background())
		if err != nil {
			log.Fatal("Error disconnecting from MongoDB", err)
		}
	}()

	loadEnv()

	port := os.Getenv("PORT")
	fmt.Println(port)

	err := db.ConnectToMongo(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()

	router.Use(middlewares.CORSMiddleware)
	router.Use(middlewares.Logger)

	router.Mount("/api", routes.APIRouter)
	router.Mount("/actions", routes.ActionRouter)
	router.Mount("/", routes.PageRouter)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Println("Starting server on port", port)
	log.Fatal(srv.ListenAndServe())
}
