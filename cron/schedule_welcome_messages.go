package cron

import (
	"time"

	"github.com/bwmarrin/discordgo"
	cron "github.com/robfig/cron/v3"
)

func ScheduleWelcomeMessages(s *discordgo.Session, welcomeChannelId string) {
	// Create a new cron instance
	c := cron.New(cron.WithLocation(time.FixedZone("GMT+7", 7*60*60)))

	// Schedule the job to run at 0 2 4 6 8 10 12 14 16 18 20 22 GMT+7 daily

	c.AddFunc("53 10 * * *", func() {
		// s.ChannelMessageSend(util.Config.TimekeepingChannelID, "Welcome!")
		// handlers.HandleResetTimekeepingStatus(s)
		// pkg.RestoreButtons(s, database.DB, util.Config.TimekeepingChannelID)
	})

	c.AddFunc("00 02 * * *", func() {
		s.ChannelMessageSend(welcomeChannelId, "Welcome!")
	})

	c.AddFunc("00 04 * * *", func() {
		s.ChannelMessageSend(welcomeChannelId, "Welcome!")
	})

	c.AddFunc("00 06 * * *", func() {
		s.ChannelMessageSend(welcomeChannelId, "Welcome!")
	})

	c.AddFunc("00 08 * * *", func() {
		s.ChannelMessageSend(welcomeChannelId, "Welcome!")
	})

	c.AddFunc("00 10 * * *", func() {
		s.ChannelMessageSend(welcomeChannelId, "Welcome!")
	})

	c.AddFunc("00 12 * * *", func() {
		s.ChannelMessageSend(welcomeChannelId, "Welcome!")
	})

	c.AddFunc("00 14 * * *", func() {
		s.ChannelMessageSend(welcomeChannelId, "Welcome!")
	})

	c.AddFunc("00 16 * * *", func() {
		s.ChannelMessageSend(welcomeChannelId, "Welcome!")
	})

	c.AddFunc("00 18 * * *", func() {
		s.ChannelMessageSend(welcomeChannelId, "Welcome!")
	})

	c.AddFunc("00 20 * * *", func() {
		s.ChannelMessageSend(welcomeChannelId, "Welcome!")
	})

	c.AddFunc("00 22 * * *", func() {
		s.ChannelMessageSend(welcomeChannelId, "Welcome!")
	})

	// Start the cron scheduler
	c.Start()
}
