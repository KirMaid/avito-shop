package usecases

import (
	"avitoshop/internal/app/entities"
	"avitoshop/internal/app/repositories"
	"context"
	"errors"
	"fmt"
)

type BuyMerchUseCase struct {
	userRepo        repositories.UserRepository
	merchRepo       repositories.MerchRepository
	transactionRepo repositories.TransactionRepository
	inventoryRepo   repositories.InventoryRepository
}

func NewBuyMerchUseCase(
	userRepo repositories.UserRepository,
	merchRepo repositories.MerchRepository,
	transactionRepo repositories.TransactionRepository,
	inventoryRepo repositories.InventoryRepository,
) *BuyMerchUseCase {
	return &BuyMerchUseCase{
		userRepo:        userRepo,
		merchRepo:       merchRepo,
		transactionRepo: transactionRepo,
		inventoryRepo:   inventoryRepo,
	}
}

func (uc *BuyMerchUseCase) BuyMerch(
	ctx context.Context,
	username string,
	merchName string,
) error {
	user, err := uc.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	merch, err := uc.merchRepo.GetByName(ctx, merchName)
	if err != nil {
		return fmt.Errorf("failed to get merch: %w", err)
	}

	if user.Balance < merch.Price {
		return errors.New("insufficient funds")
	}

	newBalance := user.Balance - merch.Price
	if err := uc.userRepo.UpdateBalance(ctx, user.ID, newBalance); err != nil {
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

	return nil
}
