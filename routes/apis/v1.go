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
	router.Handle("/test", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("api v1 test"))
	}))
	router.Handle("/users/", UserRouter())
	return http.StripPrefix("/v1", router)
}
