package services

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/ckwcfm/learn-go/rss/db"
	"github.com/ckwcfm/learn-go/rss/models"
	"github.com/ckwcfm/learn-go/rss/templates/views/contents"
	"github.com/ckwcfm/learn-go/rss/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

func getBookCollection() *mongo.Collection {
	bookCollection, err := db.GetCollection(db.BookCollection)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB", err)
	}
	return bookCollection
}

func CreateBook(book models.Book) (models.Book, error) {
	bookCollection := getBookCollection()
	if book.Title == "" || book.Author == "" {
		return models.Book{}, errors.New("title and author are required")
	}
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()
	newBook, err := bookCollection.InsertOne(context.Background(), book)
	if err != nil {
		return models.Book{}, err
	}
	book.ID = newBook.InsertedID.(primitive.ObjectID)
	return book, nil
}

func GetBooksForUser(userID primitive.ObjectID, page int, limit int) ([]models.Book, error) {

	bookCollection := getBookCollection()
	options := utils.NewMongoPagination(limit, page).GetPaginationOptions()
	cursor, err := bookCollection.Find(context.Background(), bson.M{"userId": userID}, options)
	if err != nil {
		return nil, err
	}
	var books []models.Book
	if err := cursor.All(context.Background(), &books); err != nil {
		return nil, err
	}
	return books, nil
}

func GetBookTotalPagesForUser(userID primitive.ObjectID, limit int) (int, error) {
	total, err := CountBooksForUser(userID)
	if err != nil {
		return 0, err
	}
	return int(total) / limit, nil
}

func CountBooksForUser(userID primitive.ObjectID) (int64, error) {
	bookCollection := getBookCollection()
	count, err := bookCollection.CountDocuments(context.Background(), bson.M{"userId": userID})
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetBookListData(userID primitive.ObjectID, page int, limit int) (contents.BookListData, error) {
	books, err := GetBooksForUser(userID, page, limit)
	if err != nil {
		return contents.BookListData{}, err
	}
	totalPages, err := GetBookTotalPagesForUser(userID, limit)
	if err != nil {
		return contents.BookListData{}, err
	}
	nextPage := min(page+1, totalPages)
	prevPage := max(page-1, 1)
	hasNext := page < totalPages
	hasPrev := page > 1
	return contents.BookListData{
		Books:       books,
		CurrentPage: page,
		TotalPages:  totalPages,
		NextPage:    nextPage,
		PrevPage:    prevPage,
		HasNext:     hasNext,
		HasPrev:     hasPrev,
	}, nil
}
