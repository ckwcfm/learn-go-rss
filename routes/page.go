package routes

import (
	"net/http"

	"github.com/ckwcfm/learn-go/rss/middlewares"
	"github.com/ckwcfm/learn-go/rss/routes/pages"
	"github.com/ckwcfm/learn-go/rss/routes/pages/auth"
)

func PageRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/", pages.Home)
	router.Handle("/about", middlewares.Authorization(http.HandlerFunc(pages.About)))
	router.HandleFunc("GET /login", auth.Login)
	router.HandleFunc("POST /login", auth.LoginHandler)
	router.HandleFunc("GET /register", auth.Register)
	router.HandleFunc("POST /register", auth.RegisterHandler)
	router.HandleFunc("POST /logout", auth.Logout)
	return router
}
