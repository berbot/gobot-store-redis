package redis_test

import (
	"testing"

	"github.com/botopolis/bot"
	"github.com/botopolis/bot/mock"
	"github.com/botopolis/redis"
	"github.com/stretchr/testify/assert"
)

type TestData struct {
	Name string
	Age  int
}

func TestRedis(t *testing.T) {
	assert := assert.New(t)

	store := redis.New("redis://localhost:6379")
	data := TestData{"Jean", 32}

	err := store.Set("jj", &data)
	assert.Nil(err)

	var out TestData

	err = store.Get("yy", &out)
	assert.NotNil(err)
	err = store.Get("jj", &out)
	assert.Nil(err)
	assert.Equal(data, out)

	err = store.Delete("jj")
	assert.Nil(err)
	err = store.Get("jj", &out)
	assert.NotNil(err)
}

func TestLoad(t *testing.T) {
	store := redis.New("redis://localhost:6379")
	robot := bot.New(mock.NewChat())
	store.Load(robot)

	err := store.Set("foo", "bar")
	assert.Nil(t, err)

	var out string
	err = robot.Brain.Get("foo", &out)
	assert.Nil(t, err)
	assert.Equal(t, "bar", out)
}
