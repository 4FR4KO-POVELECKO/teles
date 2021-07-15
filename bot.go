package teles

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type telegrambot struct {
	Bot    *tgbotapi.BotAPI
	ChatID int64
}

func new(token string) (*telegrambot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	bot.Debug = false

	id, result := checkChatID()
	if result {
		return &telegrambot{
			Bot:    bot,
			ChatID: id,
		}, nil
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		saveChatID(update.Message.Chat.ID)

		return &telegrambot{
			Bot:    bot,
			ChatID: update.Message.Chat.ID,
		}, nil
	}

	return nil, nil
}

func (b *telegrambot) send(text string) {
	msg := tgbotapi.NewMessage(b.ChatID, text)
	b.Bot.Send(msg)
}

func checkChatID() (int64, bool) {
	f, err := os.Open("./.teles/chat_id")
	if err != nil {
		return 0, false
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return 0, false
	}

	s := string(b)
	id, _ := strconv.ParseInt(s, 10, 64)

	return id, true
}

func saveChatID(id int64) {
	err := checkOrCreateDir("./.teles")
	if err != nil {
		return
	}

	f, err := checkOrCreateFile("./.teles/chat_id")
	if err != nil {
		return
	}

	defer f.Close()

	_, err = f.WriteString(fmt.Sprint(id))
	if err != nil {
		return
	}
}

func checkOrCreateDir(path string) error {
	result, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}

		return nil
	}
	if !result.IsDir() {
		return errors.New("Error to create directory, a file with the same name already exists")
	}

	return nil
}

func checkOrCreateFile(path string) (*os.File, error) {
	result, err := os.Stat(path)
	if os.IsNotExist(err) {
		f, err := os.Create(path)
		if err != nil {
			return nil, err
		}

		return f, nil
	}
	if result.IsDir() {
		return nil, errors.New("IsDir")
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return f, nil
}
