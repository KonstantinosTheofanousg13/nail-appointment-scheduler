package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect() {

	if err := godotenv.Load(); err != nil {
		log.Println("Note: No .env file found")
	}

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("Error: 'MONGO_URI' variable is not set.")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB: ", err)
	}

	fmt.Println("Connected to MongoDB successfully!")
	Client = client
}

func GetContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}
