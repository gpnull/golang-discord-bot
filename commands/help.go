package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gpnull/golang-github.com/utils"
)

func init() {
	utils.Commands["help"] = help
}

func help(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	msg := "help"

	_, err := s.ChannelMessageSendReply(m.ChannelID, msg, m.Reference())
	if err != nil {
		return
	}
}
