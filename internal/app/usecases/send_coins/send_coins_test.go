package usecases_test

import (
	"avitoshop/internal/app/entities"
	"avitoshop/internal/app/repositories"
	"avitoshop/internal/app/repositories/mocks"
	"avitoshop/internal/app/usecases"
	sendcoins "avitoshop/internal/app/usecases/send_coins"
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestSendCoinsUseCase_SendCoins_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mocks.NewMockUserRepository(ctrl)
	transactionRepo := mocks.NewMockTransactionRepository(ctrl)
	redisUserRepo := mocks.NewMockRedisUserRepository(ctrl)
	dbPool := &pgxpool.Pool{}

	uc := sendcoins.NewSendCoinsUseCase(dbPool, userRepo, transactionRepo, redisUserRepo, nil)

	senderUsername := "sender"
	receiverUsername := "receiver"
	amount := 100

	sender := &entities.User{
		ID:       1,
		Username: senderUsername,
		Balance:  200,
	}

	receiver := &entities.User{
		ID:       2,
		Username: receiverUsername,
		Balance:  50,
	}

	redisUserRepo.EXPECT().
		GetByUsername(gomock.Any(), senderUsername).
		Return(nil, repositories.ErrCacheMiss)

	userRepo.EXPECT().
		GetByUsername(gomock.Any(), senderUsername).
		Return(sender, nil)

	redisUserRepo.EXPECT().
		SetByUsername(gomock.Any(), senderUsername, sender).
		Return(nil)

	redisUserRepo.EXPECT().
		GetByUsername(gomock.Any(), receiverUsername).
		Return(nil, repositories.ErrCacheMiss)

	userRepo.EXPECT().
		GetByUsername(gomock.Any(), receiverUsername).
		Return(receiver, nil)

	redisUserRepo.EXPECT().
		SetByUsername(gomock.Any(), receiverUsername, receiver).
		Return(nil)

	transactionRepo.EXPECT().
		Insert(gomock.Any(), gomock.Any()).
		Return(&entities.Transaction{}, nil)

	userRepo.EXPECT().
		UpdateBalance(gomock.Any(), sender.ID, sender.Balance-amount).
		Return(nil)

	userRepo.EXPECT().
		UpdateBalance(gomock.Any(), receiver.ID, receiver.Balance+amount).
		Return(nil)

	redisUserRepo.EXPECT().
		SetByUsername(gomock.Any(), senderUsername, gomock.Any()).
		Return(nil)

	redisUserRepo.EXPECT().
		SetByUsername(gomock.Any(), receiverUsername, gomock.Any()).
		Return(nil)

	err := uc.SendCoins(context.Background(), senderUsername, receiverUsername, amount)

	assert.NoError(t, err)
}

func TestSendCoinsUseCase_SendCoins_InsufficientFunds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mocks.NewMockUserRepository(ctrl)
	transactionRepo := mocks.NewMockTransactionRepository(ctrl)
	redisUserRepo := mocks.NewMockRedisUserRepository(ctrl)
	dbPool := &pgxpool.Pool{}

	uc := sendcoins.NewSendCoinsUseCase(dbPool, userRepo, transactionRepo, redisUserRepo, nil)

	senderUsername := "sender"
	receiverUsername := "receiver"
	amount := 100

	sender := &entities.User{
		ID:       1,
		Username: senderUsername,
		Balance:  50,
	}

	redisUserRepo.EXPECT().
		GetByUsername(gomock.Any(), senderUsername).
		Return(nil, repositories.ErrCacheMiss)

	userRepo.EXPECT().
		GetByUsername(gomock.Any(), senderUsername).
		Return(sender, nil)

	redisUserRepo.EXPECT().
		SetByUsername(gomock.Any(), senderUsername, sender).
		Return(nil)

	err := uc.SendCoins(context.Background(), senderUsername, receiverUsername, amount)

	assert.Error(t, err)
	assert.Equal(t, usecases.ErrInsufficientFunds, err)
}

func TestSendCoinsUseCase_SendCoins_RedisError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mocks.NewMockUserRepository(ctrl)
	transactionRepo := mocks.NewMockTransactionRepository(ctrl)
	redisUserRepo := mocks.NewMockRedisUserRepository(ctrl)
	dbPool := &pgxpool.Pool{}

	uc := sendcoins.NewSendCoinsUseCase(dbPool, userRepo, transactionRepo, redisUserRepo, nil)

	senderUsername := "sender"
	receiverUsername := "receiver"
	amount := 100

	redisUserRepo.EXPECT().
		GetByUsername(gomock.Any(), senderUsername).
		Return(nil, errors.New("redis error"))

	err := uc.SendCoins(context.Background(), senderUsername, receiverUsername, amount)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get user(sender) from Redis")
}

func TestSendCoinsUseCase_SendCoins_TransactionError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mocks.NewMockUserRepository(ctrl)
	transactionRepo := mocks.NewMockTransactionRepository(ctrl)
	redisUserRepo := mocks.NewMockRedisUserRepository(ctrl)
	dbPool := &pgxpool.Pool{}

	uc := sendcoins.NewSendCoinsUseCase(dbPool, userRepo, transactionRepo, redisUserRepo, nil)

	senderUsername := "sender"
	receiverUsername := "receiver"
	amount := 100

	sender := &entities.User{
		ID:       1,
		Username: senderUsername,
		Balance:  200,
	}

	receiver := &entities.User{
		ID:       2,
		Username: receiverUsername,
		Balance:  50,
	}

	redisUserRepo.EXPECT().
		GetByUsername(gomock.Any(), senderUsername).
		Return(nil, repositories.ErrCacheMiss)

	userRepo.EXPECT().
		GetByUsername(gomock.Any(), senderUsername).
		Return(sender, nil)

	redisUserRepo.EXPECT().
		SetByUsername(gomock.Any(), senderUsername, sender).
		Return(nil)

	redisUserRepo.EXPECT().
		GetByUsername(gomock.Any(), receiverUsername).
		Return(nil, repositories.ErrCacheMiss)

	userRepo.EXPECT().
		GetByUsername(gomock.Any(), receiverUsername).
		Return(receiver, nil)

	redisUserRepo.EXPECT().
		SetByUsername(gomock.Any(), receiverUsername, receiver).
		Return(nil)

	transactionRepo.EXPECT().
		Insert(gomock.Any(), gomock.Any()).
		Return(nil, errors.New("transaction error"))

	err := uc.SendCoins(context.Background(), senderUsername, receiverUsername, amount)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create transaction")
}
