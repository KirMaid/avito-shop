package repositories

import (
	"avitoshop/internal/app/entities"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
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

	result, err := r.client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get inventory from Redis: %w", err)
	}

	if len(result) == 0 {
		return nil, ErrCacheMiss
	}

	var inventory []entities.Inventory
	for goodIDStr, quantityStr := range result {
		goodID, _ := strconv.Atoi(goodIDStr)
		quantity, _ := strconv.Atoi(quantityStr)

		inventory = append(inventory, entities.Inventory{
			UserID:   userID,
			GoodID:   goodID,
			Quantity: quantity,
		})
	}

	return inventory, nil
}

func (r *redisInventoryRepository) SetByUser(ctx context.Context, userID int, inventory []entities.Inventory) error {
	key := fmt.Sprintf("inventory:%d", userID)

	fields := make(map[string]interface{})
	for _, item := range inventory {
		fields[strconv.Itoa(item.GoodID)] = item.Quantity
	}

	err := r.client.HSet(ctx, key, fields).Err()
	if err != nil {
		return fmt.Errorf("failed to set inventory in Redis: %w", err)
	}

	if r.ttl > 0 {
		err = r.client.Expire(ctx, key, r.ttl).Err()
		if err != nil {
			return fmt.Errorf("failed to set TTL for inventory key: %w", err)
		}
	}

	return nil
}

func (r *redisInventoryRepository) InsertOrUpdate(ctx context.Context, inventory *entities.Inventory) error {
	key := fmt.Sprintf("inventory:%d", inventory.UserID)

	err := r.client.HSet(ctx, key, strconv.Itoa(inventory.GoodID), inventory.Quantity).Err()
	if err != nil {
		return fmt.Errorf("failed to set inventory item in Redis: %w", err)
	}

	if r.ttl > 0 {
		err = r.client.Expire(ctx, key, r.ttl).Err()
		if err != nil {
			return fmt.Errorf("failed to set TTL for inventory key: %w", err)
		}
	}

	return nil
}

func (r *redisInventoryRepository) DeleteByUser(ctx context.Context, userID int) error {
	key := fmt.Sprintf("inventory:%d", userID)

	err := r.client.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete inventory from Redis: %w", err)
	}

	return nil
}
