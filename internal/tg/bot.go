package tg

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramBot struct {
	Bot    *tgbotapi.BotAPI
	ChatID int64
}

func New(token string) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		return &TelegramBot{
			Bot:    bot,
			ChatID: update.Message.Chat.ID,
		}, nil
	}

	return nil, nil
}

func (b *TelegramBot) Send(text string) {
	msg := tgbotapi.NewMessage(b.ChatID, text)
	b.Bot.Send(msg)
}
