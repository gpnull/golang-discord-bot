package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient struct for interacting with MongoDB
type MongoClient struct {
	client      *mongo.Client
	users       *mongo.Collection
	timekeeping *mongo.Collection
}

// ConnectDB connects to MongoDB
func ConnectDB(uri string) (*MongoClient, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Ping the database to verify connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	db := client.Database("discord_bot")
	usersCollection := db.Collection("users")
	timekeepingCollection := db.Collection("timekeeping")

	return &MongoClient{
		client:      client,
		users:       usersCollection,
		timekeeping: timekeepingCollection,
	}, nil
}

// DisconnectDB disconnects from MongoDB
func (mc *MongoClient) DisconnectDB(ctx context.Context) error {
	err := mc.client.Disconnect(ctx)
	if err != nil {
		return fmt.Errorf("failed to disconnect from MongoDB: %v", err)
	}
	return nil
}
