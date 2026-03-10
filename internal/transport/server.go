package transport

import (
	"log/slog"
	"os"
	"strings"


	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewService(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI, db *pgxpool.Pool) error {
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
					// Handle the "шиз" command

				default:
					continue
			}
		}
	}

	return nil
}
