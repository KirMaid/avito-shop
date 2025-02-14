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
}

func NewSendCoinsUseCase(
	userRepo repositories.UserRepository,
	transactionRepo repositories.TransactionRepository,
) *SendCoinsUseCase {
	return &SendCoinsUseCase{
		userRepo:        userRepo,
		transactionRepo: transactionRepo,
	}
}

func (uc *SendCoinsUseCase) SendCoins(
	ctx context.Context,
	senderUsername string,
	receiverUsername string,
	amount int,
) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	sender, err := uc.userRepo.GetByUsername(ctx, senderUsername)
	if err != nil {
		return fmt.Errorf("failed to get sender: %w", err)
	}

	if sender.Balance < amount {
		return errors.New("insufficient funds")
	}

	receiver, err := uc.userRepo.GetByUsername(ctx, receiverUsername)
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

	return nil
}
