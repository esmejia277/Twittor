package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name,omitempty" json:"name"`
	LastName string             `bson:"lastName" json:"lastName,omitempty"`
	Birthday time.Time          `bson:"birthday" json:"birthday,omitempty"`
	Email    string             `bson:"email" json:"email,omitempty"`
	Password string             `bson:"password" json:"password,omitempty"`
	Avatar   string             `bson:"avatar" json:"avatar,omitempty"`
	Banner   string             `bson:"banner" json:"banner,omitempty"`
	Location string             `bson:"location" json:"location,omitempty"`
}
