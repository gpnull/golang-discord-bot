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

		channelID := db.GetChannelIDByDiscordID(button.ButtonID)
		if channelID == "" {
			fmt.Println("Error: Channel ID not found for user")
			return
		}

		s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.MessageComponentData().CustomID == button.ButtonID {
				handler.HandleTimekeepingInteraction(s, i, button.ButtonID, &buttonRestore, actionRow, channelID)
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
