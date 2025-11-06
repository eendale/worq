package redisclient

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func New(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}

func Ping(client *redis.Client) error {
	ctx := context.Background()
	return client.Ping(ctx).Err()
}
