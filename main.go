package main

import (
	"butter-bot/bot"
)

func main() {
	// token, ok := os.LookupEnv("BOT_TOKEN")

	// if !ok {
	// 	log.Fatal("Must set bot token in env: BOT_TOKEN")
	// }

	bot.Token = "MTA1MDcwMTM4NDY3ODM3NTQ1NA.GBDQoT._dH6XMF9PcAkZy-YIuuNUqsm_odNZVi-3RXZI0"
	bot.Run()
}