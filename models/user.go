package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User struct represents a Discord user in MongoDB
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	DiscordID string             `bson:"discord_id"`
	Username  string             `bson:"username"`
}
