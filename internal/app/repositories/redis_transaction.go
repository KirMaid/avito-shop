package repositories

//
//import (
//	"avitoshop/internal/app/entities"
//	"context"
//	"encoding/json"
//	"errors"
//	"fmt"
//	"github.com/redis/go-redis/v9"
//	"time"
//)
//
//type redisTransactionRepository struct {
//	client *redis.Client
//	ttl    time.Duration
//}
//
//func NewRedisTransactionRepository(client *redis.Client, ttl time.Duration) RedisTransactionRepository {
//	return &redisTransactionRepository{client: client, ttl: ttl}
//}
//
//func (r *redisTransactionRepository) GetReceivedTransactions(ctx context.Context, userID int) ([]entities.Transaction, error) {
//	key := fmt.Sprintf("transactions:received:%d", userID)
//	data, err := r.client.Get(ctx, key).Bytes()
//	if errors.Is(err, redis.Nil) {
//		return nil, ErrCacheMiss
//	} else if err != nil {
//		return nil, fmt.Errorf("failed to get received transactions from Redis: %w", err)
//	}
//
//	var transactions []entities.Transaction
//	if err := json.Unmarshal(data, &transactions); err != nil {
//		return nil, fmt.Errorf("failed to unmarshal transactions: %w", err)
//	}
//
//	return transactions, nil
//}
//
//func (r *redisTransactionRepository) SetReceivedTransactions(ctx context.Context, userID int, transactions []entities.Transaction) error {
//	key := fmt.Sprintf("transactions:received:%d", userID)
//	data, err := json.Marshal(transactions)
//	if err != nil {
//		return fmt.Errorf("failed to marshal transactions: %w", err)
//	}
//
//	if err := r.client.Set(ctx, key, data, r.ttl).Err(); err != nil {
//		return fmt.Errorf("failed to set received transactions in Redis: %w", err)
//	}
//
//	return nil
//}
//
//func (r *redisTransactionRepository) DeleteReceivedTransactions(ctx context.Context, userID int) error {
//	key := fmt.Sprintf("transactions:received:%d", userID)
//	if err := r.client.Del(ctx, key).Err(); err != nil {
//		return fmt.Errorf("failed to delete received transactions from Redis: %w", err)
//	}
//
//	return nil
//}
//
//// Аналогично для GetSentTransactions, SetSentTransactions, DeleteSentTransactions
//func (r *redisTransactionRepository) GetSentTransactions(ctx context.Context, userID int) ([]entities.Transaction, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (r *redisTransactionRepository) SetSentTransactions(ctx context.Context, userID int, transactions []entities.Transaction) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (r *redisTransactionRepository) DeleteSentTransactions(ctx context.Context, userID int) error {
//	//TODO implement me
//	panic("implement me")
//}
