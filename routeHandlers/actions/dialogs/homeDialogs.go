package dialogs

import (
	"html/template"
	"log"
	"net/http"
)

type DialogHeaderData struct {
	Title       string
	Description string
}

type DialogBodyData struct {
	Body string
}

type DialogData struct {
	Header DialogHeaderData
	Body   DialogBodyData
}

func ActionHomeDialog(w http.ResponseWriter, r *http.Request) {
	log.Println("ActionHomeDialog")
	tmpls := []string{
		"views/compontents/dialogs/logoutDialogBody.html",
		"views/compontents/dialogs/homeDialog.html",
	}
	dialogData := DialogData{
		Header: DialogHeaderData{
			Title:       "Dialog",
			Description: "Welcome to the dialog page",
		},
	}
	tmpl := template.Must(template.ParseFiles(tmpls...))
	tmpl.ExecuteTemplate(w, "dialog", dialogData)
}
