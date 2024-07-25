package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gpnull/golang-github.com/database"
	"github.com/gpnull/golang-github.com/pkg"
	"github.com/gpnull/golang-github.com/utils"
)

func init() {
	utils.Commands["restoreButtons"] = restoreButtons
}

func restoreButtons(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if !utils.HasPermissionClear(m, utils.Config.UseBotID) {
		s.ChannelMessageSend(m.ChannelID, "You do not have permission to use this command.")
		return
	}

	pkg.RestoreButtons(s, database.DB, utils.Config.TimekeepingChannelID)
}
