package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bit-bunk3r/subaru/pkg/misato"
	_ "github.com/bit-bunk3r/subaru/pkg/subaru/commands"
	"github.com/bit-bunk3r/subaru/pkg/subaru/events"
	"github.com/bit-bunk3r/subaru/pkg/subaru/sublog"
	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	sublog.SetDeveplomentLogger()
	sublog.DiscordGoBindLog()

	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))

	if err != nil {
		sublog.Logger.Panicf("unable to create client: %v", err)
	}

	events.Setup(dg)
	dg.AddHandler(misato.EventHandler)
	dg.AddHandler(func(s *discordgo.Session, _ *discordgo.Ready) {
		if len(os.Args) > 1 && os.Args[1] == "--reload-commands" {
			sublog.Logger.Info("Reloading commands...")
			misato.RegisterAll(s, os.Getenv("GUILD_ID"))
			sublog.Logger.Info("Reloaded!")
		}
	})

	dg.Identify.Intents = discordgo.IntentsAll
	dg.LogLevel = discordgo.LogWarning
	dg.State.TrackMembers = true
	dg.Identify.Properties.OS = "android"

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
