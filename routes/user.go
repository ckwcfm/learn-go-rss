package routes

import (
	"github.com/ckwcfm/learn-go/rss/routeHandlers"

	"github.com/go-chi/chi"
)

var UserRouter = chi.NewRouter()

func init() {
	UserRouter.Post("/", routeHandlers.RegisterUser)
	UserRouter.Post("/login", routeHandlers.LoginUser)
}
