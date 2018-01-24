package redis_test

import (
	"os"

	"github.com/berfarah/gobot"
	"github.com/berfarah/gobot-slack"
	"github.com/berfarah/gobot-store-redis"
)

func Example() {
	robot := gobot.New(slack.New("secret"))
	robot.Install(
		redis.New(os.Getenv("REDIS_URL")),
	)
	robot.Run()
}
