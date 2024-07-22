package models

import (
	"gorm.io/gorm"
)

// User struct represents a Discord user in SQL database
type User struct {
	gorm.Model
	DiscordID            string `gorm:"type:varchar(255);uniqueIndex"`
	Username             string
	Email                string
	Avatar               string
	Locale               string
	GlobalName           string
	Verified             bool
	Banner               string
	TimekeepingChannelID string
}

func (User) TableName() string {
	return "user"
}
