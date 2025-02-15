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

type SendCoinsUseCase struct {
	dbPool          *pgxpool.Pool
	userRepo        repositories.UserRepository
	transactionRepo repositories.TransactionRepository
	redisUserRepo   repositories.RedisUserRepository
	//redisTransactionRepo repositories.RedisTransactionRepository
}

func NewSendCoinsUseCase(
	dbPool *pgxpool.Pool,
	userRepo repositories.UserRepository,
	transactionRepo repositories.TransactionRepository,
	redisUserRepo repositories.RedisUserRepository,
	// redisTransactionRepo repositories.RedisTransactionRepository,
) *SendCoinsUseCase {
	return &SendCoinsUseCase{
		dbPool:          dbPool,
		userRepo:        userRepo,
		transactionRepo: transactionRepo,
		redisUserRepo:   redisUserRepo,
		//redisTransactionRepo: redisTransactionRepo,
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

	sender, err := uc.redisUserRepo.GetByUsername(ctx, senderUsername)

	if err != nil {
		if !errors.Is(err, repositories.ErrCacheMiss) {
			return fmt.Errorf("failed to get user(sender) from Redis: %w", err)
		}
		sender, err = uc.userRepo.GetByUsername(ctx, senderUsername)
		if err != nil {
			return fmt.Errorf("failed to get user(sender): %w", err)
		}
		if err := uc.redisUserRepo.SetByUsername(ctx, senderUsername, sender); err != nil {
			return fmt.Errorf("failed to set user(sender) in Redis: %w", err)
		}
	}

	if sender.Balance < amount {
		return usecases.ErrInsufficientFunds
	}

	receiver, err := uc.redisUserRepo.GetByUsername(ctx, receiverUsername)
	if err != nil {
		if !errors.Is(err, repositories.ErrCacheMiss) {
			return fmt.Errorf("failed to get user(receiver) from Redis: %w", err)
		}
		receiver, err = uc.userRepo.GetByUsername(ctx, receiverUsername)
		if err != nil {
			return fmt.Errorf("failed to get user(sender): %w", err)
		}
		if err := uc.redisUserRepo.SetByUsername(ctx, receiverUsername, receiver); err != nil {
			return fmt.Errorf("failed to set user(sender) in Redis: %w", err)
		}
	}

	transaction := &entities.Transaction{
		SenderID:   sender.ID,
		ReceiverID: receiver.ID,
		Amount:     amount,
	}

	tx, err := uc.dbPool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	if _, err := uc.transactionRepo.Insert(ctx, transaction); err != nil {
		return fmt.Errorf("failed to create transaction: %w", err)
	}

	sender.Balance = sender.Balance - amount
	if err := uc.userRepo.UpdateBalance(ctx, sender.ID, sender.Balance); err != nil {
		return fmt.Errorf("failed to update sender balance: %w", err)
	}

	receiver.Balance = receiver.Balance + amount
	if err := uc.userRepo.UpdateBalance(ctx, receiver.ID, receiver.Balance); err != nil {
		return fmt.Errorf("failed to update receiver balance: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	if err := uc.redisUserRepo.SetByUsername(ctx, sender.Username, sender); err != nil {
		return fmt.Errorf("failed to set user(sender) in Redis: %w", err)
	}

	if err := uc.redisUserRepo.SetByUsername(ctx, receiver.Username, receiver); err != nil {
		return fmt.Errorf("failed to set user(receiver) in Redis: %w", err)
	}

	return nil
}
