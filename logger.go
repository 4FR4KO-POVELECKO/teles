package teles

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type Logger struct {
	BotToken  string
	BotLevels []Level

	DirPath   string
	DirLevels []Level
}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) NewBot(token string, levels []Level) {

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

func (l *Logger) Trace(args ...interface{}) {
	result := l.checkToArray(Trace)
	if result {
		l.LogToFile("./log/trace.log", Trace, args...)
	}

	l.Log(Trace, args...)
}

func (l *Logger) Debug(args ...interface{}) {
	result := l.checkToArray(Debug)
	if result {
		l.LogToFile("./log/debug.log", Debug, args...)
	}

	l.Log(Debug, args...)
}

func (l *Logger) Info(args ...interface{}) {
	result := l.checkToArray(Info)
	if result {
		l.LogToFile("./log/info.log", Info, args...)
	}
	l.Log(Info, args...)
}

func (l *Logger) Warning(args ...interface{}) {
	result := l.checkToArray(Warning)
	if result {
		l.LogToFile("./log/warn.log", Warning, args...)
	}
	l.Log(Warning, args...)
}

func (l *Logger) Error(args ...interface{}) {
	result := l.checkToArray(Error)
	if result {
		l.LogToFile("./log/error.log", Error, args...)
	}

	l.Log(Error, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	result := l.checkToArray(Fatal)
	if result {
		l.LogToFile("./log/fatal.log", Fatal, args...)
	}
	l.Log(Fatal, args...)
	os.Exit(1)
}

func (l *Logger) Panic(args ...interface{}) {
	result := l.checkToArray(Panic)
	if result {
		l.LogToFile("./log/panic.log", Panic, args...)
	}
	l.Log(Panic, args...)
	panic(args)
}

// Helpers

func (l *Logger) checkToArray(level Level) bool {
	for _, value := range l.DirLevels {
		if value == level {
			return true
		}
	}

	return false
}
