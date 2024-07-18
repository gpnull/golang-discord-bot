package database

import (
	"context"
	"fmt"

	"github.com/gpnull/golang-github.com/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient struct for interacting with MongoDB
type MongoClient struct {
	client *mongo.Client
	users  *mongo.Collection
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

	return &MongoClient{
		client: client,
		users:  usersCollection,
	}, nil
}

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

// DisconnectDB disconnects from MongoDB
func (mc *MongoClient) DisconnectDB(ctx context.Context) error {
	err := mc.client.Disconnect(ctx)
	if err != nil {
		return fmt.Errorf("failed to disconnect from MongoDB: %v", err)
	}
	return nil
}
