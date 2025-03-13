package apis

import (
	"log"
	"net/http"
)

func V1Router() http.Handler {
	router := http.NewServeMux()
	log.Println("v1 router")
	router.HandleFunc("/healthz", handlerReadiness)
	router.HandleFunc("/error", handlerError)
	router.Handle("/users/", UserRouter())
	return http.StripPrefix("/v1", router)
}
