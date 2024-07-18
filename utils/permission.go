package utils

import (
	"github.com/bwmarrin/discordgo"
)

func HasPermissionClear(m *discordgo.MessageCreate, useBotId string) bool {
	// Check if useBotId exists in m.Member.Roles
	for _, role := range m.Member.Roles {
		if role == useBotId {
			return true
		}
	}

	return false
}
