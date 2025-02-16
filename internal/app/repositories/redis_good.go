package repositories

import (
	"avitoshop/internal/app/entities"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type redisGoodRepository struct {
	client *redis.Client
	ttl    time.Duration
}

func (r *redisGoodRepository) GetByIDs(ctx context.Context, goodsIDs []int) ([]entities.Good, error) {
	//TODO implement me
	panic("implement me")
}

func NewRedisGoodRepository(client *redis.Client, ttl time.Duration) RedisGoodRepository {
	return &redisGoodRepository{client: client, ttl: ttl}
}

func (r *redisGoodRepository) GetByID(ctx context.Context, id int) (*entities.Good, error) {
	key := fmt.Sprintf("good:id:%d", id)

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

	var good entities.Good
	if err := res.Scan(&good); err != nil {
		return nil, err
	}

	return &good, nil
}

func (r *redisGoodRepository) SetByID(ctx context.Context, id int, good *entities.Good) error {
	key := fmt.Sprintf("good:id:%d", id)

	err := r.client.HSet(ctx, key, map[string]interface{}{
		"id":    good.ID,
		"name":  good.Name,
		"price": good.Price,
	}).Err()

	if err != nil {
		return fmt.Errorf("failed to set good in Redis: %w", err)
	}

	return nil
}

func (r *redisGoodRepository) GetByName(ctx context.Context, name string) (*entities.Good, error) {
	key := fmt.Sprintf("good:name:%s", name)

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

	var good entities.Good
	if err := res.Scan(&good); err != nil {
		return nil, err
	}

	return &good, nil
}

func (r *redisGoodRepository) SetByName(ctx context.Context, name string, good *entities.Good) error {
	key := fmt.Sprintf("good:name:%s", name)

	err := r.client.HSet(ctx, key, map[string]interface{}{
		"id":    good.ID,
		"name":  good.Name,
		"price": good.Price,
	}).Err()

	if err != nil {
		return fmt.Errorf("failed to set good in Redis: %w", err)
	}

	return nil
}
