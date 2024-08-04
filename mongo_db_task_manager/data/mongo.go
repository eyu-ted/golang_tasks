package data

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	// "github.com/joho/godotenv"
)

var Client *mongo.Client
var TaskCollection *mongo.Collection

func ConnectDB() {
	// uri,OK := os.LookupEnv("MONGODB_URI")
	uri := os.Getenv("MONGODB_URI")
	// if OK == false{
	// 	log.Fatal("MONGODB_URI environment variable is not set")
		
	// }
	if uri == "" {
		log.Fatal("MONGODB_URI environment variable is not set")
	}

	clientOptions := options.Client().ApplyURI(uri)

	// Create a context with a longer timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	Client = client
	TaskCollection = client.Database("task_manager").Collection("tasks")
}
