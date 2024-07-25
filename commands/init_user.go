package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/gpnull/golang-github.com/database"
	"github.com/gpnull/golang-github.com/models"
	"github.com/gpnull/golang-github.com/utils"
)

func init() {
	utils.Commands["initUsers"] = initUsers
}

func initUsers(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if !utils.HasPermissionClear(m, utils.Config.UseBotID) {
		s.ChannelMessageSend(m.ChannelID, "You do not have permission to use this command.")
		return
	}

	dbClient := &database.Database{DB: database.DB}

	guildID := m.GuildID
	members, err := s.GuildMembers(guildID, "", 1000)
	if err != nil {
		fmt.Println("Error fetching guild members:", err)
		return
	}

	for _, member := range members {
		if member.User.Bot {
			continue
		}

		user := models.User{
			DiscordID:            member.User.ID,
			Username:             member.User.Username,
			Email:                member.User.Email,
			Avatar:               member.User.Avatar,
			Locale:               member.User.Locale,
			GlobalName:           member.User.GlobalName,
			Verified:             member.User.Verified,
			Banner:               member.User.Banner,
			TimekeepingChannelID: "",
		}

		if err := dbClient.CreateUser(&user); err != nil {
			fmt.Println("Error saving member information:", err)
		}
	}

	// Send a confirmation message
	s.ChannelMessageSend(m.ChannelID, "All users have been updated.")
}
