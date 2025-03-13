package pages

import (
	"html/template"
	"net/http"

	"github.com/ckwcfm/learn-go/rss/templates/layouts"
	"github.com/ckwcfm/learn-go/rss/templates/partials"
	"github.com/ckwcfm/learn-go/rss/templates/views/contents"
)

func Book(w http.ResponseWriter, r *http.Request) {
	main := layouts.MainLayout.Layout()
	sidenav := partials.Sidenav.Sidenav()
	contentData := contents.BookContentData{
		BookFormData: contents.BookFormData{
			Title:  "Book",
			Author: "Book Description",
			Error:  "333333",
		},
	}
	content := contents.Book.Content(contentData)

	mainfile := main.Key.GetFile()
	sidenavfile := sidenav.Key.GetFile()
	contentfile := content.Key.GetFile()

	files := []string{
		mainfile,
		sidenavfile,
		contentfile,
	}

	templ := template.Must(template.ParseFiles(files...))
	templ.Execute(w, contentData)

}
