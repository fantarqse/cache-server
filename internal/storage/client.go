package storage

import (
	"context"
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

func (c *Client) GetAll(ctx context.Context) error {
	panic("not implemented")
}

func (c *Client) Get(ctx context.Context) error {
	panic("not implemented")
}

func (c *Client) Put(ctx context.Context) error {
	panic("not implemented")
}

func (c *Client) Delete(ctx context.Context) error {
	panic("not implemented")
}
