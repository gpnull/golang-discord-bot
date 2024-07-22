package pkg

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/gpnull/golang-github.com/database"
	handler "github.com/gpnull/golang-github.com/handlers"
	"gorm.io/gorm"
)

func RestoreButtons(s *discordgo.Session, dbClient *gorm.DB, timeKeepingChannelId string) {
	db := &database.Database{DB: dbClient}
	buttons, err := db.GetTimeKeepingStatusButtons()
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
		_, err = s.ChannelMessageSendComplex("1202666554144325653", &discordgo.MessageSend{
			Components: []discordgo.MessageComponent{
				actionRow,
			},
		})
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		buttonRestore := discordgo.Button{
			Label:    button.Label,
			CustomID: button.ButtonID,
			Style:    button.Style,
		}

		s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			handler.HandleTimekeepingInteraction(s, i, button.ButtonID, &buttonRestore, actionRow)
		})

		// s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// 	if i.Type != discordgo.InteractionMessageComponent {
		// 		return
		// 	}

		// 	if i.MessageComponentData().CustomID != button.ButtonID {
		// 		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		// 			Type: discordgo.InteractionResponseChannelMessageWithSource,
		// 			Data: &discordgo.InteractionResponseData{
		// 				Content: "wrong1",
		// 				Flags:   discordgo.MessageFlagsEphemeral,
		// 			},
		// 		})
		// 		return
		// 	}

		// 	if i.Member.User.ID == button.ButtonID {
		// 		if button.Style == discordgo.PrimaryButton {
		// 			button.Style = discordgo.DangerButton
		// 			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		// 				Type: discordgo.InteractionResponseUpdateMessage,
		// 				Data: &discordgo.InteractionResponseData{
		// 					Content:    "nice",
		// 					Components: []discordgo.MessageComponent{actionRow},
		// 				},
		// 			})
		// 		} else {
		// 			button.Style = discordgo.PrimaryButton
		// 			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		// 				Type: discordgo.InteractionResponseUpdateMessage,
		// 				Data: &discordgo.InteractionResponseData{
		// 					Content:    "bye",
		// 					Components: []discordgo.MessageComponent{actionRow},
		// 				},
		// 			})
		// 		}
		// 	} else {
		// 		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		// 			Type: discordgo.InteractionResponseChannelMessageWithSource,
		// 			Data: &discordgo.InteractionResponseData{
		// 				Content: "wrong2",
		// 				Flags:   discordgo.MessageFlagsEphemeral,
		// 			},
		// 		})
		// 	}
		// })
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
