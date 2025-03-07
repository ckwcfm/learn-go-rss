package pages

import (
	"html/template"
	"net/http"
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
	tmpls := []string{
		"views/layouts/main.html",
		"views/partials/sidenav.html",
		"views/pages/about.html",
	}
	tmpl := template.Must(template.ParseFiles(tmpls...))
	tmpl.Execute(w, nil)
}
