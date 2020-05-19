package mongodb

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("\033[7;49;31m  Error connecting to MongoDB  \033[0m\n")
		os.Exit(1)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println("\033[7;49;31m  Error ping Mongo  \033[0m\n")
		os.Exit(1)
	}
	fmt.Println("\033[7;49;92m  Connected to MongoDB  \033[0m")
}

// GetConnection returns connection pool
func GetConnection() *mongo.Client {
	return client
}
