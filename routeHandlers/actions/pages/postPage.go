package pages

import (
	"html/template"
	"net/http"

	"github.com/ckwcfm/learn-go/rss/constants"
	"github.com/ckwcfm/learn-go/rss/models"
	"github.com/ckwcfm/learn-go/rss/services"
)

func PostPage(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(constants.UserKey).(models.User)
	posts, err := services.GetPostsForUser(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpls := []string{
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
	tmpl.ExecuteTemplate(w, "content", data)
}
