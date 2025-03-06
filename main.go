package main

import (
	"context"
	"fmt"
	"html/template"
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

	router.Mount("/v1", routes.V1Router)

	h1 := func(w http.ResponseWriter, r *http.Request) {
		type PageData struct {
			Title   string
			Message string
		}
		pageData := []PageData{
			{Title: "Hello World", Message: "Welcome to the API"},
			{Title: "Hello World2", Message: "Welcome to the API2"},
		}
		data := map[string][]PageData{
			"data": pageData,
		}
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, data)
	}

	router.Get("/", h1)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Println("Starting server on port", port)
	log.Fatal(srv.ListenAndServe())
}
