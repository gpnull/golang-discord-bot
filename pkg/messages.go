package pkg

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gpnull/golang-github.com/utils"
)

var prefix = "."

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot || !strings.HasPrefix(m.Content, prefix) {
		return
	}

	perms, err := s.State.UserChannelPermissions(s.State.User.ID, m.ChannelID)
	if err != nil {
		log.Printf("Could not get perms for channel %s: %s", m.ChannelID, err)
		return
	}

	perm := utils.IncludesPerm

	if perm(discordgo.PermissionViewChannel|discordgo.PermissionSendMessages|discordgo.PermissionEmbedLinks, perms) {
		args := strings.Split(m.Content[len(prefix):], " ")
		if cmd, ok := utils.Commands[args[0]]; ok {
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
