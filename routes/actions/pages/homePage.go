package pages

import (
	"html/template"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	log.Println("Home")
	tmpls := []string{
		"views/pages/home.html",
	}
	tmpl := template.Must(template.ParseFiles(tmpls...))
	tmpl.ExecuteTemplate(w, "content", nil)
}
