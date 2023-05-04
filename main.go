package main

import (
	"flag"
	"log"
	tgClient "telegramBot/clients/telegram"
	event_consumer "telegramBot/consumer/event-consumer"
	"telegramBot/events/telegram"
	"telegramBot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service stopped", err)
	}
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
