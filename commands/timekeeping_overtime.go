package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/gpnull/golang-github.com/database"
	"github.com/gpnull/golang-github.com/models"
	"github.com/gpnull/golang-github.com/pkg"
	"github.com/gpnull/golang-github.com/utils"
)

func init() {
	utils.Commands["regtkot"] = createTimekeepingOvertime
}

func createTimekeepingOvertime(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	dbClient := &database.Database{DB: database.DB}

	if len(args) != 2 {
		s.ChannelMessageSend(m.ChannelID,
			"Usage: .regtkot <name_of_register> <id_of_register>")
		return
	}

	buttonName := args[0] // name_of_register
	buttonID := args[1]   // id_of_register

	// Create new channel
	categoryID := utils.Config.TimekeepingOvertimeLogCategoryID
	topic := buttonID
	timekeeping_overtime_channel_log, err := s.GuildChannelCreateComplex(m.GuildID,
		discordgo.GuildChannelCreateData{
			Name:     buttonName,
			Type:     discordgo.ChannelTypeGuildText,
			ParentID: categoryID,
			Topic:    topic,
			PermissionOverwrites: []*discordgo.PermissionOverwrite{
				{
					ID:    args[1], // User ID
					Type:  discordgo.PermissionOverwriteTypeMember,
					Allow: discordgo.PermissionViewChannel,
					Deny: discordgo.PermissionManageChannels |
						discordgo.PermissionSendMessages, // Deny permission to manage channels
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

	err = dbClient.UpdateTimekeepingOvertimeChannelID(buttonID, timekeeping_overtime_channel_log.ID)
	if err != nil {
		fmt.Println("Error updating UpdateTimekeepingOvertimeChannelID func:", err)
		return
	}

	timekeepingOvertimeStatus := &models.TimekeepingOvertimeStatus{
		ButtonID:                        buttonID,
		Label:                           buttonName,
		Style:                           discordgo.SecondaryButton,
		Content:                         "Welcome...",
		TimekeepingOvertimeChannelID:    utils.Config.TimekeepingChannelID,
		TimekeepingOvertimeLogChannelID: timekeeping_overtime_channel_log.ID,
		Status:                          utils.STOPPED,
	}

	err = dbClient.SaveTimeKeepingOvertimeStatusButton(timekeepingOvertimeStatus)
	if err != nil {
		fmt.Println("Error saving button information:", err)
		return
	}

	pkg.RestoreButtons(s, database.DB, utils.Config.TimekeepingOvertimeChannelID)
}
