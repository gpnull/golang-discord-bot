package ready

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func GuildMemberLeave(s *discordgo.Session, m *discordgo.GuildMemberAdd, leaveChannelId string) {
	var err error
	// _, err = s.ChannelMessageSend(leaveChannelId, fmt.Sprintf("Bye bye, %s!", m.User.Username))
	_, err = s.ChannelMessageSendComplex(leaveChannelId, &discordgo.MessageSend{
		Content: fmt.Sprintf("Bye bye, %s!", m.User.Username),
		Embed: &discordgo.MessageEmbed{
			Title:       "Good Bye!",
			Description: fmt.Sprintf("Bye bye, %s!", m.User.Username),
			Color:       0xff0000, // Red color
		},
	})
	if err != nil {
		log.Printf("Error sending goodbye message: %s", err)
		return
	}
}
