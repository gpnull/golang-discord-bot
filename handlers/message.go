package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Message handles messages in channel
func Message(channelId string) func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Check if the message is from a bot
		if m.Author.Bot {
			return
		}

		// Check the channel and message content
		if m.ChannelID == channelId && strings.ToLower(m.Content) == "ahihi" {
			// Use ChannelMessageSendReply to reply to the specific message
			s.ChannelMessageSendReply(m.ChannelID, "ahihi", m.Reference())
		}
	}
}
