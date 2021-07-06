package teles

import (
	"fmt"
	"os"
	"time"
)

var now = time.Now().Format("2006.01.02 15:04:05")

type Logger struct {
	BotToken  string
	BotLevels []Level

	DirPath   string
	DirLevels []Level
}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) NewBot(token string, levels []Level) {}
func (l *Logger) NewDir(path string, levels []Level)  {}

func (l *Logger) Log(level Level, args ...interface{}) {
	str := fmt.Sprintf("%v", args...)
	fmt.Printf("[%s] %s %s\n", now, level, str)
}

func (l *Logger) Trace(args ...interface{}) {
	l.Log(Trace, args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.Log(Debug, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.Log(Info, args...)
}

func (l *Logger) Warning(args ...interface{}) {
	l.Log(Warning, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.Log(Error, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.Log(Fatal, args...)
	os.Exit(1)
}

func (l *Logger) Panic(args ...interface{}) {
	l.Log(Panic, args...)
	panic(args)
}
