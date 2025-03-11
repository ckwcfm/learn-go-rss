package routes

import (
	"net/http"

	"github.com/ckwcfm/learn-go/rss/middlewares"
	"github.com/ckwcfm/learn-go/rss/routes/actions/dialogs"
	"github.com/ckwcfm/learn-go/rss/routes/actions/pages"
)

func ActionRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/pages/home", pages.HomePage)
	router.Handle("/pages/about", middlewares.Authorization(http.HandlerFunc(pages.AboutPage)))
	router.HandleFunc("/dialogs/homeDialog", dialogs.ActionHomeDialog)
	return http.StripPrefix("/actions", router)
}
