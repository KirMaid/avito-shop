package repositories

import (
	"avitoshop/internal/app/entities"
	"context"
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

func (r *redisUserRepository) GetById(ctx context.Context, id int) (*entities.User, error) {
	key := fmt.Sprintf("user:id:%d", id)

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

func (r *redisUserRepository) GetUsernamesByIDs(ctx context.Context, userIDs []int) (map[int]string, []int, error) {
	usernames := make(map[int]string)
	missingIDs := make([]int, 0)
	for _, id := range userIDs {
		key := fmt.Sprintf("user:id:%d", id)

		exists, err := r.client.Exists(ctx, key).Result()
		if err != nil || exists == 0 {
			missingIDs = append(missingIDs, id)
			continue
		}

		res := r.client.HGet(ctx, key, "username")
		if res.Err() != nil {
			missingIDs = append(missingIDs, id)
			continue
		}

		var username string
		if err := res.Scan(&username); err != nil {
			missingIDs = append(missingIDs, id)
			continue
		}

		usernames[id] = username
	}

	return usernames, missingIDs, nil
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

	if r.ttl > 0 {
		err = r.client.Expire(ctx, key, r.ttl).Err()
		if err != nil {
			return fmt.Errorf("failed to set TTL for user key: %w", err)
		}
	}

	return nil
}

func (r *redisUserRepository) SetById(ctx context.Context, id int, user *entities.User) error {
	key := fmt.Sprintf("user:id:%d", id)

	err := r.client.HSet(ctx, key, map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"password": user.Password,
		"balance":  user.Balance,
	}).Err()

	if err != nil {
		return fmt.Errorf("failed to set user in Redis: %w", err)
	}

	if r.ttl > 0 {
		err = r.client.Expire(ctx, key, r.ttl).Err()
		if err != nil {
			return fmt.Errorf("failed to set TTL for user key: %w", err)
		}
	}

	return nil
}
