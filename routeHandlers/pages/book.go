package pages

import (
	"html/template"
	"net/http"

	"github.com/ckwcfm/learn-go/rss/constants"
	"github.com/ckwcfm/learn-go/rss/services"
	"github.com/ckwcfm/learn-go/rss/templates/layouts"
	"github.com/ckwcfm/learn-go/rss/templates/partials"
	"github.com/ckwcfm/learn-go/rss/templates/views/contents"
	"github.com/ckwcfm/learn-go/rss/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Book(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(constants.UserIDKey).(primitive.ObjectID)
	page := utils.GetQueryWithDefault(r, "page", 1)
	limit := 10
	bookListData, err := services.GetBookListData(userID, page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	main := layouts.MainLayout.Layout()
	sidenav := partials.Sidenav.Sidenav()
	contentData := contents.BookContentData{
		BookFormData: contents.BookFormData{
			Title:  "Book",
			Author: "Book Description",
			Error:  "333333",
		},
		BookListData: bookListData,
	}
	content := contents.Book.Content(contentData)

	mainfile := main.Key.GetFile()
	sidenavfile := sidenav.Key.GetFile()
	contentfile := content.Key.GetFile()

	files := []string{
		mainfile,
		sidenavfile,
		contentfile,
	}

	templ := template.Must(template.ParseFiles(files...))
	templ.Execute(w, contentData)

}
