package routes

import (
	"net/http"

	"github.com/ckwcfm/learn-go/rss/middlewares"
	"github.com/ckwcfm/learn-go/rss/routes/pages"
	"github.com/ckwcfm/learn-go/rss/routes/pages/auth"
)

// var PageRouter = chi.NewRouter()

// func init() {
// 	PageRouter.Get("/", pages.Home)
// 	PageRouter.With(middlewares.Authorization).Get("/about", pages.About)
// 	PageRouter.Get("/login", auth.Login)
// 	PageRouter.Post("/login", auth.LoginHandler)
// 	PageRouter.Get("/register", auth.Register)
// 	PageRouter.Post("/register", auth.RegisterHandler)
// 	PageRouter.Post("/logout", auth.Logout)
// }

func PageRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/", pages.Home)
	router.Handle("/about", middlewares.Authorization(http.HandlerFunc(pages.About)))
	router.HandleFunc("/login", auth.Login)
	router.HandleFunc("/register", auth.Register)
	router.HandleFunc("/logout", auth.Logout)
	return router
}
