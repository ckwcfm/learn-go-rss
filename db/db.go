package db

import (
	"context"
	"errors"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

type Collection struct {
	Collection *mongo.Collection
}

type CollectionName string

const (
	UserCollection CollectionName = "users"
	PostCollection CollectionName = "posts"
	BookCollection CollectionName = "books"
)
const database = "rss"

func ConnectToMongo(ctx context.Context) error {
	log.Println("Connecting to MongoDB")
	err := godotenv.Load()
	if err != nil {
		return err
	}
	MongoURI := os.Getenv("MONGO_URI")
	if MongoURI == "" {
		return errors.New("MONGO_URI is not set")
	}

	options := options.Client().ApplyURI(MongoURI)
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}
	log.Println("Connected to MongoDB")

	mongoClient = client

	createIndexes(ctx)

	return nil
}

func DisconnectMongo(ctx context.Context) error {
	err := mongoClient.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func GetMongoClient() (*mongo.Client, error) {
	if mongoClient == nil {
		return nil, errors.New("MongoDB client not initialized")
	}
	return mongoClient, nil
}

func GetCollection(name CollectionName) (*mongo.Collection, error) {
	client, err := GetMongoClient()
	if err != nil {
		return nil, err
	}
	return client.Database(database).Collection(string(name)), nil
}

func createIndexes(ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		createUserIndex(ctx)
		createPostIndex(ctx)
	}()

	wg.Wait()
}

func createUserIndex(ctx context.Context) {
	userCollection, err := GetCollection(UserCollection)

	if err != nil {
		log.Fatal("Failed to connect to MongoDB", err)
	}
	userCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	})

	userCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{"createdAt": 1},
	})
}

func createPostIndex(ctx context.Context) {
	postCollection, err := GetCollection(PostCollection)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB", err)
	}
	postCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{"createdAt": 1},
	})
	postCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{"updatedAt": 1},
	})
	postCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{"userId": 1},
	})
}
