package redis_test

import (
	"os"

	"github.com/botopolis/bot"
	"github.com/botopolis/bot/mock"
	"github.com/botopolis/redis"
)

func Example() {
	bot.New(
		mock.NewChat(),
		redis.New(os.Getenv("REDIS_URL")),
	).Run()
}
