package teles

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/4FR4KO-POVELECKO/teles/internal/tg"
)

type Logger struct {
	Bot       *tg.TelegramBot
	BotLevels []Level

	DirPath   string
	DirLevels []Level
}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) NewBot(token string, levels []Level) error {
	bot, err := tg.New(token)
	if err != nil {
		return errors.New("Error to connect the bot")
	}

	l.Bot = bot
	l.BotLevels = levels

	return nil
}

func (l *Logger) NewDir(path string, levels []Level) error {
	result, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return errors.New("Error to create directory")
		}

		l.DirPath = path
		l.DirLevels = levels

		return nil
	}
	if result.IsDir() {
		l.DirPath = path
		l.DirLevels = levels

		return nil
	}

	return errors.New("Error to create directory, a file with the same name already exists")
}

// Loggers

var now = time.Now().Format("2006.01.02 15:04:05")

func (l *Logger) Log(level Level, args ...interface{}) {
	str := fmt.Sprintf("%v", args...)
	fmt.Printf("[%s] %s %s\n", now, level, str)
}

func (l *Logger) LogToFile(path string, level Level, args ...interface{}) {
	str := fmt.Sprintf("%v", args...)
	write := fmt.Sprintf("[%s] %s %s\n", now, level, str)

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		f, err := os.Create(path)
		if err != nil {
			l.Error(errors.New("Error to create file"))
			return
		}

		defer f.Close()
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		l.Error(errors.New("Error to open file"))
		return
	}
	defer f.Close()

	_, err = f.WriteString(write)
	if err != nil {
		l.Error(errors.New("Error to write file"))
		return
	}
}

func (l *Logger) GetLogStr(level Level, args ...interface{}) string {
	str := fmt.Sprintf("%v", args...)
	write := fmt.Sprintf("[%s] %s %s\n", now, level, str)

	return write
}

func (l *Logger) Trace(args ...interface{}) {
	result := l.checkToArray(Trace, l.DirLevels)
	if result {
		l.LogToFile("./log/trace.log", Trace, args...)
	}

	result = l.checkToArray(Trace, l.BotLevels)
	if result {
		msg := l.GetLogStr(Trace, args...)
		l.Bot.Send(msg)
	}

	l.Log(Trace, args...)
}

func (l *Logger) Debug(args ...interface{}) {
	result := l.checkToArray(Debug, l.DirLevels)
	if result {
		l.LogToFile("./log/debug.log", Debug, args...)
	}

	result = l.checkToArray(Debug, l.BotLevels)
	if result {
		msg := l.GetLogStr(Debug, args...)
		l.Bot.Send(msg)
	}

	l.Log(Debug, args...)
}

func (l *Logger) Info(args ...interface{}) {
	result := l.checkToArray(Info, l.DirLevels)
	if result {
		l.LogToFile("./log/info.log", Info, args...)
	}

	result = l.checkToArray(Info, l.BotLevels)
	if result {
		msg := l.GetLogStr(Info, args...)
		l.Bot.Send(msg)
	}

	l.Log(Info, args...)
}

func (l *Logger) Warning(args ...interface{}) {
	result := l.checkToArray(Warning, l.DirLevels)
	if result {
		l.LogToFile("./log/warn.log", Warning, args...)
	}

	result = l.checkToArray(Warning, l.BotLevels)
	if result {
		msg := l.GetLogStr(Warning, args...)
		l.Bot.Send(msg)
	}

	l.Log(Warning, args...)
}

func (l *Logger) Error(args ...interface{}) {
	result := l.checkToArray(Error, l.DirLevels)
	if result {
		l.LogToFile("./log/error.log", Error, args...)
	}

	result = l.checkToArray(Error, l.BotLevels)
	if result {
		msg := l.GetLogStr(Error, args...)
		l.Bot.Send(msg)
	}

	l.Log(Error, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	result := l.checkToArray(Fatal, l.DirLevels)
	if result {
		l.LogToFile("./log/fatal.log", Fatal, args...)
	}

	result = l.checkToArray(Fatal, l.BotLevels)
	if result {
		msg := l.GetLogStr(Fatal, args...)
		l.Bot.Send(msg)
	}

	l.Log(Fatal, args...)
	os.Exit(1)
}

func (l *Logger) Panic(args ...interface{}) {
	result := l.checkToArray(Panic, l.DirLevels)
	if result {
		l.LogToFile("./log/panic.log", Panic, args...)
	}

	result = l.checkToArray(Panic, l.BotLevels)
	if result {
		msg := l.GetLogStr(Panic, args...)
		l.Bot.Send(msg)
	}

	l.Log(Panic, args...)
	panic(args)
}

// Helpers

func (l *Logger) checkToArray(level Level, array []Level) bool {
	for _, value := range array {
		if value == level {
			return true
		}
	}

	return false
}
