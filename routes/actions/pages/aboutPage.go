package pages

import (
	"html/template"
	"net/http"
)

func AboutPage(w http.ResponseWriter, r *http.Request) {
	tmpls := []string{
		"views/pages/about.html",
	}

	tmpl := template.Must(template.ParseFiles(tmpls...))

	tmpl.ExecuteTemplate(w, "content", nil)

}
