package pkg

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/gpnull/golang-github.com/database"
)

func RestoreButtons(s *discordgo.Session, dbClient *database.MongoClient, timeKeepingChannelId string) {
	buttons, err := dbClient.GetButtons(context.Background())
	if err != nil {
		fmt.Println("Error retrieving buttons:", err)
		return
	}

	// Clear existing buttons from the UI
	clearExistingButtons(s, timeKeepingChannelId)

	for _, button := range buttons {
		actionRow := discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    button.Label,
					CustomID: button.ButtonID,
					Style:    button.Style,
				},
			},
		}

		s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.Type != discordgo.InteractionMessageComponent {
				return
			}

			if i.MessageComponentData().CustomID != button.ButtonID {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "wrong",
						Flags:   discordgo.MessageFlagsEphemeral,
					},
				})
				return
			}

			if i.Member.User.ID == button.ButtonID {
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
							Content:    "tạm biệt",
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
}

// Function to clear existing buttons from the UI
func clearExistingButtons(s *discordgo.Session, timeKeepingChannelId string) {
	var messagesDeleted int
	for {
		messages, err := s.ChannelMessages(timeKeepingChannelId, 100, "", "", "")
		if err != nil || len(messages) == 0 {
			break
		}

		messageIDs := make([]string, len(messages))
		for i, msg := range messages {
			messageIDs[i] = msg.ID
		}

		err = s.ChannelMessagesBulkDelete(timeKeepingChannelId, messageIDs)
		if err != nil {
			s.ChannelMessageSend(timeKeepingChannelId, "An error occurred while deleting messages.")
			return
		}

		messagesDeleted += len(messages)
	}
}
