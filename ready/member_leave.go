package ready

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func GuildMemberLeave(s *discordgo.Session, m *discordgo.GuildMemberAdd, leaveChannelId string) {
	var err error
	_, err = s.ChannelMessageSendComplex(leaveChannelId, &discordgo.MessageSend{
		Content: fmt.Sprintf("Bye bye, %s!", m.User.Username),
		Embed: &discordgo.MessageEmbed{
			Title:       "Good Bye!",
			Description: fmt.Sprintf("Bye bye, %s!", m.User.Username),
			Color:       0xff0000,                        // Red color
			Timestamp:   time.Now().Format(time.RFC3339), // Add Timestamp
			Footer: &discordgo.MessageEmbedFooter{
				Text:    fmt.Sprintf("Left at %s", time.Now().Format("15:04:05 MST")), // Add Footer
				IconURL: "https://i.imgur.com/AfFp7pu.png",
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://i.imgur.com/AfFp7pu.png", // Add Thumbnail URL
			},
			Author: &discordgo.MessageEmbedAuthor{
				Name:    "Author Name",
				IconURL: "https://i.imgur.com/AfFp7pu.png",
				URL:     "https://discord.js.org",
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Field 1",
					Value:  "Value 1",
					Inline: true,
				},
				{
					Name:   "Field 2",
					Value:  "Value 2",
					Inline: true,
				},
			},
		},
	})
	if err != nil {
		log.Printf("Error sending goodbye message: %s", err)
		return
	}
}
