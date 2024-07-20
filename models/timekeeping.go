package models

import (
	"github.com/bwmarrin/discordgo"
)

type Timekeeping struct {
	ID       string                `bson:"_id,omitempty"`
	ButtonID string                `bson:"button_id"`
	Label    string                `bson:"label"`
	Style    discordgo.ButtonStyle `bson:"style"`
}
