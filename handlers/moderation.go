package handlers

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Moderation handles moderation commands
func Moderation(s *discordgo.Session) func(*discordgo.Session, *discordgo.MessageCreate) {
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

			// Add the "newuser" role to the specified user
			err := s.GuildMemberRoleAdd(m.GuildID, userID, "newuser")
			if err != nil {
				fmt.Println("Error adding 'newuser' role:", err)
				return
			}

			// Send a notification message
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Added 'newuser' role to <@%s>!", userID))
		}
	}
}
