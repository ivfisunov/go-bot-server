package mongodb

import (
	"bot-backend/types"
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Printf("\033[7;49;31m  Error connecting to MongoDB  \033[0m\n")
		os.Exit(1)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Printf("\033[7;49;31m  Error ping Mongo  \033[0m\n")
		os.Exit(1)
	}
	fmt.Printf("\n\033[7;49;92m  Connected to MongoDB  \033[0m\n\n")
}

// GetConnection returns connection pool
func GetConnection() *mongo.Client {
	return client
}

// SearchUserByMobile finds user in DB
func SearchUserByMobile(userObj types.MobileJSONRequest) types.FindUserByMobileResponse {
	collection := client.Database("nspk").Collection("users")

	filter := bson.M{
		"mobileNumber": userObj.Mobile,
		"state":        bson.M{"$ne": "Увольнение"},
	}

	var user types.FindUserByMobileResponse
	cxt := context.Background()
	result := collection.FindOne(cxt, filter)
	if result.Err() != nil {
		user.Message = "User NOT found"
		user.Status = "Error"
		return user
	}
	result.Decode(&user)
	user.Message = "User successfuly found"
	user.Status = "OK"
	return user
}

// SearchUserByLastName finds users in DB
func SearchUserByLastName(userObj types.UserLastNameJSONRequest) types.FindUserByLastNameResponse {
	collection := client.Database("nspk").Collection("users")

	cxt := context.Background()
	filter := bson.M{
		"lastName": primitive.Regex{Pattern: userObj.LastName, Options: "i"},
		"state":    bson.M{"$ne": "Увольнение"},
	}

	var users types.FindUserByLastNameResponse
	users.Employees = []types.User{}
	users.Message = "Users not found."
	users.Status = "Error"
	cursor, err := collection.Find(cxt, filter)
	if err != nil {
		return users
	}
	defer cursor.Close(cxt)

	for cursor.Next(cxt) {
		var user types.User
		cursor.Decode(&user)
		// if err != nil {
		// 	return users
		// }
		users.Employees = append(users.Employees, user)
	}

	if len(users.Employees) == 0 {
		return users
	}

	users.Message = "Users found."
	users.Status = "OK"
	return users
}
