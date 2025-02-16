package usecases_test

import (
	"avitoshop/internal/app/entities"
	"avitoshop/internal/app/repositories/mocks"
	usecases "avitoshop/internal/app/usecases/user_info"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUserInfoUseCase_GetInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mocks.NewMockUserRepository(ctrl)
	inventoryRepo := mocks.NewMockInventoryRepository(ctrl)
	transactionRepo := mocks.NewMockTransactionRepository(ctrl)
	goodRepo := mocks.NewMockGoodRepository(ctrl)
	redisUserRepo := mocks.NewMockRedisUserRepository(ctrl)
	redisInventoryRepo := mocks.NewMockRedisInventoryRepository(ctrl)
	redisGoodRepo := mocks.NewMockRedisGoodRepository(ctrl)
	redisTransactionRepo := mocks.NewMockRedisTransactionRepository(ctrl)

	uc := usecases.NewUserInfoUseCase(
		userRepo,
		inventoryRepo,
		transactionRepo,
		goodRepo,
		redisUserRepo,
		redisInventoryRepo,
		redisGoodRepo,
		redisTransactionRepo,
	)

	username := "testuser"
	userID := 1
	balance := 100
	user := &entities.User{
		ID:      userID,
		Balance: balance,
	}
	inventoryItems := []entities.Inventory{
		{GoodID: 1, Quantity: 2},
		{GoodID: 2, Quantity: 1},
	}
	goods := map[int]*entities.Good{
		1: {ID: 1, Name: "item1"},
		2: {ID: 2, Name: "item2"},
	}
	receivedTransactions := []entities.Transaction{
		{SenderID: 2, Amount: 50},
		{SenderID: 3, Amount: 30},
	}
	sentTransactions := []entities.Transaction{
		{ReceiverID: 2, Amount: 20},
		{ReceiverID: 3, Amount: 10},
	}
	usernames := map[int]string{
		2: "user2",
		3: "user3",
	}

	redisUserRepo.EXPECT().
		GetByUsername(gomock.Any(), username).
		Return(nil, fmt.Errorf("not found in cache")) // Кэш пустой
	userRepo.EXPECT().
		GetByUsername(gomock.Any(), username).
		Return(user, nil)
	redisUserRepo.EXPECT().
		SetByUsername(gomock.Any(), username, user).
		Return(nil)
	redisUserRepo.EXPECT().
		SetById(gomock.Any(), userID, user).
		Return(nil)

	redisInventoryRepo.EXPECT().
		GetByUser(gomock.Any(), userID).
		Return(nil, fmt.Errorf("not found in cache"))
	inventoryRepo.EXPECT().
		GetByUser(gomock.Any(), userID).
		Return(inventoryItems, nil)
	redisInventoryRepo.EXPECT().
		SetByUser(gomock.Any(), userID, inventoryItems).
		Return(nil)

	for _, item := range inventoryItems {
		redisGoodRepo.EXPECT().
			GetByID(gomock.Any(), item.GoodID).
			Return(nil, fmt.Errorf("not found in cache")) // Кэш пустой
		goodRepo.EXPECT().
			GetByID(gomock.Any(), item.GoodID).
			Return(goods[item.GoodID], nil)
		redisGoodRepo.EXPECT().
			SetByID(gomock.Any(), item.GoodID, goods[item.GoodID]).
			Return(nil)
	}

	redisTransactionRepo.EXPECT().
		GetReceivedTransactions(gomock.Any(), userID).
		Return(nil, fmt.Errorf("not found in cache")) // Кэш пустой
	transactionRepo.EXPECT().
		GetReceivedTransactions(gomock.Any(), userID).
		Return(receivedTransactions, nil)
	redisTransactionRepo.EXPECT().
		SetReceivedTransactions(gomock.Any(), userID, receivedTransactions).
		Return(nil)

	redisTransactionRepo.EXPECT().
		GetSentTransactions(gomock.Any(), userID).
		Return(nil, fmt.Errorf("not found in cache")) // Кэш пустой
	transactionRepo.EXPECT().
		GetSentTransactions(gomock.Any(), userID).
		Return(sentTransactions, nil)
	redisTransactionRepo.EXPECT().
		SetSentTransactions(gomock.Any(), userID, sentTransactions).
		Return(nil)

	userRepo.EXPECT().
		GetUsernamesByIDs(gomock.Any(), gomock.Any()).
		Return(usernames, nil)

	info, err := uc.GetInfo(context.Background(), username)

	assert.NoError(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, balance, info.Coins)
	assert.Len(t, info.Inventory, 2)
	assert.Equal(t, "item1", info.Inventory[0].Type)
	assert.Equal(t, 2, info.Inventory[0].Quantity)
	assert.Equal(t, "item2", info.Inventory[1].Type)
	assert.Equal(t, 1, info.Inventory[1].Quantity)
	assert.Len(t, info.CoinHistory.Received, 2)
	assert.Equal(t, "user2", info.CoinHistory.Received[0].FromUser)
	assert.Equal(t, 50, info.CoinHistory.Received[0].Amount)
	assert.Equal(t, "user3", info.CoinHistory.Received[1].FromUser)
	assert.Equal(t, 30, info.CoinHistory.Received[1].Amount)
	assert.Len(t, info.CoinHistory.Sent, 2)
	assert.Equal(t, "user2", info.CoinHistory.Sent[0].ToUser)
	assert.Equal(t, 20, info.CoinHistory.Sent[0].Amount)
	assert.Equal(t, "user3", info.CoinHistory.Sent[1].ToUser)
	assert.Equal(t, 10, info.CoinHistory.Sent[1].Amount)
}

func TestUserInfoUseCase_GetInfo_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mocks.NewMockUserRepository(ctrl)
	inventoryRepo := mocks.NewMockInventoryRepository(ctrl)
	transactionRepo := mocks.NewMockTransactionRepository(ctrl)
	goodRepo := mocks.NewMockGoodRepository(ctrl)
	redisUserRepo := mocks.NewMockRedisUserRepository(ctrl)
	redisInventoryRepo := mocks.NewMockRedisInventoryRepository(ctrl)
	redisGoodRepo := mocks.NewMockRedisGoodRepository(ctrl)
	redisTransactionRepo := mocks.NewMockRedisTransactionRepository(ctrl)

	uc := usecases.NewUserInfoUseCase(
		userRepo,
		inventoryRepo,
		transactionRepo,
		goodRepo,
		redisUserRepo,
		redisInventoryRepo,
		redisGoodRepo,
		redisTransactionRepo,
	)

	username := "unknownuser"

	redisUserRepo.EXPECT().
		GetByUsername(gomock.Any(), username).
		Return(nil, fmt.Errorf("not found in cache"))
	userRepo.EXPECT().
		GetByUsername(gomock.Any(), username).
		Return(nil, fmt.Errorf("user not found"))

	info, err := uc.GetInfo(context.Background(), username)

	assert.Error(t, err)
	assert.Nil(t, info)
	assert.Contains(t, err.Error(), "failed to get user")
}

func TestUserInfoUseCase_GetInfo_InventoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mocks.NewMockUserRepository(ctrl)
	inventoryRepo := mocks.NewMockInventoryRepository(ctrl)
	transactionRepo := mocks.NewMockTransactionRepository(ctrl)
	goodRepo := mocks.NewMockGoodRepository(ctrl)
	redisUserRepo := mocks.NewMockRedisUserRepository(ctrl)
	redisInventoryRepo := mocks.NewMockRedisInventoryRepository(ctrl)
	redisGoodRepo := mocks.NewMockRedisGoodRepository(ctrl)
	redisTransactionRepo := mocks.NewMockRedisTransactionRepository(ctrl)

	uc := usecases.NewUserInfoUseCase(
		userRepo,
		inventoryRepo,
		transactionRepo,
		goodRepo,
		redisUserRepo,
		redisInventoryRepo,
		redisGoodRepo,
		redisTransactionRepo,
	)

	username := "testuser"
	userID := 1
	balance := 100

	user := &entities.User{
		ID:      userID,
		Balance: balance,
	}

	redisUserRepo.EXPECT().
		GetByUsername(gomock.Any(), username).
		Return(nil, fmt.Errorf("not found in cache"))
	userRepo.EXPECT().
		GetByUsername(gomock.Any(), username).
		Return(user, nil)
	redisUserRepo.EXPECT().
		SetByUsername(gomock.Any(), username, user).
		Return(nil)
	redisUserRepo.EXPECT().
		SetById(gomock.Any(), userID, user).
		Return(nil)

	redisInventoryRepo.EXPECT().
		GetByUser(gomock.Any(), userID).
		Return(nil, fmt.Errorf("not found in cache"))
	inventoryRepo.EXPECT().
		GetByUser(gomock.Any(), userID).
		Return(nil, fmt.Errorf("database error"))

	info, err := uc.GetInfo(context.Background(), username)

	assert.Error(t, err)
	assert.Nil(t, info)
	assert.Contains(t, err.Error(), "failed to get inventory")
}
