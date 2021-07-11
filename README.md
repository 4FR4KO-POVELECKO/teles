# Teles ![](.github/logo.jpg)
Teles is a small logger for Go.

Logging in:
- Telegram bot 
- File ```.log```
- Terminal 

## Installation

```bash
go get -u github.com/4FR4KO-POVELECKO/teles
```

## Examples

First, create bot in [BotFather](https://telegram.me/BotFather).

Start:
```go
package main

import (
	"github.com/4FR4KO-POVELECKO/teles"
)

func main() {
	// Create new logger
	logger := teles.New()

	// Choose levels
	levels := []teles.Level{
		teles.Panic,
		teles.Fatal,
		teles.Error,
		teles.Warning,
		teles.Info,
		teles.Debug,
		teles.Trace,
	}

	// Create logger to directory
	err := logger.NewDir("./log", levels)
	if err != nil {
		logger.Error(err)
	}

	// Create logger to tg bot
	err = logger.NewBot("BOT_TOKEN", levels)
	if err != nil {
		logger.Error(err)
	}
}

```

Usage:
```go
logger.Log(teles.Info, "text")
logger.Trace("text")
logger.Debug("text")
logger.Info("text")
logger.Error("text")
logger.Warning("text")
logger.Fatal("text")
logger.Panic("text")
```

Levels:
```go
package teles

type Level string

const (
	Panic   Level = "PANIC:"
	Fatal   Level = "FATAL:"
	Error   Level = "ERROR:"
	Warning Level = "WARNING:"
	Info    Level = "INFO:"
	Debug   Level = "DEBUG:"
	Trace   Level = "TRACE:"
)
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)