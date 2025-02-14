package usecases

import (
	"avitoshop/internal/app/entities"
	"avitoshop/internal/app/repositories"
	"context"
	"errors"
	"fmt"
)

type SendCoinsUseCase struct {
	userRepo        repositories.UserRepository
	transactionRepo repositories.TransactionRepository
	coinHistoryRepo repositories.CoinHistoryRepository
}

func NewSendCoinsUseCase(
	userRepo repositories.UserRepository,
	transactionRepo repositories.TransactionRepository,
	coinHistoryRepo repositories.CoinHistoryRepository,
) *SendCoinsUseCase {
	return &SendCoinsUseCase{
		userRepo:        userRepo,
		transactionRepo: transactionRepo,
		coinHistoryRepo: coinHistoryRepo,
	}
}

func (uc *SendCoinsUseCase) SendCoins(ctx context.Context, senderID int, toUsername string, amount int) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	sender, err := uc.userRepo.GetByID(ctx, senderID)
	if err != nil {
		return fmt.Errorf("failed to get sender: %w", err)
	}

	if sender.Balance < amount {
		return errors.New("insufficient funds")
	}

	receiver, err := uc.userRepo.GetByUsername(ctx, toUsername)
	if err != nil {
		return fmt.Errorf("failed to get receiver: %w", err)
	}

	transaction := &entities.Transaction{
		SenderID:   sender.ID,
		ReceiverID: receiver.ID,
		Amount:     amount,
	}
	if _, err := uc.transactionRepo.Insert(ctx, transaction); err != nil {
		return fmt.Errorf("failed to create transaction: %w", err)
	}

	newSenderBalance := sender.Balance - amount
	if err := uc.userRepo.UpdateBalance(ctx, sender.ID, newSenderBalance); err != nil {
		return fmt.Errorf("failed to update sender balance: %w", err)
	}

	newReceiverBalance := receiver.Balance + amount
	if err := uc.userRepo.UpdateBalance(ctx, receiver.ID, newReceiverBalance); err != nil {
		return fmt.Errorf("failed to update receiver balance: %w", err)
	}

	senderHistory := &entities.CoinHistory{
		UserID:        sender.ID,
		ChangeAmount:  -amount,
		OperationType: "send",
	}
	if _, err := uc.coinHistoryRepo.Insert(ctx, senderHistory); err != nil {
		return fmt.Errorf("failed to record sender history: %w", err)
	}

	// TODO receive вынести в константы в модели
	receiverHistory := &entities.CoinHistory{
		UserID:        receiver.ID,
		ChangeAmount:  amount,
		OperationType: "receive",
	}
	if _, err := uc.coinHistoryRepo.Insert(ctx, receiverHistory); err != nil {
		return fmt.Errorf("failed to record receiver history: %w", err)
	}

	return nil
}
