package commands

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/gpnull/golang-github.com/database"
	"github.com/gpnull/golang-github.com/models"
	"github.com/gpnull/golang-github.com/pkg"
	"github.com/gpnull/golang-github.com/utils"
)

func init() {
	utils.Commands["createTimekeeping"] = createTimekeeping
}

// .createTimekeeping name id_discord time_start time_end

func createTimekeeping(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	dbClient := &database.Database{DB: database.DB}

	if len(args) != 4 {
		s.ChannelMessageSend(m.ChannelID, "Usage: .createTimekeeping <name_of_register> <id_of_register>")
		return
	}

	buttonName := args[0] // name_of_register
	buttonID := args[1]   // id_of_register
	timeStart, err := strconv.Atoi(args[2])
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Invalid time_start value")
		return
	}

	timeEnd, err := strconv.Atoi(args[3])
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Invalid time_end value")
		return
	}

	if timeEnd-timeStart != 2 || timeEnd%2 != 0 || timeStart%2 != 0 {
		s.ChannelMessageSend(m.ChannelID, "Invalid shift schedule")
		return
	}

	// Create new channel
	categoryID := utils.Config.TimekeepingLogCategoryID
	topic := buttonID + " | Schedule: " + strconv.Itoa(timeStart) + "-" + strconv.Itoa(timeEnd)
	timekeeping_channel_log, err := s.GuildChannelCreateComplex(m.GuildID, discordgo.GuildChannelCreateData{
		Name:     buttonName,
		Type:     discordgo.ChannelTypeGuildText,
		ParentID: categoryID,
		Topic:    topic,
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
		TimekeepingChannelID:    utils.Config.TimekeepingChannelID,
		TimekeepingLogChannelID: timekeeping_channel_log.ID,
		Status:                  utils.STOPPED,
		TimeStart:               timeStart,
		TimeEnd:                 timeEnd,
	}

	err = dbClient.SaveTimeKeepingStatusButton(timekeepingStatus)
	if err != nil {
		fmt.Println("Error saving button information:", err)
		return
	}

	pkg.RestoreButtons(s, database.DB, utils.Config.TimekeepingChannelID)
}
