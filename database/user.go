package database

import (
	"context"
	"fmt"

	"github.com/gpnull/golang-github.com/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateUser creates a new user in MongoDB or updates an existing user if the discord_id already exists
func (mc *MongoClient) CreateUser(ctx context.Context, user *models.User) error {
	filter := bson.M{"discord_id": user.DiscordID}
	update := bson.M{"$set": user}
	_, err := mc.users.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("failed to create/update user: %v", err)
	}
	return nil
}
