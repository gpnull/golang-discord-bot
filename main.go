package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	_ "github.com/gpnull/golang-github.com/commands"
	"github.com/gpnull/golang-github.com/database"
	"github.com/gpnull/golang-github.com/ready"

	"github.com/bwmarrin/discordgo"
	util "github.com/gpnull/golang-github.com/utils"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	s, err := discordgo.New("Bot " + util.Config.Token)
	if err != nil {
		log.Panicf("Error creating session: %s", err)
	}

	// Connect to MongoDB
	dbClient, err := database.ConnectDB(util.Config.MongoURI)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	defer dbClient.DisconnectDB(context.Background())

	s.AddHandler(func(s *discordgo.Session, e *discordgo.Ready) {
		ready.Status(s, e)
	})
	s.AddHandler(func(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
		ready.GuildMemberAdd(s, m, dbClient, util.Config.AutoRoleId, util.Config.WelcomeChannelID)
	})
	s.AddHandler(func(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
		ready.GuildMemberLeave(s, m, util.Config.LeaveChannelID)
	})
	s.AddHandler(messageCreate)

	s.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages |
		discordgo.IntentsGuilds | discordgo.IntentsGuildMembers) // GuildMembers intent

	err = s.Open()
	if err != nil {
		log.Panicf("Unable to open session: %s", err)
	}

	// Wait for os terminate events, cleanly close connection when encountered
	closeChan := make(chan os.Signal, 1)
	signal.Notify(closeChan, syscall.SIGTERM, syscall.SIGINT, os.Interrupt, syscall.SIGTERM)
	<-closeChan
	log.Print("OS termination received, closing WS and DB")
	s.Close()
	log.Print("Connections closed, bye bye")
}

var prefix = "."

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot || !strings.HasPrefix(m.Content, prefix) {
		return
	}

	perms, err := s.State.UserChannelPermissions(s.State.User.ID, m.ChannelID)
	if err != nil {
		log.Printf("Could not get perms for channel %s: %s", m.ChannelID, err)
		return
	}

	// Check user's role
	hasRole := false
	for _, roleID := range m.Member.Roles {
		if roleID == util.Config.UseBotID {
			hasRole = true
			break
		}
	}

	if !hasRole {
		s.ChannelMessageSend(m.ChannelID, "You do not have permission to use this command.")
		return
	}

	perm := util.IncludesPerm

	if perm(discordgo.PermissionViewChannel|discordgo.PermissionSendMessages|discordgo.PermissionEmbedLinks, perms) {
		args := strings.Split(m.Content[len(prefix):], " ")
		if cmd, ok := util.Commands[args[0]]; ok {
			if len(args) == 1 {
				args = []string{}
			} else {
				args = args[1:]
			}

			cmd(s, m, args)
		}
	} else if perm(discordgo.PermissionSendMessages, perms) {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("I seem to be missing permissions. Below, false indicates a lacking permission. Please grant these permissions on my role.\n```"+
			"Read Text Channels & See Voice Channels: %t\nSend Messages: true\nEmbed Links: %t```",
			perm(discordgo.PermissionViewChannel, perms), perm(discordgo.PermissionEmbedLinks, perms)))
	}
}
