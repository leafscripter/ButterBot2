package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	BotToken string
)

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {

	// Ignore bot message
	if message.Author.ID == discord.State.User.ID {
		return 
	}

	switch {
	case strings.Contains(message.Content, "character"):
		discord.ChannelMessageSend(message.ChannelID, "Let's find a character!")
	case strings.Contains(message.Content, "episode"):
		discord.ChannelMessageSend(message.ChannelID, "Let's find an episode!")
	}
}

func Run() {
	// discord, err := discordgo.New("Bot" + "MTA1MDcwMTM4NDY3ODM3NTQ1NA.GBDQoT._dH6XMF9PcAkZy-YIuuNUqsm_odNZVi-3RXZI0")
	discord, err := discordgo.New("Bot " + BotToken)

	if err != nil {
		log.Fatal(err)
	}

	discord.AddHandler(newMessage)

	// Open session
	discord.Open()
	defer discord.Close()

	fmt.Println("Bot running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c	// Wait for interruption / kill signal
}