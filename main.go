package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gpnull/golang-github.com/gpnull/golang-discord-bot/config"
	"github.com/gpnull/golang-github.com/gpnull/golang-discord-bot/database"
	"github.com/gpnull/golang-github.com/gpnull/golang-discord-bot/handlers"

	"github.com/bwmarrin/discordgo"
)

const (
	defaultConfigPath = "config.yaml"
)

var (
	CONFIG_FILE_PATH string
)

func main() {
	// Flag
	flag.StringVar(&CONFIG_FILE_PATH, "config", defaultConfigPath, "Config path")
	flag.StringVar(&CONFIG_FILE_PATH, "c", defaultConfigPath, "Config path (shorthand)")

	cfg, err := config.LoadConfig(CONFIG_FILE_PATH)
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	// Connect to MongoDB
	dbClient, err := database.ConnectDB(cfg.MongoDBURI)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	defer dbClient.DisconnectDB(context.Background())

	// Create Discord bot
	dg, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		fmt.Println("Error creating Discord bot:", err)
		return
	}

	// Add handlers
	dg.AddHandler(handlers.Welcome(dbClient))
	dg.AddHandler(handlers.Goodbye())
	dg.AddHandler(handlers.Moderation(dg))
	dg.AddHandler(handlers.Message(cfg.IdChannelMessage))

	// Connect to Discord
	err = dg.Open()
	if err != nil {
		fmt.Println("Error connecting to Discord:", err)
		return
	}
	defer dg.Close()

	fmt.Println("Bot is ready!")

	// Wait for stop signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	fmt.Println("Shutting down bot...")
}
