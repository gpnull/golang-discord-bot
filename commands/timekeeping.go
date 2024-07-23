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
	database.ConnectDB(util.Config.DbURL)
	dbClient := &database.Database{DB: database.DB}
	defer database.CloseDB()

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
			// {
			// 	ID:    "1202667195944009779", // Role ID
			// 	Type:  discordgo.PermissionOverwriteTypeRole,
			// 	Allow: discordgo.PermissionViewChannel | discordgo.PermissionManageChannels, // Allow view and manage channels
			// },
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

	// button := discordgo.Button{
	// 	Label:    buttonName,
	// 	CustomID: buttonID,
	// 	Style:    discordgo.PrimaryButton,
	// }

	// actionRow := discordgo.ActionsRow{
	// 	Components: []discordgo.MessageComponent{button},
	// }

	// _, err = s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
	// 	Content: "Welcome...",
	// 	Components: []discordgo.MessageComponent{
	// 		actionRow,
	// 	},
	// })
	// if err != nil {
	// 	fmt.Println("Error sending message:", err)
	// 	return
	// }

	timekeepingStatus := &models.TimekeepingStatus{
		ButtonID:                buttonID,
		Label:                   buttonName,
		Style:                   discordgo.PrimaryButton,
		Content:                 "Welcome...",
		TimekeepingChannelID:    util.Config.TimekeepingChannelID,
		TimekeepingLogChannelID: timekeeping_channel_log.ID,
	}

	err = dbClient.SaveTimeKeepingStatusButton(timekeepingStatus)
	if err != nil {
		fmt.Println("Error saving button information:", err)
		return
	}

	pkg.RestoreButtons(s, database.DB, util.Config.TimekeepingChannelID)

	// s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// 	if i.MessageComponentData().CustomID == buttonID {
	// 		handler.HandleTimekeepingInteraction(s, i, buttonID, &button, actionRow, timekeeping_channel_log.ID)
	// 	}
	// })

}
