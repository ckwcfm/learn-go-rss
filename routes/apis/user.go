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
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("user test router"))
	})
	router.HandleFunc("POST /register", routeHandlers.RegisterUser)
	router.HandleFunc("POST /login", routeHandlers.LoginUser)
	return http.StripPrefix("/users", router)
}
