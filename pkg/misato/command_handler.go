package misato

import (
	"github.com/bit-bunk3r/subaru/pkg/subaru/sublog"
	"github.com/bwmarrin/discordgo"
)

type callback = func(Context) error

type Context struct {
	Session *discordgo.Session
	Event   *discordgo.InteractionCreate
}

type command struct {
	info    discordgo.ApplicationCommand
	handler func(Context) error
}

var cmdLibrary map[string]command

func Register(callback callback, data discordgo.ApplicationCommand) {
	if cmdLibrary == nil {
		cmdLibrary = map[string]command{}
	}

	if data.Description == "" {
		panic("Command description must be not empty!")
	}

	cmdLibrary[data.Name] = command{
		info:    data,
		handler: callback,
	}
}

func EventHandler(sess *discordgo.Session, event *discordgo.InteractionCreate) {
	if cmd, ok := cmdLibrary[event.ApplicationCommandData().Name]; ok {
		err := cmd.handler(Context{
			Session: sess,
			Event:   event,
		})

		if err != nil {
			sublog.Logger.Warnf("Failed to execute command %s: %v", event.ApplicationCommandData().Name, err)
		}
	}
}

func RegisterAll(sess *discordgo.Session, guildID string) {

	deleteCmds, err := sess.ApplicationCommands(sess.State.Application.ID, guildID)

	if err == nil {
		for _, cmd := range deleteCmds {
			err = sess.ApplicationCommandDelete(sess.State.Application.ID, guildID, cmd.ID)
			if err != nil {
				sublog.Logger.Warnf("Unable to delete command %s: %v", cmd.Name, err)
			}
		}
	}

	if cmdLibrary == nil {
		sublog.Logger.Warnf("Cannot find any command")
	}

	for _, cmd := range cmdLibrary {
		if _, err := sess.ApplicationCommandCreate(sess.State.Application.ID, guildID, &cmd.info); err != nil {
			sublog.Logger.Errorf("Failed to register command: %v", err)
		}
	}
}
