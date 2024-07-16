package ready

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/gpnull/golang-github.com/database"
)

func GuildMemberLeave(dbClient *database.MongoClient, s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	// Replace "WelcomeChannelID" with the actual ID of the channel you want to send the message to
	var err error
	_, err = s.ChannelMessageSend("1202666419519619166", fmt.Sprintf("Bye bye, %s!", m.User.Username))
	if err != nil {
		log.Printf("Error sending welcome message: %s", err)
		return
	}
}
