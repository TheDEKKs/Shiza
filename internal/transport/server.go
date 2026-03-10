package transport

import (
	"log/slog"
	"strings"
	"thedekk/Shiza/internal/api"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewService(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) {
	for up := range updates {
		if up.Message == nil {
			continue
		}

		if up.Message.Text != "" {
			fields := strings.Fields(up.Message.Text) // разбивает по пробелам
   			if len(fields) == 0 {
       	 	continue
    		}

    		firstWord := strings.ToLower(fields[0])

    		switch firstWord {
				case "шиз":
					answer, err := api.Request(up.Message.Text)
					if err != nil {
						slog.Error("Error making API request", "error", err)
						continue
					}

					msg := tgbotapi.NewMessage(up.Message.Chat.ID, *answer)
					if _, err := bot.Send(msg); err != nil {
						slog.Error("Error sending message", "error", err)
						continue
					}

					// Handle the "шиз" command

				default:
					continue
			}
		}
	}

	
}
