package usecases

import (
	"avitoshop/internal/app/entities"
	"avitoshop/internal/app/repositories"
	"avitoshop/internal/app/usecases"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BuyMerchUseCase struct {
	dbPool             *pgxpool.Pool
	userRepo           repositories.UserRepository
	merchRepo          repositories.MerchRepository
	inventoryRepo      repositories.InventoryRepository
	redisUserRepo      repositories.RedisUserRepository
	redisMerchRepo     repositories.RedisMerchRepository
	redisInventoryRepo repositories.RedisInventoryRepository
}

func NewBuyMerchUseCase(
	dbPool *pgxpool.Pool,
	userRepo repositories.UserRepository,
	merchRepo repositories.MerchRepository,
	inventoryRepo repositories.InventoryRepository,
	redisUserRepo repositories.RedisUserRepository,
	redisMerchRepo repositories.RedisMerchRepository,
	redisInventoryRepo repositories.RedisInventoryRepository,
) *BuyMerchUseCase {
	return &BuyMerchUseCase{
		dbPool:             dbPool,
		userRepo:           userRepo,
		merchRepo:          merchRepo,
		inventoryRepo:      inventoryRepo,
		redisUserRepo:      redisUserRepo,
		redisMerchRepo:     redisMerchRepo,
		redisInventoryRepo: redisInventoryRepo,
	}
}

func (uc *BuyMerchUseCase) BuyMerch(
	ctx context.Context,
	username string,
	merchName string,
) error {
	user, err := uc.redisUserRepo.GetByUsername(ctx, username)

	if err != nil {
		if !errors.Is(err, repositories.ErrCacheMiss) {
			return fmt.Errorf("failed to get user from Redis: %w", err)
		}
		user, err = uc.userRepo.GetByUsername(ctx, username)
		if err != nil {
			return fmt.Errorf("failed to get user: %w", err)
		}
		if err := uc.redisUserRepo.SetByUsername(ctx, username, user); err != nil {
			return fmt.Errorf("failed to set user in Redis: %w", err)
		}
	}

	merch, err := uc.redisMerchRepo.GetByName(ctx, merchName)
	if err != nil {
		if !errors.Is(err, repositories.ErrCacheMiss) {
			return fmt.Errorf("failed to get user from Redis: %w", err)
		}
		merch, err = uc.merchRepo.GetByName(ctx, merchName)
		if err != nil {
			return fmt.Errorf("failed to get merch: %w", err)
		}
		if err := uc.redisMerchRepo.SetByName(ctx, merchName, merch); err != nil {
			return fmt.Errorf("failed to set merch in Redis: %w", err)
		}
	}

	if user.Balance < merch.Price {
		return usecases.ErrInsufficientFunds
	}

	user.Balance = user.Balance - merch.Price

	tx, err := uc.dbPool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	if err := uc.userRepo.UpdateBalance(ctx, user.ID, user.Balance); err != nil {
		return fmt.Errorf("failed to update user balance: %w", err)
	}

	inventoryItem := &entities.Inventory{
		UserID:   user.ID,
		Type:     merch.Name,
		Quantity: 1,
	}

	if err := uc.inventoryRepo.InsertOrUpdate(ctx, inventoryItem); err != nil {
		return fmt.Errorf("failed to update inventory: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	if err := uc.redisUserRepo.SetByUsername(ctx, user.Username, user); err != nil {
		return fmt.Errorf("failed to update user balance in Redis: %w", err)
	}

	//if err := uc.redisInventoryRepo.SetByUser(ctx, inventoryItem.UserID, inventoryItem); err != nil {
	//	return fmt.Errorf("failed to update inventory in Redis: %w", err)
	//}

	return nil
}
