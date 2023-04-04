package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bit-bunk3r/subaru/pkg/subaru/config"
	"github.com/bit-bunk3r/subaru/pkg/subaru/sublog"
	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config := config.Load()

	sublog.SetDeveplomentLogger()
	sublog.DiscordGoBindLog()

	dg, err := discordgo.New("Bot " + config.DiscordToken)

	if err != nil {
		sublog.Logger.Panicf("unable to create client: %v", err)
	}

	dg.Identify.Intents = discordgo.IntentsAll
	dg.LogLevel = discordgo.LogInformational

	err = dg.Open()
	if err != nil {
		sublog.Logger.Panicf("unable to open connection with discord: %v", err)
	}

	sublog.Logger.Info("Connected to discord! Press Control + C to Close (or send SIGTERM).")

	defer dg.Close()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
