package apis

import (
	"net/http"

	"github.com/ckwcfm/learn-go/rss/routeHandlers"
)

func UserRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("POST /register", routeHandlers.RegisterUser)
	router.HandleFunc("POST /login", routeHandlers.LoginUser)
	return http.StripPrefix("/users", router)
}
