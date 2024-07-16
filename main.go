package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	_ "github.com/gpnull/golang-github.com/commands"

	"github.com/bwmarrin/discordgo"
	"github.com/gpnull/golang-github.com/util"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	s, err := discordgo.New("Bot " + util.Config.Token)
	if err != nil {
		log.Panicf("Error creating session: %s", err)
	}

	s.AddHandler(ready)
	s.AddHandler(messageCreate)
	s.AddHandler(guildMemberAdd) // Add handler for guild member add event

	s.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages | discordgo.IntentsGuilds | discordgo.IntentsGuildMembers) // Add GuildMembers intent

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

func ready(s *discordgo.Session, e *discordgo.Ready) {
	s.UpdateStatusComplex(discordgo.UpdateStatusData{
		Activities: []*discordgo.Activity{
			{
				Name: "I am a cool bot",
				Type: discordgo.ActivityTypeGame,
			},
		},
		Status: string(discordgo.StatusOnline),
	})

	log.Print(s.State.User.Username + " is online")
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

func guildMemberAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	fmt.Println("User: \n", m)
	// Replace "WelcomeRoleID" with the actual ID of the role you want to assign
	err := s.GuildMemberRoleAdd(m.GuildID, m.User.ID, "1202667195944009779")
	if err != nil {
		log.Printf("Error adding role to new member: %s", err)
		return
	}

	// Replace "WelcomeChannelID" with the actual ID of the channel you want to send the message to
	_, err = s.ChannelMessageSend("1202666419519619166", fmt.Sprintf("Welcome to the server, %s!", m.User.Username))
	if err != nil {
		log.Printf("Error sending welcome message: %s", err)
		return
	}
}
