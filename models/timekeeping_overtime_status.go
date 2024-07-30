package models

import (
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

type TimekeepingOvertimeStatus struct {
	gorm.Model
	ButtonID                        string `gorm:"type:varchar(255);uniqueIndex"`
	Label                           string
	Style                           discordgo.ButtonStyle
	Content                         string
	TimekeepingOvertimeChannelID    string
	TimekeepingOvertimeLogChannelID string
	Status                          string `gorm:"type:enum('working', 'stopped')"`
}

func (TimekeepingOvertimeStatus) TableName() string {
	return "timekeeping_overtime_status"
}
