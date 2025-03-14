package pages

import (
	"net/http"

	"github.com/ckwcfm/learn-go/rss/constants"
	"github.com/ckwcfm/learn-go/rss/services"
	"github.com/ckwcfm/learn-go/rss/templates/views/contents"
	"github.com/ckwcfm/learn-go/rss/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BookPage(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(constants.UserIDKey).(primitive.ObjectID)
	// Get page number from query params, default to 1 if not provided
	page := utils.GetQueryWithDefault(r, "page", 1)
	limit := 10
	bookListData, err := services.GetBookListData(userID, page, limit)
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
		BookListData: bookListData,
	})

	content.Render(w, r)
}
