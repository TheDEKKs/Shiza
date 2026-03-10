package main

import (
	"fmt"
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
		fmt.Println(err)
		return
	}
	bot.Debug = false
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)


	transport.NewService(updates, bot)
}
