package models

import (
	"gorm.io/gorm"
)

// User struct represents a Discord user in SQL database
type UserDoctor struct {
	gorm.Model
	DiscordID                    string `gorm:"type:varchar(255);uniqueIndex"`
	Username                     string
	Email                        string
	Avatar                       string
	Locale                       string
	GlobalName                   string
	Verified                     bool
	Banner                       string
	TimekeepingChannelID         string
	TimekeepingOvertimeChannelID string
}

func (UserDoctor) TableName() string {
	return "user_doctor"
}
