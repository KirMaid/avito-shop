package repositories

import (
	"avitoshop/internal/app/entities"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type redisUserRepository struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisUserRepository(client *redis.Client, ttl time.Duration) RedisUserRepository {
	return &redisUserRepository{client: client, ttl: ttl}
}

func (r *redisUserRepository) GetByUsername(ctx context.Context, username string) (*entities.User, error) {
	key := fmt.Sprintf("user:username:%s", username)

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

	var user entities.User
	if err := res.Scan(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *redisUserRepository) GetUsernamesByIDs(ctx context.Context, userIDs []int) (map[int]string, error) {
	keys := make([]string, 0, len(userIDs))
	for _, id := range userIDs {
		keys = append(keys, fmt.Sprintf("user:id:%d", id))
	}

	results, err := r.client.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get usernames from Redis: %w", err)
	}

	usernames := make(map[int]string)
	for i, result := range results {
		if result == nil {
			continue
		}

		var user entities.User
		if err := json.Unmarshal([]byte(result.(string)), &user); err != nil {
			return nil, fmt.Errorf("failed to unmarshal user: %w", err)
		}

		usernames[userIDs[i]] = user.Username
	}

	return usernames, nil
}

func (r *redisUserRepository) SetByUsername(ctx context.Context, username string, user *entities.User) error {
	key := fmt.Sprintf("user:username:%s", username)

	err := r.client.HSet(ctx, key, map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"password": user.Password,
		"balance":  user.Balance,
	}).Err()

	if err != nil {
		return fmt.Errorf("failed to set user in Redis: %w", err)
	}

	return nil
}

func (r *redisUserRepository) DeleteByUsername(ctx context.Context, username string) error {
	key := fmt.Sprintf("user:%s", username)
	if err := r.client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("failed to delete user from Redis: %w", err)
	}

	return nil
}
