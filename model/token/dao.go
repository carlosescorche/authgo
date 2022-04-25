package token

import (
	"context"

	"github.com/carlosescorche/authgo/utils/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func dbCollection() (*mongo.Collection, error) {
	if collection != nil {
		return collection, nil
	}

	database, err := db.Get()
	if err != nil {
		return nil, err
	}

	collection = database.Collection("tokens")

	_, err = collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.M{
				"userId": 1,
			}, Options: nil,
		},
	)

	if err != nil {
		return nil, err
	}

	return collection, nil
}

func insert(userID primitive.ObjectID) (*Token, error) {
	collection, err := dbCollection()
	if err != nil {
		return nil, err
	}

	token := newToken(userID)

	_, err = collection.InsertOne(context.Background(), token)
	if err != nil {
		return nil, err
	}

	return token, nil

}

func findByID(tokenID string) (*Token, error) {
	collection, err := dbCollection()
	if err != nil {
		return nil, err
	}

	_id, err := primitive.ObjectIDFromHex(tokenID)
	if err != nil {
		return nil, err
	}

	token := &Token{}
	filter := bson.M{"_id": _id}

	if err = collection.FindOne(context.Background(), filter).Decode(token); err != nil {
		return nil, err
	}

	return token, nil
}
