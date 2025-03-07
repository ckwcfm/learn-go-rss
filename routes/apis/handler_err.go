package apis

import (
	"net/http"

	"github.com/ckwcfm/learn-go/rss/utils"
)

func handlerError(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, 400, "Something went wrong")
}
