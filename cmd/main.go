package main

import (
	"fmt"
	"log"
	"thedekk/Shiza/internal/env"
	"thedekk/Shiza/internal/transport"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	var config env.Config
	config.Load()
	fmt.Println("start bot")
	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = false
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)


	if err = transport.NewService(updates, bot); err != nil {

		log.Printf("", err)
	}
}
