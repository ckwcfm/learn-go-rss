package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ckwcfm/learn-go/rss/db"
	"github.com/ckwcfm/learn-go/rss/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SeedDatabase() {
	seedBooks()
	log.Println("Database seeded")
}

func seedBooks() {
	userId, err := primitive.ObjectIDFromHex("67ca6b6d39da5bd7450309da")
	if err != nil {
		log.Fatal("Failed to convert user ID to ObjectID", err)
	}
	books := make([]interface{}, 100)
	for i := range 100 {
		books[i] = models.Book{
			Title:     fmt.Sprintf("Book %d", i),
			Author:    fmt.Sprintf("Author %d", i),
			UserId:    userId,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	}

	bookCollection, err := db.GetCollection(db.BookCollection)

	if err != nil {
		log.Fatal("Failed to connect to MongoDB", err)
	}
	_, err = bookCollection.InsertMany(context.Background(), books)
	if err != nil {
		log.Fatal("Failed to insert books", err)
	}

}
