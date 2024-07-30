package commands

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/gpnull/golang-github.com/utils"
)

func init() {
	utils.Commands["clear"] = clear
	utils.Commands["clearall"] = clearAll
}

func clear(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if !utils.HasPermissionClear(m, utils.Config.UseBotID) {
		s.ChannelMessageSend(m.ChannelID, "You do not have permission to use this command.")
		return
	}

	numMessages := 0
	if len(args) > 0 {
		numMessages, _ = strconv.Atoi(args[0])
	}
	if numMessages == 0 {
		s.ChannelMessageSend(m.ChannelID, "The quantity to be deleted is invalid.")
		return
	}

	messages, err := s.ChannelMessages(m.ChannelID, numMessages, "", "", "")
	if err != nil {
		return
	}

	for _, msg := range messages {
		err = s.ChannelMessageDelete(m.ChannelID, msg.ID)
		if err != nil {
			continue
		}
		// time.Sleep(1 * time.Second) // Avoid rate limiting
	}

	// s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Successfully cleared %d messages.", len(messages)))
}

func clearAll(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if !utils.HasPermissionClear(m, utils.Config.UseBotID) {
		s.ChannelMessageSend(m.ChannelID, "You do not have permission to use this command.")
		return
	}

	var messagesDeleted int
	for {
		messages, err := s.ChannelMessages(m.ChannelID, 100, "", "", "")
		if err != nil || len(messages) == 0 {
			break
		}

		messageIDs := make([]string, len(messages))
		for i, msg := range messages {
			messageIDs[i] = msg.ID
		}

		err = s.ChannelMessagesBulkDelete(m.ChannelID, messageIDs)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "An error occurred while deleting messages.")
			return
		}

		messagesDeleted += len(messages)
	}

	// s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Successfully deleted %d messages.", messagesDeleted))
}
