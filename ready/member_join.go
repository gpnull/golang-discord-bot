package ready

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gpnull/golang-github.com/database"
	"github.com/gpnull/golang-github.com/models"
)

func GuildMemberAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd,
	dbClient *database.MongoClient, autoRoleId, welcomeChannelId string) {
	user := models.User{
		DiscordID:            m.User.ID,
		Username:             m.User.Username,
		Email:                m.User.Email,
		Avatar:               m.User.Avatar,
		Locale:               m.User.Locale,
		GlobalName:           m.User.GlobalName,
		Verified:             m.User.Verified,
		Banner:               m.User.Banner,
		TimekeepingChannelID: "",
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
			Color:       0x00ff00,                        // Green color
			Timestamp:   time.Now().Format(time.RFC3339), // Add Timestamp
			Footer: &discordgo.MessageEmbedFooter{
				Text:    fmt.Sprintf("Join at %s", time.Now().Format("15:04:05 MST")), // Add Footer
				IconURL: "https://i.imgur.com/AfFp7pu.png",
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://i.imgur.com/AfFp7pu.png", // Add Thumbnail URL
			},
			Author: &discordgo.MessageEmbedAuthor{
				Name:    "Author Name",
				IconURL: "https://i.imgur.com/AfFp7pu.png",
				URL:     "https://discord.js.org",
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Field 1",
					Value:  "Value 1",
					Inline: true,
				},
				{
					Name:   "Field 2",
					Value:  "Value 2",
					Inline: true,
				},
			},
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
