package db

import (
	"context"
	"time"

	"github.com/esmejia277/twittor/app/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckIfUserExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := mongoClient.Database("twittor")
	collection := db.Collection("Users")

	var result models.User

	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID

}
