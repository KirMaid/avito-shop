package usecases

import (
	"avitoshop/internal/app/repositories"
	"context"
	"fmt"
)

type UserInfoUseCase struct {
	userRepo        repositories.UserRepository
	inventoryRepo   repositories.InventoryRepository
	transactionRepo repositories.TransactionRepository
}

func NewUserInfoUseCase(
	userRepo repositories.UserRepository,
	inventoryRepo repositories.InventoryRepository,
	transactionRepo repositories.TransactionRepository,
) *UserInfoUseCase {
	return &UserInfoUseCase{
		userRepo:        userRepo,
		inventoryRepo:   inventoryRepo,
		transactionRepo: transactionRepo,
	}
}

func (s *UserInfoUseCase) GetInfo(ctx context.Context, username string) (*UserInfoDTO, error) {
	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	inventoryEntities, err := s.inventoryRepo.GetByUser(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get inventory: %w", err)
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

	received := make([]ReceivedDTO, 0, len(receivedEntities))
	for _, tx := range receivedEntities {
		received = append(received, ReceivedDTO{
			FromUser: tx.SenderID,
			Amount:   tx.Amount,
		})
	}

	sentEntities, err := s.transactionRepo.GetSentTransactions(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get sent transactions: %w", err)
	}

	sent := make([]SentDTO, 0, len(sentEntities))
	for _, tx := range sentEntities {
		sent = append(sent, SentDTO{
			ToUser: tx.ReceiverID,
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
