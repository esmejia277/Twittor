package db

import (
	"context"
	"fmt"
	"time"

	"github.com/esmejia277/twittor/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchProfile(id string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := mongoClient.Database("twittor")
	collection := db.Collection("Users")
	var profile models.User
	objectId, _ := primitive.ObjectIDFromHex(id)
	condition := bson.M{"_id": objectId}
	err := collection.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		fmt.Println("Not found: " + err.Error())
		return profile, err
	}
	return profile, nil

}
