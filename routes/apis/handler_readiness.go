package apis

import (
	"net/http"

	"github.com/ckwcfm/learn-go/rss/utils"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 200, struct{}{})
}
