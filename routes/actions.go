package routes

import (
	"net/http"

	"github.com/ckwcfm/learn-go/rss/middlewares"
	"github.com/ckwcfm/learn-go/rss/routeHandlers/actions"
	"github.com/ckwcfm/learn-go/rss/routeHandlers/actions/dialogs"
	"github.com/ckwcfm/learn-go/rss/routeHandlers/actions/pages"
)

func ActionRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/pages/home", pages.HomePage)
	router.Handle("/pages/", middlewares.IsUser(protectedPageRoutes()))
	router.Handle("POST /posts", middlewares.IsUser(http.HandlerFunc(actions.CreatePost)))
	router.Handle("POST /books", middlewares.IsUser(http.HandlerFunc(actions.CreateBook)))
	router.HandleFunc("/dialogs/homeDialog", dialogs.ActionHomeDialog)
	return http.StripPrefix("/actions", router)
}

func protectedPageRoutes() http.Handler {
	router := http.NewServeMux()
	router.Handle("/pages/about", http.HandlerFunc(pages.AboutPage))
	router.Handle("/pages/posts", http.HandlerFunc(pages.PostPage))
	router.Handle("/pages/books", http.HandlerFunc(pages.BookPage))
	return router
}
