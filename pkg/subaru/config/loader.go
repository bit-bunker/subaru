package config

import "os"

func Load() Subaru {
	sub := Subaru{
		DiscordToken: os.Getenv("DISCORD_TOKEN"),
		GuildID:      os.Getenv("GUILD_ID"),
	}
	return sub
}
