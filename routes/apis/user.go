package apis

import (
	"net/http"

	"github.com/ckwcfm/learn-go/rss/routeHandlers"
)

// var UserRouter = chi.NewRouter()

// func init() {
// 	UserRouter.Post("/", routeHandlers.RegisterUser)
// 	UserRouter.Post("/login", routeHandlers.LoginUser)

// }

func UserRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/", routeHandlers.RegisterUser)
	router.HandleFunc("/login", routeHandlers.LoginUser)
	return http.StripPrefix("/users", router)
}
