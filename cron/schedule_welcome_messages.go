package cron

import (
	"time"

	"github.com/bwmarrin/discordgo"
	cron "github.com/robfig/cron/v3"
)

func ScheduleWelcomeMessages(s *discordgo.Session, welcomeChannelId string) {
	// Create a new cron instance
	c := cron.New(cron.WithLocation(time.FixedZone("GMT+7", 7*60*60)))

	// Schedule the job to run at 15:20 GMT+7 daily
	c.AddFunc("58 15 * * *", func() {
		s.ChannelMessageSend(welcomeChannelId, "Welcome!")
	})

	// Start the cron scheduler
	c.Start()
}
