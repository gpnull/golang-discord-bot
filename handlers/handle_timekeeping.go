package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/gpnull/golang-github.com/database"
	"github.com/gpnull/golang-github.com/models"
	"github.com/gpnull/golang-github.com/utils"
)

func HandleTimekeepingInteraction(s *discordgo.Session, i *discordgo.InteractionCreate,
	buttonID string, button *discordgo.Button, actionRow discordgo.ActionsRow, channelID string) {
	if i.Type != discordgo.InteractionMessageComponent {
		return
	}

	if i.Member.User.ID == buttonID {
		dbClient := &database.Database{DB: database.DB}

		buttonResult, err := dbClient.GetTimeKeepingStatusButtonByID(buttonID)
		if err != nil {
			fmt.Println("Error retrieving button:", err)
			return
		}

		timekeepingStatus := &models.TimekeepingStatus{
			ButtonID:                buttonResult.ButtonID,
			Label:                   buttonResult.Label,
			Style:                   buttonResult.Style,
			Content:                 buttonResult.Content,
			TimekeepingChannelID:    buttonResult.TimekeepingChannelID,
			TimekeepingLogChannelID: buttonResult.TimekeepingLogChannelID,
			Status:                  buttonResult.Status,
			TimeStart:               buttonResult.TimeStart,
			TimeEnd:                 buttonResult.TimeEnd,
		}

		now := utils.GetDayTimeNow()
		var content string
		var status string
		if button.Style == discordgo.SecondaryButton {
			button.Style = discordgo.SuccessButton
			timekeepingStart(s, now, channelID)
			content = "Work has started."
			status = utils.WORKING
		} else {
			button.Style = discordgo.SecondaryButton
			timekeepingEnd(s, now, channelID)
			content = "Work has ended."
			status = utils.STOPPED
		}

		// Update the button style in the actionRow
		actionRow.Components[0] = button

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Content:    content,
				Components: []discordgo.MessageComponent{actionRow},
			},
		})

		timekeepingStatus.Style = button.Style
		timekeepingStatus.Content = content
		timekeepingStatus.Status = status

		// Save the timekeeping status
		err = dbClient.SaveTimeKeepingStatusButton(timekeepingStatus)
		if err != nil {
			fmt.Println("Error saving button information:", err)
		}
	} else {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Invalid operation, please try again",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
	}
}

func HandleTimekeepingOvertimeInteraction(s *discordgo.Session, i *discordgo.InteractionCreate,
	buttonID string, button *discordgo.Button, actionRow discordgo.ActionsRow, channelID string) {
	if i.Type != discordgo.InteractionMessageComponent {
		return
	}

	if i.Member.User.ID == buttonID {
		dbClient := &database.Database{DB: database.DB}

		buttonResult, err := dbClient.GetTimeKeepingOvertimeStatusButtonByID(buttonID)
		if err != nil {
			fmt.Println("Error retrieving button:", err)
			return
		}

		timekeepingOvertimeStatus := &models.TimekeepingOvertimeStatus{
			ButtonID:                        buttonResult.ButtonID,
			Label:                           buttonResult.Label,
			Style:                           buttonResult.Style,
			Content:                         buttonResult.Content,
			TimekeepingOvertimeChannelID:    buttonResult.TimekeepingOvertimeChannelID,
			TimekeepingOvertimeLogChannelID: buttonResult.TimekeepingOvertimeLogChannelID,
			Status:                          buttonResult.Status,
		}

		now := utils.GetDayTimeNow()
		var content string
		var status string
		if button.Style == discordgo.SecondaryButton {
			button.Style = discordgo.SuccessButton
			timekeepingStart(s, now, channelID)
			content = "Work has started."
			status = utils.WORKING
		} else {
			button.Style = discordgo.SecondaryButton
			timekeepingEnd(s, now, channelID)
			content = "Work has ended."
			status = utils.STOPPED
		}

		// Update the button style in the actionRow
		actionRow.Components[0] = button

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Content:    content,
				Components: []discordgo.MessageComponent{actionRow},
			},
		})

		timekeepingOvertimeStatus.Style = button.Style
		timekeepingOvertimeStatus.Content = content
		timekeepingOvertimeStatus.Status = status

		// Save the timekeeping status
		err = dbClient.SaveTimeKeepingOvertimeStatusButton(timekeepingOvertimeStatus)
		if err != nil {
			fmt.Println("Error saving button information:", err)
		}
	} else {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Invalid operation, please try again",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
	}
}

func timekeepingStart(s *discordgo.Session, nowTime, channelId string) {
	message := fmt.Sprintf("🚑\n🩺 Work has started at: %s", nowTime)
	_, err := s.ChannelMessageSend(channelId, message)
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}

func timekeepingEnd(s *discordgo.Session, nowTime, channelId string) {
	message := fmt.Sprintf("💤 Work has ended at: %s", nowTime)
	_, err := s.ChannelMessageSend(channelId, message)
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}

func HandleResetTimekeepingStatus(s *discordgo.Session) {
	dbClient := &database.Database{DB: database.DB}

	buttons, err := dbClient.GetTimeKeepingStatusButtons()
	if err != nil {
		fmt.Println("Error retrieving buttons:", err)
		return
	}

	now := utils.GetDayTimeNow()

	for _, button := range buttons {
		if button.Status == utils.WORKING {
			button.Status = utils.STOPPED
			button.Style = discordgo.SecondaryButton
			button.Content = "Work has ended."

			err := dbClient.SaveTimeKeepingStatusButton(button)
			if err != nil {
				fmt.Println("Error updating button information:", err)
			} else {
				timekeepingEnd(s, now, button.TimekeepingLogChannelID)
			}
		}
	}
}
