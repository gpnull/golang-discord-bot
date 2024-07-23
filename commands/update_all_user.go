package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/gpnull/golang-github.com/database"
	"github.com/gpnull/golang-github.com/models"
	util "github.com/gpnull/golang-github.com/utils"
)

func init() {
	util.Commands["updateAllUser"] = updateAllUser
}

func updateAllUser(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if !util.HasPermissionClear(m, util.Config.UseBotID) {
		s.ChannelMessageSend(m.ChannelID, "You do not have permission to use this command.")
		return
	}

	// Connect to the database
	database.ConnectDB(util.Config.DbURL)
	defer database.CloseDB()
	dbClient := &database.Database{DB: database.DB}

	// Get the guild ID from the message
	guildID := m.GuildID

	// Get all members in the guild
	members, err := s.GuildMembers(guildID, "", 1000)
	if err != nil {
		fmt.Println("Error fetching guild members:", err)
		return
	}

	// Iterate through each member and create/update user in the database
	for _, member := range members {
		// Skip bots
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
