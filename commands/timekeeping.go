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
	if len(args) != 2 {
		s.ChannelMessageSend(m.ChannelID, "Usage: .createTimekeeping <name_of_register> <id_of_register>")
		return
	}

	buttonName := args[0] // name_of_register
	buttonID := args[1]   // id_of_register

	// Create new channel
	categoryID := "1202666419519619164"
	_, err := s.GuildChannelCreateComplex(m.GuildID, discordgo.GuildChannelCreateData{
		Name:     buttonName,
		Type:     discordgo.ChannelTypeGuildText,
		ParentID: categoryID,
		Topic:    buttonID,
	})
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Cannot create new channel: "+err.Error())
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

	// Save button info into MongoDB
	dbClient, err := database.ConnectDB(util.Config.MongoURI)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	defer dbClient.DisconnectDB(context.Background())

	timekeeping := &models.Timekeeping{
		ID:       buttonID,
		ButtonID: buttonID,
		Label:    buttonName,
		Style:    discordgo.PrimaryButton,
	}

	err = dbClient.SaveButton(context.Background(), timekeeping)
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
					Content: "không đúng",
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
						Content:    "cảm ơn",
						Components: []discordgo.MessageComponent{actionRow},
					},
				})
			} else {
				button.Style = discordgo.PrimaryButton
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseUpdateMessage,
					Data: &discordgo.InteractionResponseData{
						Content:    "tạm biệt",
						Components: []discordgo.MessageComponent{actionRow},
					},
				})
			}
		} else {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "không đúng",
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
		}
	})
}
