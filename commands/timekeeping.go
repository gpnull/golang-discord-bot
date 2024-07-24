package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/gpnull/golang-github.com/database"
	"github.com/gpnull/golang-github.com/models"
	"github.com/gpnull/golang-github.com/pkg"
	util "github.com/gpnull/golang-github.com/utils"
)

func init() {
	util.Commands["createTimekeeping"] = createTimekeeping
}

func createTimekeeping(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	dbClient := &database.Database{DB: database.DB}

	if len(args) != 2 {
		s.ChannelMessageSend(m.ChannelID, "Usage: .createTimekeeping <name_of_register> <id_of_register>")
		return
	}

	buttonName := args[0] // name_of_register
	buttonID := args[1]   // id_of_register

	// Create new channel
	categoryID := util.Config.TimekeepingLogCategoryID
	timekeeping_channel_log, err := s.GuildChannelCreateComplex(m.GuildID, discordgo.GuildChannelCreateData{
		Name:     buttonName,
		Type:     discordgo.ChannelTypeGuildText,
		ParentID: categoryID,
		Topic:    buttonID,
		PermissionOverwrites: []*discordgo.PermissionOverwrite{
			{
				ID:    args[1], // User ID
				Type:  discordgo.PermissionOverwriteTypeMember,
				Allow: discordgo.PermissionViewChannel,
				Deny:  discordgo.PermissionManageChannels | discordgo.PermissionSendMessages, // Deny permission to manage channels
			},
			{
				ID:   m.GuildID, // Default role (everyone)
				Type: discordgo.PermissionOverwriteTypeRole,
				Deny: discordgo.PermissionViewChannel,
			},
		},
	})
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Cannot create new channel: "+err.Error())
		return
	}

	err = dbClient.UpdateUserChannelID(buttonID, timekeeping_channel_log.ID)
	if err != nil {
		fmt.Println("Error updating user channel ID:", err)
		return
	}

	timekeepingStatus := &models.TimekeepingStatus{
		ButtonID:                buttonID,
		Label:                   buttonName,
		Style:                   discordgo.SecondaryButton,
		Content:                 "Welcome...",
		TimekeepingChannelID:    util.Config.TimekeepingChannelID,
		TimekeepingLogChannelID: timekeeping_channel_log.ID,
		Status:                  util.STOPPED,
	}

	err = dbClient.SaveTimeKeepingStatusButton(timekeepingStatus)
	if err != nil {
		fmt.Println("Error saving button information:", err)
		return
	}

	pkg.RestoreButtons(s, database.DB, util.Config.TimekeepingChannelID)
}
