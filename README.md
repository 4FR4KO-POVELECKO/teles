# Teles ![](.github/rsz_1unnamed.jpg)
Teles is a small logger for Go.

Logging in telegram bot, file ```.log``` , terminal. 


## Installation

```bash
go get -u github.com/4FR4KO-POVELECKO/teles
```

## Examples

First, create bot in [BotFather](https://telegram.me/BotFather).

Create log directory:
```bash
mkdir log
```

Start:
```go
package main

import (
	"log"
	"github.com/4FR4KO-POVELECKO/teles"
)

func main() {
	err := teles.Start("TOKEN")
	if err != nil {
		log.Fatal(err)
	}
}
```

Usage:
```go
teles.Logger("log message")
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)