package routes

import (
	"net/http"

	"github.com/ckwcfm/learn-go/rss/middlewares"
	"github.com/ckwcfm/learn-go/rss/routeHandlers/pages"
	"github.com/ckwcfm/learn-go/rss/routeHandlers/pages/auth"
)

func PageRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/", pages.Home)
	router.Handle("GET /about", middlewares.IsUser(http.HandlerFunc(pages.About)))
	router.Handle("GET /posts", middlewares.IsUser(http.HandlerFunc(pages.Post)))
	router.Handle("GET /books", middlewares.IsUser(http.HandlerFunc(pages.Book)))
	router.HandleFunc("GET /login", auth.Login)
	router.HandleFunc("POST /login", auth.LoginHandler)
	router.HandleFunc("GET /register", auth.Register)
	router.HandleFunc("POST /register", auth.RegisterHandler)
	router.HandleFunc("POST /logout", auth.Logout)
	return router
}
