package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	util "github.com/gpnull/golang-github.com/utils"
)

func init() {
	util.Commands["test"] = test
}

func test(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	for i, arg := range args {
		fmt.Printf("arg %d: %s\n", i, arg)
	}
}
