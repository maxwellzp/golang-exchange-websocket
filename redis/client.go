package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client
var Ctx = context.Background()

func Init() {
	Client = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		DB:   0,
	})

	if _, err := Client.Ping(Ctx).Result(); err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}
}

func GetRate(pair string) (string, error) {
	return Client.Get(Ctx, pair).Result()
}
