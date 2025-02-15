package repositories

import (
	"avitoshop/internal/app/entities"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type redisInventoryRepository struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisInventoryRepository(client *redis.Client, ttl time.Duration) RedisInventoryRepository {
	return &redisInventoryRepository{client: client, ttl: ttl}
}

func (r *redisInventoryRepository) GetByUser(ctx context.Context, userID int) ([]entities.Inventory, error) {
	key := fmt.Sprintf("inventory:%d", userID)
	data, err := r.client.Get(ctx, key).Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, ErrCacheMiss
	} else if err != nil {
		return nil, fmt.Errorf("failed to get inventory from Redis: %w", err)
	}

	var inventory []entities.Inventory
	if err := json.Unmarshal(data, &inventory); err != nil {
		return nil, fmt.Errorf("failed to unmarshal inventory: %w", err)
	}

	return inventory, nil
}

func (r *redisInventoryRepository) SetByUser(ctx context.Context, userID int, inventory []entities.Inventory) error {
	key := fmt.Sprintf("inventory:%d", userID)
	data, err := json.Marshal(inventory)
	if err != nil {
		return fmt.Errorf("failed to marshal inventory: %w", err)
	}

	if err := r.client.Set(ctx, key, data, r.ttl).Err(); err != nil {
		return fmt.Errorf("failed to set inventory in Redis: %w", err)
	}

	return nil
}

func (r *redisInventoryRepository) DeleteByUser(ctx context.Context, userID int) error {
	key := fmt.Sprintf("inventory:%d", userID)
	if err := r.client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("failed to delete inventory from Redis: %w", err)
	}

	return nil
}
