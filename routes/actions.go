package routes

import (
	"github.com/ckwcfm/learn-go/rss/routes/actions/dialogs"

	"github.com/go-chi/chi"
)

var ActionRouter = chi.NewRouter()

func init() {
	ActionRouter.HandleFunc("/dialogs/homeDialog", dialogs.ActionHomeDialog)
}
