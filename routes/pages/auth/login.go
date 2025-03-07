package auth

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/ckwcfm/learn-go/rss/services"
	"github.com/ckwcfm/learn-go/rss/utils"
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
		renderError(w, "Missing form data")
		return
	}
	email := r.FormValue("email")
	password := r.FormValue("password")

	log.Println(email, password)
	log.Println(r.Form)

	token, err := services.Login(email, password)
	if err != nil {
		renderError(w, "Invalid email or password")
		return
	}

	// set the cookie
	http.SetCookie(w, utils.CreateTokenCookie(token))
	// redirect to the about page
	w.Header().Add("HX-Redirect", "/about")
}

func renderError(w http.ResponseWriter, error string) {
	tmpl := template.Must(template.ParseFiles("views/pages/auth/login.html"))
	tmpl.ExecuteTemplate(w, "error", struct {
		Error string
	}{
		Error: error,
	})
}
