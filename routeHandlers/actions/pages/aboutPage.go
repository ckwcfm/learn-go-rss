package pages

import (
	"html/template"
	"log"
	"net/http"

	"github.com/ckwcfm/learn-go/rss/constants"
	"github.com/ckwcfm/learn-go/rss/models"
)

func AboutPage(w http.ResponseWriter, r *http.Request) {
	log.Println("AboutPage")
	User := r.Context().Value(constants.UserKey).(models.User)
	log.Println(User)
	tmpls := []string{
		"views/pages/about.html",
	}
	type PageData struct {
		User models.User
	}
	data := PageData{
		User: User,
	}
	tmpl := template.Must(template.ParseFiles(tmpls...))

	tmpl.ExecuteTemplate(w, "content", data)

}
