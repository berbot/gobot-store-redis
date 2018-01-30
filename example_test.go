package redis_test

import (
	"os"

	"github.com/botopolis/bot"
	"github.com/botopolis/redis"
	"github.com/botopolis/slack"
)

func Example() {
	bot.New(
		slack.New("secret"),
		redis.New(os.Getenv("REDIS_URL")),
	).Run()
}
