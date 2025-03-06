package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var requiredEnvVars = []string{
	"JWT_SECRET",
	"MONGO_URI",
	"PORT",
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			log.Fatalf("Environment variable %s is not set", envVar)
		}
	}
}
