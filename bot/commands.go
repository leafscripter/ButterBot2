package bot

import (
	"github.com/bwmarrin/discordgo"
)

var (
	help = &discordgo.ApplicationCommand{
		Name: "help",
		Description: "Shows how to use bot",
	}

	characters = &discordgo.ApplicationCommand{
		Name: "characters",
		Description: "Retrieves character from Rick and Morty API",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type: discordgo.ApplicationCommandOptionString,
				Name: "character-name",
				Description: "Name of character",
				Required: true,
			},
		},
	}

	episodes = &discordgo.ApplicationCommand{
		Name: "episodes",
		Description: "Retrieves episodes from Rick and Morty API",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type: discordgo.ApplicationCommandOptionString,
				Name: "episode-title",
				Description: "Title of the episode",
				Required: true,
			},
		},
	}

)
