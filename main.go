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

	bot_token := os.Getenv("BOT_TOKEN")
	bot.BotToken =  bot_token
	bot.Run()
}