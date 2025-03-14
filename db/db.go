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
		configUserCollection(ctx)
		configPostCollection(ctx)
		configBookCollection(ctx)
	}()

	wg.Wait()
}

func configUserCollection(ctx context.Context) {
	userCollection, err := GetCollection(UserCollection)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB", err)
	}
	createUserSchema(userCollection, ctx)
	createUserIndex(userCollection, ctx)

}

func createUserSchema(collection *mongo.Collection, ctx context.Context) {
	const bcryptHashLength = 60
	schema := bson.D{
		{Key: "bsonType", Value: "object"},
		{Key: "title", Value: "User Schema"},
		{Key: "required", Value: bson.A{"email", "password", "createdAt", "updatedAt"}},
		{Key: "properties", Value: bson.D{
			{Key: "email", Value: bson.D{
				{Key: "bsonType", Value: "string"},
				{Key: "pattern", Value: "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"},
				{Key: "description", Value: "email must be a string and is required"},
			}},
			{Key: "password", Value: bson.D{
				{Key: "bsonType", Value: "string"},
				{Key: "minLength", Value: bcryptHashLength},
				{Key: "description", Value: "password must be a string and is required"},
			}},
			{Key: "createdAt", Value: bson.D{
				{Key: "bsonType", Value: "date"},
				{Key: "description", Value: "createdAt must be a date and is required"},
			}},
			{Key: "updatedAt", Value: bson.D{
				{Key: "bsonType", Value: "date"},
				{Key: "description", Value: "updatedAt must be a date"},
			}},
		}},
	}

	command := bson.D{
		{Key: "collMod", Value: "users"},
		{Key: "validator", Value: bson.D{
			{Key: "$jsonSchema", Value: schema},
		}},
		{Key: "validationLevel", Value: "strict"},
	}
	err := collection.Database().RunCommand(ctx, command).Err()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB", err)
	}
}

func createUserIndex(collection *mongo.Collection, ctx context.Context) {
	collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	})
	collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{"createdAt": 1},
	})
}

func configPostCollection(ctx context.Context) {
	postCollection, err := GetCollection(PostCollection)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB", err)
	}
	createPostIndex(postCollection, ctx)
}

func createPostIndex(collection *mongo.Collection, ctx context.Context) {
	collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{"createdAt": 1},
	})
	collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{"updatedAt": 1},
	})
	collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{"userId": 1},
	})
}

func configBookCollection(ctx context.Context) {
	bookCollection, err := GetCollection(BookCollection)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB Book Index", err)
	}
	createBookSchema(bookCollection, ctx)
	createBookIndex(bookCollection, ctx)
}

func createBookSchema(collection *mongo.Collection, ctx context.Context) {

	schema := bson.D{
		{Key: "bsonType", Value: "object"},
		{Key: "title", Value: "Book Schema"},
		{Key: "required", Value: bson.A{"title", "author", "userId", "createdAt", "updatedAt"}},
		{Key: "properties", Value: bson.D{
			{Key: "title", Value: bson.D{
				{Key: "bsonType", Value: "string"},
				{Key: "minLength", Value: 1},
				{Key: "maxLength", Value: 100},
				{Key: "description", Value: "title must be a string and is required"},
			}},
			{Key: "author", Value: bson.D{
				{Key: "bsonType", Value: "string"},
				{Key: "minLength", Value: 1},
				{Key: "maxLength", Value: 100},
				{Key: "description", Value: "author must be a string and is required"},
			}},
			{Key: "createdAt", Value: bson.D{
				{Key: "bsonType", Value: "date"},
				{Key: "description", Value: "createdAt must be a date and is required"},
			}},
			{Key: "updatedAt", Value: bson.D{
				{Key: "bsonType", Value: "date"},
				{Key: "description", Value: "updatedAt must be a date"},
			}},
			{Key: "userId", Value: bson.D{
				{Key: "bsonType", Value: "objectId"},
				{Key: "description", Value: "userId must be an objectId and is required"},
			}},
		}},
	}

	command := bson.D{
		{Key: "collMod", Value: "books"},
		{Key: "validator", Value: bson.D{
			{Key: "$jsonSchema", Value: schema},
		}},
		{Key: "validationLevel", Value: "strict"},
		{Key: "validationAction", Value: "error"},
	}

	var result bson.M
	err := collection.Database().RunCommand(ctx, command).Decode(&result)
	log.Println("Schema result:", result)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB Book Index line 188 ", err)
	}

}

func createBookIndex(collection *mongo.Collection, ctx context.Context) {
	collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{"createdAt": 1},
	})
	collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{"updatedAt": 1},
	})
	collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{"userId": 1},
	})
	collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{"title": 1},
	})
	collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{"author": 1},
	})
}
