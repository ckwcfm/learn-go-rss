package routes

import (
	"github.com/ckwcfm/learn-go/rss/middlewares"
	"github.com/ckwcfm/learn-go/rss/routes/pages"
	"github.com/ckwcfm/learn-go/rss/routes/pages/auth"
	"github.com/go-chi/chi"
)

var PageRouter = chi.NewRouter()

func init() {
	PageRouter.Get("/", pages.Home)
	PageRouter.With(middlewares.Authorization).Get("/about", pages.About)
	PageRouter.Get("/login", auth.Login)
	PageRouter.Post("/login", auth.LoginHandler)
	PageRouter.Get("/register", auth.Register)
	PageRouter.Post("/register", auth.RegisterHandler)
	PageRouter.Post("/logout", auth.Logout)
}
