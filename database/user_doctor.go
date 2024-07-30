package database

import (
	"fmt"

	"github.com/gpnull/golang-github.com/models"
	"gorm.io/gorm/clause"
)

func (db *Database) CreateUser(user *models.UserDoctor) error {
	result := db.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(user)
	if result.Error != nil {
		return fmt.Errorf("failed to create/update user: %v", result.Error)
	}
	return nil
}

func (db *Database) UpdateTimekeepingChannelID(discordId, timekeepingChannelId string) error {
	result := db.DB.Model(&models.UserDoctor{}).Where("discord_id = ?", discordId).
		Update("timekeeping_channel_id", timekeepingChannelId)
	return result.Error
}

func (db *Database) GetTimekeepingChannelIDByDiscordID(discordId string) string {
	var user models.UserDoctor
	result := db.DB.Model(&models.UserDoctor{}).Where("discord_id = ?", discordId).First(&user)
	if result.Error != nil {
		return ""
	}
	return user.TimekeepingChannelID
}

// for OT

func (db *Database) UpdateTimekeepingOvertimeChannelID(discordId, timekeepingOvertimeChannelId string) error {
	result := db.DB.Model(&models.UserDoctor{}).Where("discord_id = ?", discordId).
		Update("timekeeping_overtime_channel_id", timekeepingOvertimeChannelId)
	return result.Error
}

func (db *Database) GetTimekeepingOvertimeChannelIDByDiscordID(discordId string) string {
	var user models.UserDoctor
	result := db.DB.Model(&models.UserDoctor{}).Where("discord_id = ?", discordId).First(&user)
	if result.Error != nil {
		return ""
	}
	return user.TimekeepingOvertimeChannelID
}
