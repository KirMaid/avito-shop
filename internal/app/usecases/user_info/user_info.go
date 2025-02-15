package usecases

import (
	"avitoshop/internal/app/entities"
	"avitoshop/internal/app/repositories"
	"context"
	"errors"
	"fmt"
)

type UserInfoUseCase struct {
	userRepo           repositories.UserRepository
	inventoryRepo      repositories.InventoryRepository
	transactionRepo    repositories.TransactionRepository
	redisUserRepo      repositories.RedisUserRepository
	redisInventoryRepo repositories.RedisInventoryRepository
	//redisTransactionRepo repositories.RedisTransactionRepository
}

func NewUserInfoUseCase(
	userRepo repositories.UserRepository,
	inventoryRepo repositories.InventoryRepository,
	transactionRepo repositories.TransactionRepository,
	redisUserRepo repositories.RedisUserRepository,
	redisInventoryRepo repositories.RedisInventoryRepository,
	// redisTransactionRepo repositories.RedisTransactionRepository,
) *UserInfoUseCase {
	return &UserInfoUseCase{
		userRepo:           userRepo,
		inventoryRepo:      inventoryRepo,
		transactionRepo:    transactionRepo,
		redisUserRepo:      redisUserRepo,
		redisInventoryRepo: redisInventoryRepo,
		//redisTransactionRepo: redisTransactionRepo,
	}
}

func (s *UserInfoUseCase) GetInfo(ctx context.Context, username string) (*UserInfoDTO, error) {
	user, err := s.redisUserRepo.GetByUsername(ctx, username)

	if err != nil {
		if !errors.Is(err, repositories.ErrCacheMiss) {
			return nil, fmt.Errorf("failed to get user from Redis: %w", err)
		}

		user, err = s.userRepo.GetByUsername(ctx, username)
		if err != nil {
			return nil, fmt.Errorf("failed to get user: %w", err)
		}

		if err := s.redisUserRepo.SetByUsername(ctx, username, user); err != nil {
			return nil, fmt.Errorf("failed to set user in Redis: %w", err)
		}
	}

	inventoryEntities, err := s.redisInventoryRepo.GetByUser(ctx, user.ID)
	if err != nil {
		if !errors.Is(err, repositories.ErrCacheMiss) {
			return nil, fmt.Errorf("failed to get user inventory from Redis: %w", err)
		}

		inventoryEntities, err = s.redisInventoryRepo.GetByUser(ctx, user.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get inventory: %w", err)
		}

		if err := s.redisInventoryRepo.SetByUser(ctx, user.ID, inventoryEntities); err != nil {
			return nil, fmt.Errorf("failed to set inventory in Redis: %w", err)
		}
	}

	inventory := make([]InventoryDTO, 0, len(inventoryEntities))
	for _, item := range inventoryEntities {
		inventory = append(inventory, InventoryDTO{
			Type:     item.Type,
			Quantity: item.Quantity,
		})
	}

	receivedEntities, err := s.transactionRepo.GetReceivedTransactions(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get received transactions: %w", err)
	}

	sentEntities, err := s.transactionRepo.GetSentTransactions(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get sent transactions: %w", err)
	}

	userIDSet := make(map[int]struct{})
	for _, tx := range receivedEntities {
		userIDSet[tx.SenderID] = struct{}{}
	}
	for _, tx := range sentEntities {
		userIDSet[tx.ReceiverID] = struct{}{}
	}

	userIDs := make([]int, 0, len(userIDSet))
	for id := range userIDSet {
		userIDs = append(userIDs, id)
	}

	userUsernames, err := s.redisUserRepo.GetUsernamesByIDs(ctx, userIDs)
	if err != nil {
		if !errors.Is(err, repositories.ErrCacheMiss) {
			return nil, fmt.Errorf("failed to get usernames from Redis: %w", err)
		}

		userUsernames, err = s.userRepo.GetUsernamesByIDs(ctx, userIDs)
		if err != nil {
			return nil, fmt.Errorf("failed to get usernames: %w", err)
		}

		for id, username := range userUsernames {
			if err := s.redisUserRepo.SetByUsername(ctx, username, &entities.User{ID: id, Username: username}); err != nil {
				return nil, fmt.Errorf("failed to set username in Redis: %w", err)
			}
		}
	}

	received := make([]ReceivedDTO, 0, len(receivedEntities))
	for _, tx := range receivedEntities {
		received = append(received, ReceivedDTO{
			FromUser: userUsernames[tx.SenderID],
			Amount:   tx.Amount,
		})
	}

	sent := make([]SentDTO, 0, len(sentEntities))
	for _, tx := range sentEntities {
		sent = append(sent, SentDTO{
			ToUser: userUsernames[tx.ReceiverID],
			Amount: tx.Amount,
		})
	}

	return &UserInfoDTO{
		Coins:     user.Balance,
		Inventory: inventory,
		CoinHistory: CoinHistoryDTO{
			Received: received,
			Sent:     sent,
		},
	}, nil
}
