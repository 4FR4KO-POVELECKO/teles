package teles

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type chat struct {
	chatID int64
}

type bot struct {
	botapi *tgbotapi.BotAPI
}

var c = &chat{}

// newBot ...
// Get token, returns the bot type, or an error.
func newBot(token string) (*bot, error) {
	b, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	b.Debug = false
	log.Printf("Authorized on account %s", b.Self.UserName)

	return &bot{botapi: b}, nil
}

// sendMessage ...
// Get message, returns error.
//
// Parses json file, returns chatID
// If there is no file or it is empty, the bot waits for the '/start' command
// Creates a json file and writes the chatID
//
// Send message
func (b *bot) sendMessage(message string) error {
	if b == nil {
		return errors.New("Not valid bot")
	}

	jsonFile, _ := os.Open("chat_ID.json")
	defer jsonFile.Close()

	byteVal, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteVal, &c.chatID)

	if c.chatID == 0 {
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60

		updates, err := b.botapi.GetUpdatesChan(u)
		if err != nil {
			return err
		}

		for update := range updates {
			switch update.Message.Command() {
			case "start":
				c.chatID = update.Message.Chat.ID

				file, err := json.MarshalIndent(c.chatID, "", " ")
				if err != nil {
					return err
				}

				err = ioutil.WriteFile("chat_ID.json", file, 0644)
				if err != nil {
					return err
				}
			}
			break
		}
	}

	msg := tgbotapi.NewMessage(c.chatID, message)
	if _, err := b.botapi.Send(msg); err != nil {
		return err
	}

	return nil
}
