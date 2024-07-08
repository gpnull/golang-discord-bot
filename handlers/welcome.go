package handlers

import (
	"context"
	"fmt"

	"github.com/gpnull/golang-github.com/gpnull/golang-discord-bot/database"
	"github.com/gpnull/golang-github.com/gpnull/golang-discord-bot/models"

	"github.com/bwmarrin/discordgo"
)

// Welcome handles member join events
func Welcome(dbClient *database.MongoClient) func(*discordgo.Session, *discordgo.GuildMemberAdd) {
	return func(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
		// Save member information to MongoDB
		fmt.Print(m.User)
		user := models.User{
			DiscordID: m.User.ID,
			Username:  m.User.Username,
		}
		if err := dbClient.CreateUser(context.Background(), &user); err != nil {
			fmt.Println("Error saving member information:", err)
			return
		}

		// Add "newuser" role
		guildID := m.GuildID
		err := s.GuildMemberRoleAdd(guildID, m.User.ID, "newuser")
		if err != nil {
			fmt.Println("Error adding 'newuser' role:", err)
			return
		}

		// Send welcome message
		s.ChannelMessageSend(m.GuildID, fmt.Sprintf("Welcome %s to the server!", m.User.Mention()))
	}
}
