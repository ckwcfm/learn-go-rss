package pages

import (
	"html/template"
	"net/http"

	"github.com/ckwcfm/learn-go/rss/constants"
	"github.com/ckwcfm/learn-go/rss/models"
)

func About(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value(constants.UserKey).(models.User)
	tmpls := []string{
		"views/layouts/main.html",
		"views/partials/sidenav.html",
		"views/pages/about.html",
	}
	tmpl := template.Must(template.ParseFiles(tmpls...))
	type PageData struct {
		User models.User
	}
	data := PageData{
		User: user,
	}
	tmpl.Execute(w, data)
}
