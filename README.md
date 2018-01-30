# Gobot Store Redis

For use with bot

## Basic usage

```go
package main

import(
	"github.com/botopolis/bot"
	"github.com/botopolis/redis"
)

func main() {
	// define adapter
	robot := bot.New(adapter)

	robot.Install(
		redis.Plugin,
	)
}
```

## Customized configuration options

```go
package main

import(
	"github.com/botopolis/bot"
	"github.com/botopolis/redis"
)

func main() {
	// define adapter
	store := redis.New(os.Getenv("REDIS_URL))
	robot := bot.New(adapter)

	robot.Install(
    store.Plugin,
	)
}
```
