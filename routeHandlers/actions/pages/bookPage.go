package pages

import (
	"net/http"

	"github.com/ckwcfm/learn-go/rss/constants"
	"github.com/ckwcfm/learn-go/rss/services"
	"github.com/ckwcfm/learn-go/rss/templates/views/contents"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BookPage(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(constants.UserIDKey).(string)
	userID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	books, err := services.GetBooksForUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	content := contents.Book.Content(contents.BookContentData{
		BookFormData: contents.BookFormData{
			Title:  "Book",
			Author: "Book Author",
			Error:  "",
		},
		BookListData: contents.BookListData{
			Books: books,
		},
	})

	content.Render(w, r)
}
