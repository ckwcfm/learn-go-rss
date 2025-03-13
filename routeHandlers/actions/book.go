package actions

import (
	"log"
	"net/http"

	"github.com/ckwcfm/learn-go/rss/constants"
	"github.com/ckwcfm/learn-go/rss/models"
	"github.com/ckwcfm/learn-go/rss/services"
	"github.com/ckwcfm/learn-go/rss/templates/views/contents"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateBook")
	userId := r.Context().Value(constants.UserIDKey).(string)
	userID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	title := r.FormValue("title")
	author := r.FormValue("author")
	if title == "" || author == "" {
		http.Error(w, "Title and author are required", http.StatusBadRequest)
		return
	}
	book := models.Book{
		Title:  title,
		Author: author,
		UserId: userID,
	}
	newBook, err := services.CreateBook(book)
	log.Println("newBook", newBook)
	if err != nil {
		log.Println("error creating book", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("creating book form")
	bookForm := contents.Book.Form(contents.BookFormData{
		Title:  "",
		Author: "",
		Error:  "",
	})
	bookListItem := contents.Book.OobListItem(newBook)
	bookForm.Render(w, r)
	bookListItem.Render(w, r)
}
