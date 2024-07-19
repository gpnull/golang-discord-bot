package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TimekeepingExtra struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Day   string             `bson:"day"`
	Start primitive.DateTime `bson:"start"`
	End   primitive.DateTime `bson:"end"`
}
