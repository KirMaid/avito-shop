package repositories

import (
	"avitoshop/internal/app/entities"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type redisMerchRepository struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisMerchRepository(client *redis.Client, ttl time.Duration) RedisMerchRepository {
	return &redisMerchRepository{client: client, ttl: ttl}
}

func (r *redisMerchRepository) GetByName(ctx context.Context, name string) (*entities.Merch, error) {
	key := fmt.Sprintf("merch:name:%s", name)

	exists, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to check if key exists in Redis: %w", err)
	}

	if exists == 0 {
		return nil, ErrCacheMiss
	}

	res := r.client.HGetAll(ctx, key)
	if res.Err() != nil {
		return nil, res.Err()
	}

	var merch entities.Merch
	if err := res.Scan(&merch); err != nil {
		return nil, err
	}

	return &merch, nil
}

func (r *redisMerchRepository) SetByName(ctx context.Context, name string, merch *entities.Merch) error {
	key := fmt.Sprintf("merch:name:%s", name)

	err := r.client.HSet(ctx, key, map[string]interface{}{
		"id":    merch.ID,
		"name":  merch.Name,
		"price": merch.Price,
	}).Err()

	if err != nil {
		return fmt.Errorf("failed to set merch in Redis: %w", err)
	}

	return nil
}
