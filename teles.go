package teles

import (
	"log"
	"os"
)

var b *bot

// Start create a new bot
func Start(token string) error {
	bot, err := newBot(token)
	if err != nil {
		return err
	}

	b = bot

	return nil
}

// Logger ...
// Get message string
//
// Creates a .log file and writes the message, need to create a directory 'log'
// Send message to the Telegram bot
// Log message
func Logger(msg string) error {
	f, err := os.OpenFile("log/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	logger := log.New(f, "Logger", log.LstdFlags)

	if err := b.sendMessage(msg); err != nil {
		return err
	}
	logger.Println(msg)

	return nil
}
