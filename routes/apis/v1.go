package apis

import (
	"github.com/go-chi/chi"
)

var V1Router = chi.NewRouter()

func init() {
	// V1Router.Mount("/users", UserRouter)
	// V1Router.Get("/healthz", handlerReadiness)
	// V1Router.Get("/error", handlerError)
}
