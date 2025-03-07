package pages

import (
	"html/template"
	"log"
	"net/http"

	"github.com/ckwcfm/learn-go/rss/constants"
	"github.com/ckwcfm/learn-go/rss/models"
	"github.com/ckwcfm/learn-go/rss/services"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpls := []string{
		"views/layouts/main.html",
		"views/partials/sidenav.html",
		"views/pages/home.html",
	}

	tmpl := template.Must(template.ParseFiles(tmpls...))
	tmpl.Execute(w, nil)
}

func About(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(constants.UserIDKey).(string)
	log.Println(userID)
	tmpls := []string{
		"views/layouts/main.html",
		"views/partials/sidenav.html",
		"views/pages/about.html",
	}
	tmpl := template.Must(template.ParseFiles(tmpls...))
	user, err := services.GetUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	type PageData struct {
		User models.User
	}
	data := PageData{
		User: user,
	}
	tmpl.Execute(w, data)
}
