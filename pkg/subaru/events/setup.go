package events

import (
	"github.com/bit-bunk3r/subaru/pkg/subaru/events/welcomer"
	"github.com/bwmarrin/discordgo"
)

func Setup(dg *discordgo.Session) {
	dg.AddHandler(welcomer.OnMemberJoin)
	dg.AddHandler(welcomer.OnMemberRemove)
}
