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

	sender, err := uc.getUser(ctx, senderUsername)
	if err != nil {
		return fmt.Errorf("failed to get sender: %w", err)
	}

	if sender.Balance < amount {
		return usecases.ErrInsufficientFunds
	}

	receiver, err := uc.getUser(ctx, receiverUsername)
	if err != nil {
		return fmt.Errorf("failed to get receiver: %w", err)
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

	sender.Balance -= amount
	if err := uc.userRepo.UpdateBalance(ctx, sender.ID, sender.Balance); err != nil {
		return fmt.Errorf("failed to update sender balance: %w", err)
	}

	receiver.Balance += amount
	if err := uc.userRepo.UpdateBalance(ctx, receiver.ID, receiver.Balance); err != nil {
		return fmt.Errorf("failed to update receiver balance: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	defer uc.updateCache(ctx, sender, receiver)

	return nil
}

func (uc *SendCoinsUseCase) getUser(ctx context.Context, username string) (*entities.User, error) {
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

func (uc *SendCoinsUseCase) updateCache(ctx context.Context, sender, receiver *entities.User) {
	//go func() {
	_ = uc.redisUserRepo.SetByUsername(ctx, sender.Username, sender)
	_ = uc.redisUserRepo.SetById(ctx, sender.ID, sender)
	_ = uc.redisUserRepo.SetByUsername(ctx, receiver.Username, receiver)
	_ = uc.redisUserRepo.SetById(ctx, receiver.ID, receiver)
	//}()
}
