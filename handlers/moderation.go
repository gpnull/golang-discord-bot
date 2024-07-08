package handlers

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Moderation handles moderation commands
func Moderation(s *discordgo.Session, idRole string) func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Check if the message is a command
		if !strings.HasPrefix(m.Content, "/") {
			return
		}

		// Split the command and arguments
		args := strings.Fields(m.Content)
		command := args[0]

		// Handle the /approve command
		if command == "/approve" && len(args) == 2 {
			userID := args[1]

			// Add the "auto-role" role to the specified user
			err := s.GuildMemberRoleAdd(m.GuildID, userID, idRole)
			if err != nil {
				fmt.Println("Error adding 'auto-role' role:", err)
				return
			}

			// Send a notification message
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Added 'auto-role' role to <@%s>!", userID))
		}
	}
}
