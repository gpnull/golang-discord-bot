package ready

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func Status(s *discordgo.Session, e *discordgo.Ready) {
	s.UpdateStatusComplex(discordgo.UpdateStatusData{
		Activities: []*discordgo.Activity{
			{
				Name: "I am a cool bot",
				Type: discordgo.ActivityTypeCustom,
			},
		},
		Status: string(discordgo.StatusOnline),
	})

	log.Print(s.State.User.Username + " is online")
}
