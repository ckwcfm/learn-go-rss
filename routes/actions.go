package routes

import (
	"net/http"

	"github.com/ckwcfm/learn-go/rss/routes/actions/dialogs"
	"github.com/ckwcfm/learn-go/rss/routes/actions/pages"
)

func ActionRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/pages/home", pages.HomePage)
	router.HandleFunc("/pages/about", pages.AboutPage)
	router.HandleFunc("/dialogs/homeDialog", dialogs.ActionHomeDialog)
	return http.StripPrefix("/actions", router)
}
