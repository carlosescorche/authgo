package user

import (
	"context"
	"fmt"

	"github.com/carlosescorche/authgo/utils/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	col := database.Collection("users")

	_, err = col.Indexes().CreateMany(
		context.Background(),
		[]mongo.IndexModel{{
			Keys:    bson.M{"username": 1},
			Options: options.Index().SetUnique(true),
		}, {
			Keys:    bson.M{"email": 1},
			Options: options.Index().SetUnique(true),
		}},
	)

	if err != nil {
		return nil, err
	}

	collection = col
	return collection, nil
}

func insert(user *User) (*User, error) {
	var collection, err = dbCollection()
	if err != nil {
		return nil, err
	}

	if _, err := collection.InsertOne(context.Background(), user); err != nil {
		return nil, err
	}

	return user, nil
}

func findByLogin(login string) (*User, error) {
	var collection, err = dbCollection()
	if err != nil {
		return nil, err
	}

	user := &User{}
	filter := bson.M{"$or": []bson.M{{"username": login}, {"email": login}}}
	err = collection.FindOne(context.Background(), filter).Decode(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func findByID(userID string) (*User, error) {
	var collection, err = dbCollection()
	if err != nil {
		return nil, err
	}

	_id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUserId, err)
	}

	user := &User{}
	filter := bson.M{"_id": _id}
	if err = collection.FindOne(context.Background(), filter).Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}
