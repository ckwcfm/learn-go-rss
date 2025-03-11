package routes

import (
	"log"
	"net/http"

	"github.com/ckwcfm/learn-go/rss/routes/apis"
)

func APIRouter() http.Handler {
	router := http.NewServeMux()
	log.Println("api router")
	router.Handle("/v1/", apis.V1Router())

	router.Handle("/test", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("api test"))
	}))
	return http.StripPrefix("/api", router)
}
