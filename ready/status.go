package ready

import (
	"log"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Status(s *discordgo.Session, e *discordgo.Ready) {
	go func() {
		for {
			statusIndex := rand.Intn(3)
			statuses := []string{"perfect nil", "perfect null", "perfect bot"}

			s.UpdateStatusComplex(discordgo.UpdateStatusData{
				IdleSince: nil, // Set this to a time to show idle status
				Activities: []*discordgo.Activity{
					{
						Name: statuses[statusIndex],
						Type: discordgo.ActivityTypeWatching,
					},
				},
				AFK:    true,
				Status: string(discordgo.StatusOnline),
			})

			time.Sleep(3 * time.Second) // Wait for 3 seconds before changing status
		}
	}()

	log.Print(s.State.User.Username + " is online")
}
