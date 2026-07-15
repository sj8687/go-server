package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserCollection *mongo.Collection

func ConnectDB() {

	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI("mongodb://localhost:27017"),
	)

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(
		context.Background(), 10*time.Second
	)
	defer cancel()

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	UserCollection = client.Database("golang_auth").Collection("users")

	log.Println("Mongo Connected")
}