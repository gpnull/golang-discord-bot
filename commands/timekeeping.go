package commands

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/gpnull/golang-github.com/database"
	"github.com/gpnull/golang-github.com/models"
	util "github.com/gpnull/golang-github.com/utils"
)

func init() {
	util.Commands["createTimekeeping"] = createTimekeeping
}

func createTimekeeping(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	database.ConnectDB(util.Config.DbURL)
	defer database.CloseDB()
	dbClient := &database.Database{DB: database.DB}

	if len(args) != 2 {
		s.ChannelMessageSend(m.ChannelID, "Usage: .createTimekeeping <name_of_register> <id_of_register>")
		return
	}

	buttonName := args[0] // name_of_register
	buttonID := args[1]   // id_of_register

	// Create new channel
	categoryID := "1202666419519619164"
	timekeeping_channel, err := s.GuildChannelCreateComplex(m.GuildID, discordgo.GuildChannelCreateData{
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

	err = dbClient.UpdateUserChannelID(context.Background(), buttonID, timekeeping_channel.ID)
	if err != nil {
		fmt.Println("Error updating user channel ID:", err)
		return
	}

	button := discordgo.Button{
		Label:    buttonName,
		CustomID: buttonID,
		Style:    discordgo.PrimaryButton,
	}

	actionRow := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{button},
	}

	_, err = s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
		// Content: "Click the button below:",
		Components: []discordgo.MessageComponent{
			actionRow,
		},
	})
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

	timekeepingStatus := &models.TimekeepingStatus{
		ButtonID: buttonID,
		Label:    buttonName,
		Style:    discordgo.PrimaryButton,
	}

	err = dbClient.SaveTimeKeepingStatusButton(timekeepingStatus)
	if err != nil {
		fmt.Println("Error saving button information:", err)
		return
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionMessageComponent {
			return
		}

		if i.MessageComponentData().CustomID != buttonID {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "wrong",
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
			return
		}

		if i.Member.User.ID == buttonID {
			if button.Style == discordgo.PrimaryButton {
				button.Style = discordgo.DangerButton
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseUpdateMessage,
					Data: &discordgo.InteractionResponseData{
						Content:    "nice",
						Components: []discordgo.MessageComponent{actionRow},
					},
				})
			} else {
				button.Style = discordgo.PrimaryButton
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseUpdateMessage,
					Data: &discordgo.InteractionResponseData{
						Content:    "bye",
						Components: []discordgo.MessageComponent{actionRow},
					},
				})
			}
		} else {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "wrong",
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
		}
	})
}
