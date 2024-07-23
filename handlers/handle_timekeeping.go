package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/gpnull/golang-github.com/utils"
)

func HandleTimekeepingInteraction(s *discordgo.Session, i *discordgo.InteractionCreate,
	buttonID string, button *discordgo.Button, actionRow discordgo.ActionsRow, channelID string) {
	if i.Type != discordgo.InteractionMessageComponent {
		return
	}

	now := utils.GetDayTimeNow()

	if i.Member.User.ID == buttonID {
		if button.Style == discordgo.PrimaryButton {
			button.Style = discordgo.DangerButton
			timekeepingStart(s, now, channelID)

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
				Data: &discordgo.InteractionResponseData{
					Content:    "Work has started.",
					Components: []discordgo.MessageComponent{actionRow},
				},
			})
		} else {
			button.Style = discordgo.PrimaryButton
			timekeepingEnd(s, now, channelID)

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
				Data: &discordgo.InteractionResponseData{
					Content:    "Work has ended.",
					Components: []discordgo.MessageComponent{actionRow},
				},
			})
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
	message := fmt.Sprintf("Work has started at: %s", nowTime)
	_, err := s.ChannelMessageSend(channelId, message)
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}

func timekeepingEnd(s *discordgo.Session, nowTime, channelId string) {
	message := fmt.Sprintf("Work has ended at: %s", nowTime)
	_, err := s.ChannelMessageSend(channelId, message)
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}
