# Gobot Store Redis

For use with gobot

## Basic usage

```go
package main

import(
	"github.com/berfarah/gobot"
	"github.com/berfarah/gobot-store-redis"
)

func main() {
	// define adapter
	robot := gobot.New(adapter)

	robot.Install(
		redis.Plugin,
	)
}
```

## Customized configuration options

```go
package main

import(
	"github.com/berfarah/gobot"
	"github.com/berfarah/gobot-store-redis"
)

func main() {
	// define adapter
	store := redis.New(os.Getenv("REDIS_URL))
	robot := gobot.New(adapter)

	robot.Install(
    store.Plugin,
	)
}
```
