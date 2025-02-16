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

var ErrFailedToGetGood = errors.New("failed to get good")

type BuyGoodUseCase struct {
	dbPool             *pgxpool.Pool
	userRepo           repositories.UserRepository
	goodRepo           repositories.GoodRepository
	inventoryRepo      repositories.InventoryRepository
	redisUserRepo      repositories.RedisUserRepository
	redisGoodRepo      repositories.RedisGoodRepository
	redisInventoryRepo repositories.RedisInventoryRepository
}

func NewBuyGoodUseCase(
	dbPool *pgxpool.Pool,
	userRepo repositories.UserRepository,
	goodRepo repositories.GoodRepository,
	inventoryRepo repositories.InventoryRepository,
	redisUserRepo repositories.RedisUserRepository,
	redisGoodRepo repositories.RedisGoodRepository,
	redisInventoryRepo repositories.RedisInventoryRepository,
) *BuyGoodUseCase {
	return &BuyGoodUseCase{
		dbPool:             dbPool,
		userRepo:           userRepo,
		goodRepo:           goodRepo,
		inventoryRepo:      inventoryRepo,
		redisUserRepo:      redisUserRepo,
		redisGoodRepo:      redisGoodRepo,
		redisInventoryRepo: redisInventoryRepo,
	}
}

func (uc *BuyGoodUseCase) BuyGood(
	ctx context.Context,
	username string,
	goodName string,
) error {
	user, err := uc.getUser(ctx, username)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	good, err := uc.getGood(ctx, goodName)
	if err != nil {
		return ErrFailedToGetGood
	}

	if user.Balance < good.Price {
		return usecases.ErrInsufficientFunds
	}

	user.Balance -= good.Price

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
		GoodID:   good.ID,
		Quantity: 1,
	}

	inventoryItem, err = uc.inventoryRepo.InsertOrUpdate(ctx, inventoryItem)

	if err != nil {
		return fmt.Errorf("failed to update inventory: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	defer uc.updateCache(ctx, user, inventoryItem)

	return nil
}

func (uc *BuyGoodUseCase) getUser(ctx context.Context, username string) (*entities.User, error) {
	user, err := uc.redisUserRepo.GetByUsername(ctx, username)
	if err == nil {
		return user, nil
	}

	user, err = uc.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from database: %w", err)
	}

	return user, nil
}

func (uc *BuyGoodUseCase) getGood(ctx context.Context, goodName string) (*entities.Good, error) {
	good, err := uc.redisGoodRepo.GetByName(ctx, goodName)
	if err == nil {
		return good, nil
	}

	good, err = uc.goodRepo.GetByName(ctx, goodName)
	if err != nil {
		return nil, fmt.Errorf("failed to get good from database: %w", err)
	}

	//go func() {
	_ = uc.redisGoodRepo.SetByName(ctx, goodName, good)
	_ = uc.redisGoodRepo.SetByID(ctx, good.ID, good)
	//}()

	return good, nil
}

func (uc *BuyGoodUseCase) updateCache(ctx context.Context, user *entities.User, inventoryItem *entities.Inventory) {
	//go func() {
	_ = uc.redisInventoryRepo.DeleteByUser(ctx, user.ID)
	_ = uc.redisUserRepo.SetByUsername(ctx, user.Username, user)
	_ = uc.redisUserRepo.SetById(ctx, user.ID, user)
	//}()
}
