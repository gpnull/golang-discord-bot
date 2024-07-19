package database

import (
	"context"
	"fmt"

	"github.com/gpnull/golang-github.com/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (mc *MongoClient) Test(ctx context.Context, timekeeping *models.Timekeeping) error {
	filter := bson.M{"discord_id": timekeeping.ID}
	update := bson.M{"$set": timekeeping}
	_, err := mc.timekeeping.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("failed to create/update timekeeping: %v", err)
	}
	return nil
}
