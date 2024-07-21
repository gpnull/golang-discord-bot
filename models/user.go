package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User struct represents a Discord user in MongoDB
type User struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty"`
	DiscordID            string             `bson:"discord_id"`
	Username             string             `bson:"username"`
	Email                string             `bson:"email"`
	Avatar               string             `bson:"avatar"`
	Locale               string             `bson:"locale"`
	GlobalName           string             `bson:"global_name"`
	Verified             bool               `bson:"verified"`
	Banner               string             `bson:"banner"`
	TimekeepingChannelID string             `bson:"timekeeping_channel_id"`
}
