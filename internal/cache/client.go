package cache

import "github.com/redis/go-redis/v9"

type Client struct {
	redis *redis.Client
}

func New(addr string) *Client {
	return &Client{
		redis: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}
}
