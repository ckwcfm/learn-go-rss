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

	router := http.NewServeMux()
	router.Handle("/", routes.PageRouter())
	router.Handle("/actions/", routes.ActionRouter())
	router.Handle("/api/", routes.APIRouter())

	fmt.Println("Starting server on port", port)
	chain := middlewares.Chain(
		middlewares.Logger,
		middlewares.CORSMiddleware,
	)
	server := &http.Server{
		Handler: chain(router),
		Addr:    ":" + port,
	}
	log.Fatal(server.ListenAndServe())

}
