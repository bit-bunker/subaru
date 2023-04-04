package config

import "os"

func Load() Subaru {
	discordToken := os.Getenv("DISCORD_TOKEN")
	sub := Subaru{discordToken}
	return sub
}
