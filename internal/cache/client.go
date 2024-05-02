package cache

import (
	"fmt"

	"github.com/redis/go-redis/v9"

	"cache-server/internal/config"
)

type Client struct {
	redis *redis.Client
}

func New(cfg config.Redis) *Client {
	return &Client{
		redis: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}
}
