package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
	rnm "github.com/pitakill/rickandmortyapigowrapper"
)

func charactersHandler (s *discordgo.Session, i *discordgo.InteractionCreate) {

	// Search the API and find matching results

	characterName := i.ApplicationCommandData().Options
	_ = characterName

	character, err := rnm.GetCharacter(1)

	if err != nil {
		log.Panicf("Could not get %v", character)
	}

	session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Retrieving",
		},
	})
}

func helpHandler (s *discordgo.Session, i *discordgo.InteractionCreate) {

	session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource, 
		Data: &discordgo.InteractionResponseData{
			Content: "Hello there!",
		},
	})
}

func episodesHandler(s*discordgo.Session, i *discordgo.InteractionCreate) {

	session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Retrieving episodes",
		},
	})
}