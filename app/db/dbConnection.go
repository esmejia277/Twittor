package db

import (
	"context"
	"log"
	"time"

	"github.com/esmejia277/twittor/app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb://127.0.0.1:27017/?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false")

func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	return client
}

func IsDBConnected() int {
	err := mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}

func InsertIntoDatabase(model models.User) (string, bool, error) {
	ctx, cancel := (context.WithTimeout(context.Background(), 15*time.Second))
	defer cancel()
	db := mongoClient.Database("twittor")
	collection := db.Collection("Users")

	model.Password, _ = HashPassword(model.Password)
	result, error := collection.InsertOne(ctx, model)
	if error != nil {
		return "", false, error
	}
	ObjectID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjectID.String(), true, nil

}
