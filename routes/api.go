package routes

import (
	"log"
	"net/http"
)

func APIRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/v1", func(w http.ResponseWriter, r *http.Request) {
		log.Println("V1 bbb")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("V1 bbb"))
	})

	return http.StripPrefix("/api", router)
}
