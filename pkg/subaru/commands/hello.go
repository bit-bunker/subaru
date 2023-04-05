package commands

import (
	"fmt"
	"runtime"

	"github.com/bit-bunk3r/subaru/pkg/misato"
	"github.com/bwmarrin/discordgo"
)

func Hello(ctx misato.Context) error {
	ctx.Session.InteractionRespond(ctx.Event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Olá! Estou usando o `%s`", runtime.Version()),
		},
	})

	return nil
}

func init() {
	misato.Register(Hello, discordgo.ApplicationCommand{
		Name:        "hello",
		Description: "Just testing, heh.",
	})
}
