package utils

import (
	"github.com/bwmarrin/discordgo"
)

func HasPermissionClear(m *discordgo.MessageCreate, roleUseBotId string) bool {
	// Check if roleUseBotId exists in m.Member.Roles
	for _, role := range m.Member.Roles {
		if role == roleUseBotId {
			return true
		}
	}

	return false
}
