package pages

import (
	"html/template"
	"log"
	"net/http"

	"github.com/ckwcfm/learn-go/rss/constants"
	"github.com/ckwcfm/learn-go/rss/models"
	"github.com/ckwcfm/learn-go/rss/services"
)

func Post(w http.ResponseWriter, r *http.Request) {
	log.Println("Post")
	user := r.Context().Value(constants.UserKey).(models.User)
	posts, err := services.GetPostsForUser(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpls := []string{
		"views/layouts/main.html",
		"views/partials/sidenav.html",
		"views/pages/post.html",
	}
	tmpl := template.Must(template.ParseFiles(tmpls...))
	type FormData struct {
		Title   string
		Content string
	}
	type PageData struct {
		User  models.User
		Posts []models.Post
		Error string
		Form  FormData
	}
	data := PageData{
		User:  user,
		Posts: posts,
		Error: "",
		Form:  FormData{},
	}
	tmpl.Execute(w, data)
}
