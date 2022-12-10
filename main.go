package main

import (
	"butter-bot/bot"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("local.env")
	
	if err != nil {
		log.Fatal(err)
	}

	botToken := os.Getenv("BOT_TOKEN")
	guildID := os.Getenv("GUILD_ID")

	bot.Token = botToken
	bot.GuildID = guildID

	bot.Run()
}