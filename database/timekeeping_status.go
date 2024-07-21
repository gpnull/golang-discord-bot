package database

import (
	"context"
	"fmt"

	"github.com/gpnull/golang-github.com/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (mc *MongoClient) SaveTimeKeepingStatusButton(ctx context.Context, timekeeping *models.TimekeepingStatus) error {
	filter := bson.M{"_id": timekeeping.ID}
	update := bson.M{"$set": timekeeping}
	_, err := mc.timekeepingStatus.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("failed to create/update button: %v", err)
	}
	return nil
}

func (mc *MongoClient) GetTimeKeepingStatusButtons(ctx context.Context) ([]*models.TimekeepingStatus, error) {
	cursor, err := mc.timekeepingStatus.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve buttons: %v", err)
	}
	defer cursor.Close(ctx)

	var buttons []*models.TimekeepingStatus
	for cursor.Next(ctx) {
		var button models.TimekeepingStatus
		if err := cursor.Decode(&button); err != nil {
			return nil, fmt.Errorf("failed to decode button: %v", err)
		}
		buttons = append(buttons, &button)
	}

	return buttons, nil
}
