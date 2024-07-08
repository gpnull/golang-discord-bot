package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Goodbye handles member leave events
func Goodbye() func(*discordgo.Session, *discordgo.GuildMemberRemove) {
	return func(s *discordgo.Session, m *discordgo.GuildMemberRemove) {
		// Send goodbye message
		s.ChannelMessageSend(m.GuildID, fmt.Sprintf("Goodbye %s!", m.User.Username))
	}
}
