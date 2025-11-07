package tg_bot_adviser_read

import (
	"flag"
	"log"
	"tg-bot-adviser-read/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(mustToken())
}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
