package commands

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
	util "github.com/gpnull/golang-github.com/utils"
)

func init() {
	util.Commands["clear"] = clear
}

func clear(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	// Check if the user has permission to delete messages
	if !util.HasPermissionClear(m, util.Config.UseBotID) {
		s.ChannelMessageSend(m.ChannelID, "You do not have permission to use this command.")
		return
	}

	// Parse the number of messages to delete
	numMessages := 0
	if len(args) > 0 {
		numMessages, _ = strconv.Atoi(args[0])
	}
	if numMessages == 0 {
		s.ChannelMessageSend(m.ChannelID, "The quantity to be deleted is invalid.")
		return
	}

	// Fetch messages to delete
	messages, err := s.ChannelMessages(m.ChannelID, numMessages, "", "", "")
	if err != nil {
		return
	}

	// Delete the fetched messages
	for _, msg := range messages {
		err = s.ChannelMessageDelete(m.ChannelID, msg.ID)
		if err != nil {
			continue
		}
		// time.Sleep(1 * time.Second) // Avoid rate limiting
	}

	// Notify the user about the deletion
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Successfully cleared %d messages.", len(messages)))
}
