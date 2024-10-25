package config

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongodbUri := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(mongodbUri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	DB = client.Database(os.Getenv("DB_NAME"))
	fmt.Println("Database status : Connect to MongoDB Succeed!")
}

func GetBlogCollection() (*mongo.Collection, error) {
	if DB == nil {
		return nil, errors.New("database connection is not initialized")
	}
	return DB.Collection("blog"), nil
}
