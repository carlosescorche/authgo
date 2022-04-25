package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Get() (*mongo.Database, error) {

	clientOptions := options.Client().ApplyURI(os.Getenv("mongodb"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	database := client.Database(os.Getenv("db"))

	return database, nil
}
