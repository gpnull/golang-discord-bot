package ready

import (
	"context"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/gpnull/golang-github.com/database"
	"github.com/gpnull/golang-github.com/models"
)

func GuildMemberAdd(dbClient *database.MongoClient, s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	user := models.User{
		DiscordID: m.User.ID,
		Username:  m.User.Username,
	}

	if err := dbClient.CreateUser(context.Background(), &user); err != nil {
		fmt.Println("Error saving member information:", err)
		return
	}

	// Replace "WelcomeRoleID" with the actual ID of the role you want to assign
	err := s.GuildMemberRoleAdd(m.GuildID, m.User.ID, "1202667195944009779")
	if err != nil {
		log.Printf("Error adding role to new member: %s", err)
		return
	}

	// Replace "WelcomeChannelID" with the actual ID of the channel you want to send the message to
	_, err = s.ChannelMessageSend("1202666419519619166", fmt.Sprintf("Welcome to the server, %s!", m.User.Username))
	if err != nil {
		log.Printf("Error sending welcome message: %s", err)
		return
	}
}
