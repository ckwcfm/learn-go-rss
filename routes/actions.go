package routes

import (
	"github.com/ckwcfm/learn-go/rss/routes/actions/dialogs"
	"github.com/ckwcfm/learn-go/rss/routes/actions/pages"
	"github.com/go-chi/chi"
)

var ActionRouter = chi.NewRouter()

func init() {
	ActionRouter.Get("/pages/home", pages.HomePage)
	ActionRouter.Get("/pages/about", pages.AboutPage)
	ActionRouter.HandleFunc("/dialogs/homeDialog", dialogs.ActionHomeDialog)
}
