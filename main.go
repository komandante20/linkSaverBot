package main

import (
	"flag"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	//t := mustToken()
	//tgClient := telegram.New(tgBotHost, mustToken())

}

func mustToken() string {
	token := flag.String(
		"tg-token",
		"",
		"telegram bot token",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("incorrect token for telegram bot")
	}

	return *token
}
