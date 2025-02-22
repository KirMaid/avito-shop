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

type redisTransactionRepository struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisTransactionRepository(client *redis.Client, ttl time.Duration) RedisTransactionRepository {
	return &redisTransactionRepository{client: client, ttl: ttl}
}

func (r *redisTransactionRepository) GetReceivedTransactions(ctx context.Context, userID int) ([]entities.Transaction, error) {
	key := fmt.Sprintf("transactions:received:%d", userID)
	data, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, ErrCacheMiss
		}
		return nil, fmt.Errorf("failed to get received transactions from Redis: %w", err)
	}

	var transactions []entities.Transaction
	if err := json.Unmarshal(data, &transactions); err != nil {
		return nil, fmt.Errorf("failed to unmarshal received transactions: %w", err)
	}

	return transactions, nil
}

func (r *redisTransactionRepository) SetReceivedTransactions(ctx context.Context, userID int, transactions []entities.Transaction) error {
	key := fmt.Sprintf("transactions:received:%d", userID)
	data, err := json.Marshal(transactions)
	if err != nil {
		return fmt.Errorf("failed to marshal received transactions: %w", err)
	}

	if err := r.client.Set(ctx, key, data, r.ttl).Err(); err != nil {
		return fmt.Errorf("failed to set received transactions in Redis: %w", err)
	}

	if r.ttl > 0 {
		err = r.client.Expire(ctx, key, r.ttl).Err()
		if err != nil {
			return fmt.Errorf("failed to set TTL for transactions: %w", err)
		}
	}

	return nil
}

func (r *redisTransactionRepository) GetSentTransactions(ctx context.Context, userID int) ([]entities.Transaction, error) {
	key := fmt.Sprintf("transactions:sent:%d", userID)
	data, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, ErrCacheMiss
		}
		return nil, fmt.Errorf("failed to get sent transactions from Redis: %w", err)
	}

	var transactions []entities.Transaction
	if err := json.Unmarshal(data, &transactions); err != nil {
		return nil, fmt.Errorf("failed to unmarshal sent transactions: %w", err)
	}

	return transactions, nil
}

func (r *redisTransactionRepository) SetSentTransactions(ctx context.Context, userID int, transactions []entities.Transaction) error {
	key := fmt.Sprintf("transactions:sent:%d", userID)
	data, err := json.Marshal(transactions)
	if err != nil {
		return fmt.Errorf("failed to marshal sent transactions: %w", err)
	}

	if err := r.client.Set(ctx, key, data, r.ttl).Err(); err != nil {
		return fmt.Errorf("failed to set sent transactions in Redis: %w", err)
	}

	if r.ttl > 0 {
		err = r.client.Expire(ctx, key, r.ttl).Err()
		if err != nil {
			return fmt.Errorf("failed to set TTL for transactions: %w", err)
		}
	}

	return nil
}

func (r *redisTransactionRepository) AddReceivedTransaction(ctx context.Context, userID int, transaction *entities.Transaction) error {
	transactions, err := r.GetReceivedTransactions(ctx, userID)
	if err != nil && !errors.Is(err, ErrCacheMiss) {
		return fmt.Errorf("failed to get received transactions: %w", err)
	}

	transactions = append(transactions, *transaction)
	return r.SetReceivedTransactions(ctx, userID, transactions)
}

func (r *redisTransactionRepository) AddSentTransaction(ctx context.Context, userID int, transaction *entities.Transaction) error {
	transactions, err := r.GetSentTransactions(ctx, userID)
	if err != nil && !errors.Is(err, ErrCacheMiss) {
		return fmt.Errorf("failed to get sent transactions: %w", err)
	}

	transactions = append(transactions, *transaction)
	return r.SetSentTransactions(ctx, userID, transactions)
}
