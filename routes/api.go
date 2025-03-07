package routes

import (
	"github.com/ckwcfm/learn-go/rss/routes/apis"

	"github.com/go-chi/chi"
)

var APIRouter = chi.NewRouter()

func init() {
	APIRouter.Mount("/v1", apis.V1Router)
}
