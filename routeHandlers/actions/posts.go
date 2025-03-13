package actions

import (
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/ckwcfm/learn-go/rss/constants"
	"github.com/ckwcfm/learn-go/rss/models"
	"github.com/ckwcfm/learn-go/rss/services"
	"github.com/go-playground/validator/v10"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating post: [post.go CreatePost | line 14]")
	user := r.Context().Value(constants.UserKey).(models.User)
	title := r.FormValue("title")
	content := r.FormValue("content")
	post := models.Post{
		Title:     title,
		Content:   content,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	validate := validator.New()
	if err := validate.Struct(post); err != nil {
		renderPostForm(w, err, title, content)
		return
	}
	newPost, err := services.CreatePost(post)

	if err != nil {
		renderPostForm(w, err, title, content)
		return
	}

	tmpls := []string{
		"views/pages/post.html",
	}
	type FormData struct {
		Title   string
		Content string
	}
	type PageData struct {
		Form  FormData
		Error string
	}
	pageData := PageData{
		Form: FormData{
			Title:   "",
			Content: "",
		},
		Error: "",
	}
	tmpl := template.Must(template.ParseFiles(tmpls...))
	tmpl.ExecuteTemplate(w, "post-form", pageData)
	tmpl.ExecuteTemplate(w, "oob-post-item", newPost)

}

func renderPostForm(w http.ResponseWriter, err error, title string, content string) {
	tmpls := []string{
		"views/pages/post.html",
	}
	tmpl := template.Must(template.ParseFiles(tmpls...))

	type FormData struct {
		Title   string
		Content string
	}
	type PageData struct {
		Error string
		Form  FormData
	}
	data := PageData{
		Error: err.Error(),
		Form: FormData{
			Title:   title,
			Content: content,
		},
	}
	w.WriteHeader(http.StatusUnprocessableEntity)
	tmpl.ExecuteTemplate(w, "post-form", data)

}
