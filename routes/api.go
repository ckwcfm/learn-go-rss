package routes

import (
	"net/http"

	"github.com/ckwcfm/learn-go/rss/routes/apis"
)

func APIRouter() http.Handler {
	router := http.NewServeMux()
	router.Handle("/v1/", apis.V1Router())
	return http.StripPrefix("/api", router)
}
