package cron

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gpnull/golang-github.com/database"
	"github.com/gpnull/golang-github.com/handlers"
	"github.com/gpnull/golang-github.com/pkg"
	util "github.com/gpnull/golang-github.com/utils"
	cron "github.com/robfig/cron/v3"
)

func ResetTimekeepingStatus(s *discordgo.Session) {
	// Create a new cron instance
	c := cron.New(cron.WithLocation(time.FixedZone("GMT+7", 7*60*60)))

	// Schedule the job to run at 0 2 4 6 8 10 12 14 16 18 20 22 GMT+7 daily
	c.AddFunc("25 16 * * *", func() {
		s.ChannelMessageSend(util.Config.TimekeepingChannelID, "Welcome!")
		handlers.HandleResetTimekeepingStatus(s)
		pkg.RestoreButtons(s, database.DB, util.Config.TimekeepingChannelID)
	})

	c.AddFunc("00 02 * * *", func() {
		s.ChannelMessageSend(util.Config.TimekeepingChannelID, "Welcome!")
	})

	c.AddFunc("00 04 * * *", func() {
		s.ChannelMessageSend(util.Config.TimekeepingChannelID, "Welcome!")
	})

	c.AddFunc("00 06 * * *", func() {
		s.ChannelMessageSend(util.Config.TimekeepingChannelID, "Welcome!")
	})

	c.AddFunc("00 08 * * *", func() {
		s.ChannelMessageSend(util.Config.TimekeepingChannelID, "Welcome!")
	})

	c.AddFunc("00 10 * * *", func() {
		s.ChannelMessageSend(util.Config.TimekeepingChannelID, "Welcome!")
	})

	c.AddFunc("00 12 * * *", func() {
		s.ChannelMessageSend(util.Config.TimekeepingChannelID, "Welcome!")
	})

	c.AddFunc("00 14 * * *", func() {
		s.ChannelMessageSend(util.Config.TimekeepingChannelID, "Welcome!")
	})

	c.AddFunc("00 16 * * *", func() {
		s.ChannelMessageSend(util.Config.TimekeepingChannelID, "Welcome!")
	})

	c.AddFunc("00 18 * * *", func() {
		s.ChannelMessageSend(util.Config.TimekeepingChannelID, "Welcome!")
	})

	c.AddFunc("00 20 * * *", func() {
		s.ChannelMessageSend(util.Config.TimekeepingChannelID, "Welcome!")
	})

	c.AddFunc("00 22 * * *", func() {
		s.ChannelMessageSend(util.Config.TimekeepingChannelID, "Welcome!")
	})

	// Start the cron scheduler
	c.Start()
}
