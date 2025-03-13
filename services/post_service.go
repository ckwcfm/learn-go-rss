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
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getPostCollection() *mongo.Collection {
	postCollection, err := db.GetCollection(db.PostCollection)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB", err)
	}
	return postCollection
}

func GetPostsForUser(userID primitive.ObjectID) ([]models.Post, error) {
	postCollection := getPostCollection()
	opts := options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}})
	cursor, err := postCollection.Find(context.Background(), bson.M{"userId": userID}, opts)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for cursor.Next(context.Background()) {
		var post models.Post
		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func CreatePost(post models.Post) (models.Post, error) {
	postCollection := getPostCollection()
	// TODO: validate post
	if post.Title == "" || post.Content == "" {
		return models.Post{}, errors.New("title and content are required")
	}
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()
	newPost, err := postCollection.InsertOne(context.Background(), post)
	if err != nil {
		return models.Post{}, err
	}
	post.ID = newPost.InsertedID.(primitive.ObjectID)
	return post, nil
}
