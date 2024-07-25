package cron

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gpnull/golang-github.com/database"
	"github.com/gpnull/golang-github.com/handlers"
	"github.com/gpnull/golang-github.com/pkg"
	"github.com/gpnull/golang-github.com/utils"
	cron "github.com/robfig/cron/v3"
)

func ResetTimekeepingStatus(s *discordgo.Session) {
	// Create a new cron instance
	c := cron.New(cron.WithLocation(time.FixedZone("GMT+7", 7*60*60)))

	// Function to add a cron job
	addCronJob := func(spec string) {
		c.AddFunc(spec, func() {
			handlers.HandleResetTimekeepingStatus(s)
			pkg.RestoreButtons(s, database.DB, utils.Config.TimekeepingChannelID)
		})
	}

	// Schedule the jobs
	for hour := 0; hour < 24; hour += 2 {
		addCronJob(fmt.Sprintf("00 %02d * * *", hour))
	}

	// Start the cron scheduler
	c.Start()
}

// func ResetTimekeepingStatus(s *discordgo.Session) {
// 	// Create a new cron instance
// 	c := cron.New(cron.WithLocation(time.FixedZone("GMT+7", 7*60*60)))

// 	// Schedule the job to run at 0 2 4 6 8 10 12 14 16 18 20 22 GMT+7 daily
// 	c.AddFunc("00 00 * * *", func() {
// 		handlers.HandleResetTimekeepingStatus(s)
// 		pkg.RestoreButtons(s, database.DB, util.Config.TimekeepingChannelID)
// 	})

// 	c.AddFunc("00 02 * * *", func() {
// 		handlers.HandleResetTimekeepingStatus(s)
// 		pkg.RestoreButtons(s, database.DB, util.Config.TimekeepingChannelID)
// 	})

// 	c.AddFunc("00 04 * * *", func() {
// 		handlers.HandleResetTimekeepingStatus(s)
// 		pkg.RestoreButtons(s, database.DB, util.Config.TimekeepingChannelID)
// 	})

// 	c.AddFunc("00 06 * * *", func() {
// 		handlers.HandleResetTimekeepingStatus(s)
// 		pkg.RestoreButtons(s, database.DB, util.Config.TimekeepingChannelID)
// 	})

// 	c.AddFunc("00 08 * * *", func() {
// 		handlers.HandleResetTimekeepingStatus(s)
// 		pkg.RestoreButtons(s, database.DB, util.Config.TimekeepingChannelID)
// 	})

// 	c.AddFunc("00 10 * * *", func() {
// 		handlers.HandleResetTimekeepingStatus(s)
// 		pkg.RestoreButtons(s, database.DB, util.Config.TimekeepingChannelID)
// 	})

// 	c.AddFunc("00 12 * * *", func() {
// 		handlers.HandleResetTimekeepingStatus(s)
// 		pkg.RestoreButtons(s, database.DB, util.Config.TimekeepingChannelID)
// 	})

// 	c.AddFunc("00 14 * * *", func() {
// 		handlers.HandleResetTimekeepingStatus(s)
// 		pkg.RestoreButtons(s, database.DB, util.Config.TimekeepingChannelID)
// 	})

// 	c.AddFunc("00 16 * * *", func() {
// 		handlers.HandleResetTimekeepingStatus(s)
// 		pkg.RestoreButtons(s, database.DB, util.Config.TimekeepingChannelID)
// 	})

// 	c.AddFunc("00 18 * * *", func() {
// 		handlers.HandleResetTimekeepingStatus(s)
// 		pkg.RestoreButtons(s, database.DB, util.Config.TimekeepingChannelID)
// 	})

// 	c.AddFunc("00 20 * * *", func() {
// 		handlers.HandleResetTimekeepingStatus(s)
// 		pkg.RestoreButtons(s, database.DB, util.Config.TimekeepingChannelID)
// 	})

// 	c.AddFunc("00 22 * * *", func() {
// 		handlers.HandleResetTimekeepingStatus(s)
// 		pkg.RestoreButtons(s, database.DB, util.Config.TimekeepingChannelID)
// 	})

// 	// Start the cron scheduler
// 	c.Start()
// }
