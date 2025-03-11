package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ckwcfm/learn-go/rss/db"
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
	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Root")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Root"))
	}))
	router.Handle("/api/", routes.APIRouter())

	fmt.Println("Starting server on port", port)
	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	log.Fatal(server.ListenAndServe())

}

func APIMux() http.Handler {
	router := http.NewServeMux()
	router.Handle("/v1", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("V1")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("V1"))
	}))
	return http.StripPrefix("/bbb", router)
}
