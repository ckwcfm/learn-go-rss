package auth

import (
	"html/template"
	"net/http"

	"github.com/ckwcfm/learn-go/rss/models"
	"github.com/ckwcfm/learn-go/rss/services"
	"github.com/ckwcfm/learn-go/rss/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	tmpls := []string{
		"views/layouts/main.html",
		"views/pages/auth/register.html",
	}
	tmpl := template.Must(template.ParseFiles(tmpls...))
	tmpl.Execute(w, nil)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	email := r.FormValue("email")
	password := r.FormValue("password")
	user := models.User{
		Email:    email,
		Password: password,
	}
	err = services.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token, err := services.CreateToken(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, utils.CreateTokenCookie(token))
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
