package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Getenv(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}
	return os.Getenv(key)
}

func GetConnectURI() string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}
	return "mongodb+srv://" + Getenv("mongoUsername") + ":" + Getenv("mongoPassword") + "@cluster0.4lioy.mongodb.net/eight-sup?retryWrites=true&w=majority"
}
