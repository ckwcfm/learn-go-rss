package pages

import (
	"html/template"
	"log"
	"net/http"

	"github.com/ckwcfm/learn-go/rss/constants"
	"github.com/ckwcfm/learn-go/rss/models"
	"github.com/ckwcfm/learn-go/rss/services"
)

func AboutPage(w http.ResponseWriter, r *http.Request) {
	log.Println("AboutPage")
	log.Println(r.Context().Value(constants.UserIDKey))

	userID, ok := r.Context().Value(constants.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	User, err := services.GetUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
