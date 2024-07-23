package database

import (
	"fmt"

	"github.com/gpnull/golang-github.com/models"
	"gorm.io/gorm/clause"
)

func (db *Database) CreateUser(user *models.User) error {
	result := db.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(user)
	if result.Error != nil {
		return fmt.Errorf("failed to create/update user: %v", result.Error)
	}
	return nil
}

func (db *Database) UpdateUserChannelID(discordId, timekeepingChannelId string) error {
	result := db.DB.Model(&models.User{}).Where("discord_id = ?", discordId).Update("timekeeping_channel_id", timekeepingChannelId)
	return result.Error
}

func (db *Database) GetChannelIDByDiscordID(discordId string) string {
	var user models.User
	result := db.DB.Model(&models.User{}).Where("discord_id = ?", discordId).First(&user)
	if result.Error != nil {
		return ""
	}
	return user.TimekeepingChannelID
}
