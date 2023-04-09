package welcomer

import (
	"fmt"
	"os"
	"time"

	"github.com/bit-bunk3r/subaru/pkg/subaru/sublog"
	"github.com/bwmarrin/discordgo"
)

func OnMemberJoin(sess *discordgo.Session, member *discordgo.GuildMemberAdd) {
	g, _ := sess.GuildWithCounts(member.GuildID)
	m, err := sess.ChannelMessageSendEmbed(os.Getenv("WELCOMER_CHANNEL_ID"), &discordgo.MessageEmbed{
		//Title:       "Bem-vindo",
		Description: fmt.Sprintf("%s, seja bem vindo ao BitBunker!\n\n\nVocê é nosso %d membro!", member.Mention(), g.ApproximateMemberCount),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: member.AvatarURL(""),
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("ID: %s", member.User.ID),
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0x90EE90,
	})

	if err != nil {
		sublog.Logger.Errorf("Unable to send (join) welcome message: %v", err)
	} else {
		sublog.Logger.Debugf("Sent welcome message: %v", m.ID)
	}
}

func OnMemberRemove(sess *discordgo.Session, member *discordgo.GuildMemberRemove) {
	g, err := sess.GuildWithCounts(member.GuildID)
	if err != nil {
		sublog.Logger.Warnf("unable to get guild: %v", err)
	}

	m, err := sess.ChannelMessageSendEmbed(os.Getenv("WELCOMER_CHANNEL_ID"), &discordgo.MessageEmbed{
		Description: fmt.Sprintf("%s, foi uma pena que você saiu do servidor, espero que você volte logo!\n\nVocê era nosso %d membro :(", member.Mention(), g.ApproximateMemberCount),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: member.AvatarURL(""),
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("ID: %s", member.User.ID),
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0xffcccb,
	})

	if err != nil {
		sublog.Logger.Errorf("Unable to send (leave) welcome message: %v", err)
	} else {
		sublog.Logger.Debugf("Sent welcome message: %v", m.ID)
	}
}
