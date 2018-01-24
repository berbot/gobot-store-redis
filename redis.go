package redis

import (
	"github.com/berfarah/gobot"
	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack"
)

// Store conforms to the gobot.Store interface
// It uses msgpack to cache records in Redis
type Store struct {
	redis *redis.Client
	cache *cache.Codec
}

// New allows for configuration of the redis plugin with a
// custom address
func New(address string) *Store {
	redis := redis.NewClient(&redis.Options{Addr: address})
	return &Store{
		redis: redis,
		cache: &cache.Codec{
			Redis: redis,
			Marshal: func(v interface{}) ([]byte, error) {
				return msgpack.Marshal(v)
			},
			Unmarshal: func(b []byte, v interface{}) error {
				return msgpack.Unmarshal(b, v)
			},
		},
	}
}

// Set writes the given interface to the key as msgpack
func (s *Store) Set(key string, i interface{}) error {
	return s.cache.Set(&cache.Item{Key: key, Object: i})
}

// Get reads at the key into the interface from msgpack
func (s *Store) Get(key string, i interface{}) error {
	return s.cache.Get(key, i)
}

// Delete removes the record at the key
func (s *Store) Delete(key string) error {
	return s.cache.Delete(key)
}

// Load installs this store as a gobot.Plugin
func (s *Store) Load(r *gobot.Robot) { r.Brain.SetStore(s) }

// Unload disconnects from Redis
func (s *Store) Unload(r *gobot.Robot) { s.redis.Close() }
