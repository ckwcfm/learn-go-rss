package services

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/ckwcfm/learn-go/rss/db"
	"github.com/ckwcfm/learn-go/rss/models"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func getUserCollection() *mongo.Collection {
	userCollection, err := db.GetCollection(db.UserCollection)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB", err)
	}
	return userCollection
}

func CreateUser(user models.User) error {
	userCollection := getUserCollection()
	user.Password = hashPassword(user.Password)

	_, err := userCollection.InsertOne(context.Background(), user)
	return err
}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash password", err)
	}
	return string(hashedPassword)
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := getUserCollection().FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	return user, err
}

func ValidateUser(email, password string) (models.User, error) {
	user, err := GetUserByEmail(email)
	if err != nil {
		return models.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func CreateToken(userID string) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 24 hour expiration
	})
	return token.SignedString([]byte(jwtSecret))
}
