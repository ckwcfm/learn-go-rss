package services

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/ckwcfm/learn-go/rss/db"
	"github.com/ckwcfm/learn-go/rss/models"
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

func GetBooksForUser(userID primitive.ObjectID) ([]models.Book, error) {
	bookCollection := getBookCollection()
	cursor, err := bookCollection.Find(context.Background(), bson.M{"userId": userID})
	if err != nil {
		return nil, err
	}
	var books []models.Book
	if err := cursor.All(context.Background(), &books); err != nil {
		return nil, err
	}
	return books, nil
}
