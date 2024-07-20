package database

import (
	"context"
	"fmt"

	"github.com/gpnull/golang-github.com/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (mc *MongoClient) SaveButton(ctx context.Context, timekeeping *models.Timekeeping) error {
	filter := bson.M{"_id": timekeeping.ID}
	update := bson.M{"$set": timekeeping}
	_, err := mc.timekeeping.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("failed to create/update button: %v", err)
	}
	return nil
}

func (mc *MongoClient) GetButtons(ctx context.Context) ([]*models.Timekeeping, error) {
	cursor, err := mc.timekeeping.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve buttons: %v", err)
	}
	defer cursor.Close(ctx)

	var buttons []*models.Timekeeping
	for cursor.Next(ctx) {
		var button models.Timekeeping
		if err := cursor.Decode(&button); err != nil {
			return nil, fmt.Errorf("failed to decode button: %v", err)
		}
		buttons = append(buttons, &button)
	}

	return buttons, nil
}
