package db

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"random.chars.jp/git/misskey/config"
)

// NewRedisClient returns a new Redis client with the parameters
// specified in the config file.
func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			config.External.Redis.Host,
			config.External.Redis.Port),
		DB:       config.External.Redis.DB,
		Password: config.External.Redis.Password,
	})
}
