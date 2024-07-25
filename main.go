package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/gpnull/golang-github.com/commands"
	"github.com/gpnull/golang-github.com/cron"
	"github.com/gpnull/golang-github.com/database"
	"github.com/gpnull/golang-github.com/pkg"
	"github.com/gpnull/golang-github.com/ready"

	"github.com/bwmarrin/discordgo"
	"github.com/gpnull/golang-github.com/utils"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	s, err := discordgo.New("Bot " + utils.Config.Token)
	if err != nil {
		log.Panicf("Error creating session: %s", err)
	}

	database.ConnectDB(utils.Config.DbURL)
	defer database.CloseDB()
	database.Migrate()
	dbClient := database.DB

	// Restore previously created buttons
	pkg.RestoreButtons(s, dbClient, utils.Config.TimekeepingChannelID)

	s.AddHandler(func(s *discordgo.Session, e *discordgo.Ready) {
		ready.Status(s, e)
	})
	s.AddHandler(func(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
		ready.GuildMemberAdd(s, m, dbClient, utils.Config.AutoRoleId, utils.Config.WelcomeChannelID)
	})
	s.AddHandler(func(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
		ready.GuildMemberLeave(s, m, utils.Config.LeaveChannelID)
	})
	s.AddHandler(pkg.MessageCreate)

	s.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages |
		discordgo.IntentsGuilds | discordgo.IntentsGuildMembers) // GuildMembers intent

	err = s.Open()
	if err != nil {
		log.Panicf("Unable to open session: %s", err)
	}

	cron.ResetTimekeepingStatus(s)

	// Wait for os terminate events, cleanly close connection when encountered
	closeChan := make(chan os.Signal, 1)
	signal.Notify(closeChan, syscall.SIGTERM, syscall.SIGINT, os.Interrupt, syscall.SIGTERM)
	<-closeChan
	log.Print("OS termination received, closing WS and DB")
	s.Close()
	log.Print("Connections closed, bye bye")
}
