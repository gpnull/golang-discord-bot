package ready

import (
	"context"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/gpnull/golang-github.com/database"
	"github.com/gpnull/golang-github.com/models"
)

func GuildMemberAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd,
	dbClient *database.MongoClient, autoRoleId, welcomeChannelId string) {
	user := models.User{
		DiscordID:  m.User.ID,
		Username:   m.User.Username,
		Email:      m.User.Email,
		Avatar:     m.User.Avatar,
		Locale:     m.User.Locale,
		GlobalName: m.User.GlobalName,
		Verified:   m.User.Verified,
		Banner:     m.User.Banner,
	}

	if err := dbClient.CreateUser(context.Background(), &user); err != nil {
		fmt.Println("Error saving member information:", err)
		return
	}

	// Use a modal to send a welcome message
	_, err := s.ChannelMessageSendComplex(welcomeChannelId, &discordgo.MessageSend{
		Content: fmt.Sprintf("Welcome to the server, %s!", m.User.Username),
		Embed: &discordgo.MessageEmbed{
			Title:       "Welcome!",
			Description: fmt.Sprintf("Hello, %s! Welcome to the server.", m.User.Username),
			Color:       0x00ff00, // Green color
		},
	})
	if err != nil {
		log.Printf("Error sending welcome message: %s", err)
		return
	}

	err = s.GuildMemberRoleAdd(m.GuildID, m.User.ID, autoRoleId)
	if err != nil {
		log.Printf("Error adding role to new member: %s", err)
		return
	}
}
