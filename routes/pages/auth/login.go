package auth

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/ckwcfm/learn-go/rss/services"
)

func Login(w http.ResponseWriter, r *http.Request) {
	tmpls := []string{
		"views/layouts/main.html",
		"views/pages/auth/login.html",
	}
	tmpl := template.Must(template.ParseFiles(tmpls...))
	var data struct {
		Email    string
		Password string
	}
	if env := os.Getenv("ENV"); env == "development" {
		data.Email = "kevinwong@gmail.com"
		data.Password = "12345678"
	}

	tmpl.Execute(w, data)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("LoginHandler")
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	email := r.FormValue("email")
	password := r.FormValue("password")

	log.Println(email, password)
	log.Println(r.Form)

	token, err := services.Login(email, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    "Bearer " + token,
		Path:     "/",
		MaxAge:   3600,
		Secure:   os.Getenv("ENV") == "production",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)
	w.Header().Add("HX-Redirect", "/about")
}
