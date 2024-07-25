package utils

import (
	"encoding/json"
	"os"
)

type configuration struct {
	Token string `json:"token"`
	DbURL string `json:"db_url"`

	WelcomeChannelID         string `json:"welcome_channel_id"`
	LeaveChannelID           string `json:"leave_channel_id"`
	TimekeepingChannelID     string `json:"timekeeping_channel_id"`
	TimekeepingLogCategoryID string `json:"timekeeping_log_category_id"`

	AutoRoleId string `json:"auto_role_id"`
	UseBotID   string `json:"use_bot_id"`
}

var Config configuration

func init() {
	f, err := os.Open("config.json")
	if err != nil {
		panic("error opening config.json: " + err.Error())
	}
	err = json.NewDecoder(f).Decode(&Config)
	if err != nil {
		panic("error decoding config.json: " + err.Error())
	}
}
