package transport

import (
	"log/slog"
	"thedekk/Shiza/internal/api"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewService(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) {
	for up := range updates {
		if up.Message == nil {
			continue
		}

		if up.Message.Text != "" {
			answer, err := api.Request(up.Message.Text)
			if err != nil {
				slog.Error("Error making API request", "error", err)
				continue
			}
			if *answer == "NULL ANSWER" {
				continue
			}
			msg := tgbotapi.NewMessage(up.Message.Chat.ID, *answer)

			msg.ReplyToMessageID = up.Message.MessageID

			if _, err := bot.Send(msg); err != nil {
				slog.Error("Error sending message", "error", err)
				continue
			}
		}
	}

	
}
