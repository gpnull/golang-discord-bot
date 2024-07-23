package models

import (
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

type TimekeepingStatus struct {
	gorm.Model
	ButtonID                string `gorm:"type:varchar(255);uniqueIndex"`
	Label                   string
	Style                   discordgo.ButtonStyle
	Content                 string
	TimekeepingChannelID    string
	TimekeepingLogChannelID string
}

func (TimekeepingStatus) TableName() string {
	return "timekeeping_status"
}
