package pkg

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/gpnull/golang-github.com/database"
	handler "github.com/gpnull/golang-github.com/handlers"
	"github.com/gpnull/golang-github.com/utils"
	util "github.com/gpnull/golang-github.com/utils"
	"gorm.io/gorm"
)

func RestoreButtons(s *discordgo.Session, dbClient *gorm.DB, timeKeepingChannelId string) {
	db := &database.Database{DB: dbClient}
	buttons, err := db.GetTimeKeepingStatusButtons()
	if err != nil {
		fmt.Println("Error retrieving buttons:", err)
		return
	}

	now := util.GetDayTimeNow()
	hourNow, err := strconv.Atoi(utils.GetHourNow())
	if err != nil {
		fmt.Println("Error converting hour:", err)
		return
	}
	for _, button := range buttons {
		if hourNow < button.TimeStart || hourNow > button.TimeEnd {
			if button.Status == util.WORKING {
				button.Status = util.STOPPED
				button.Style = discordgo.SecondaryButton
				button.Content = "Work has ended."

				err := db.SaveTimeKeepingStatusButton(button)
				if err != nil {
					fmt.Println("Error updating button information:", err)
				} else {
					timekeepingEnd(s, now, button.TimekeepingLogChannelID)
				}
			}
		}
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
		_, err = s.ChannelMessageSendComplex(util.Config.TimekeepingChannelID, &discordgo.MessageSend{
			Content: button.Content,
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

func timekeepingEnd(s *discordgo.Session, nowTime, channelId string) {
	message := fmt.Sprintf("💤 Work has ended at: %s", nowTime)
	_, err := s.ChannelMessageSend(channelId, message)
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}
