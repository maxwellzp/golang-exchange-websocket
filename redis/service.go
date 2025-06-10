package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type Service struct {
	client *redis.Client
	ctx    context.Context
}

func NewService(add string) *Service {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: add,
		DB:   0,
	})

	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}

	return &Service{
		client: client,
		ctx:    ctx,
	}
}

func (s *Service) GetRate(pair string) (string, error) {
	return s.client.Get(s.ctx, pair).Result()
}
func (s *Service) SetRate(pair, value string) error {
	return s.client.Set(s.ctx, pair, value, 0).Err()
}
