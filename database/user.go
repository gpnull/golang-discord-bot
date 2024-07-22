package database

import (
	"context"
	"fmt"

	"github.com/gpnull/golang-github.com/models"
	"gorm.io/gorm/clause"
)

// CreateUser creates a new user in SQL database or updates an existing user if the discord_id already exists
func (db *Database) CreateUser(ctx context.Context, user *models.User) error {
	result := db.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(user)
	if result.Error != nil {
		return fmt.Errorf("failed to create/update user: %v", result.Error)
	}
	return nil
}

// UpdateUserChannelID updates the timekeeping_channel_id for a user in SQL database
func (db *Database) UpdateUserChannelID(ctx context.Context, discordId, timekeepingChannelId string) error {
	result := db.DB.Model(&models.User{}).Where("discord_id = ?", discordId).Update("timekeeping_channel_id", timekeepingChannelId)
	return result.Error
}
