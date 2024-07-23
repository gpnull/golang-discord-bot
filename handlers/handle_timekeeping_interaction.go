package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func HandleTimekeepingInteraction(s *discordgo.Session, i *discordgo.InteractionCreate,
	buttonID string, button *discordgo.Button, actionRow discordgo.ActionsRow) {
	if i.Type != discordgo.InteractionMessageComponent {
		return
	}

	if i.MessageComponentData().CustomID != buttonID {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "wrong1",
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
				Content: "wrong2",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
	}
}
