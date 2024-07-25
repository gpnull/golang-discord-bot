package utils

import (
	"github.com/bwmarrin/discordgo"
)

func HasPermissionClear(m *discordgo.MessageCreate, roleUseBotId string) bool {
	for _, role := range m.Member.Roles {
		if role == roleUseBotId {
			return true
		}
	}

	return false
}
