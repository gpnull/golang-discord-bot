package commands

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	util "github.com/gpnull/golang-github.com/utils"
)

func init() {
	util.Commands["ping"] = ping
}

func ping(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	msg, err := s.ChannelMessageSend(m.ChannelID, "Pong...")
	if err != nil {
		return
	}
	msge, err := s.ChannelMessageEdit(m.ChannelID, msg.ID, "Pong?")
	if err != nil {
		return
	}

	start, _ := time.Parse(time.RFC3339, m.Timestamp.String())
	end, _ := time.Parse(time.RFC3339, msg.Timestamp.String())
	sendLatency := end.Sub(start)

	start = end
	end, _ = time.Parse(time.RFC3339, msge.EditedTimestamp.String())
	editLatency := end.Sub(start)

	_, err = s.ChannelMessageEdit(m.ChannelID, msg.ID, fmt.Sprintf("Pong!\n``` - Send latency: %dms (%dμs)\n - Edit latency: %dms (%dμs)\n - API latency: %dms (%dμs)```",
		sendLatency/1e6, sendLatency/1e3, editLatency/1e6, editLatency/1e3, s.HeartbeatLatency().Milliseconds(), s.HeartbeatLatency().Microseconds()))
	if err != nil {
		return
	}
}
