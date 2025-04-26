package routes

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBinstance creates a MongoDB client
func DBinstance() *mongo.Client {
	// Optional: load .env if needed
	/*
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	*/

	MongoDb := os.Getenv("MONGODB_URL")
	if MongoDb == "" {
		log.Fatal("MONGODB_URL not set in environment")
	}

	clientOptions := options.Client().
		ApplyURI(MongoDb).
		SetTLSConfig(&tls.Config{}) // ADD THIS LINE for TLS!

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ Connected to MongoDB!")
	return client
}

// Client Database instance
var Client *mongo.Client = DBinstance()

// OpenCollection connects to a specific collection
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("cluster0").Collection(collectionName)
}
