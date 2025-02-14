package usecases_test

import (
	"avitoshop/internal/app/entities"
	"avitoshop/internal/app/repositories/mocks"
	usecase "avitoshop/internal/app/usecases/user_info"
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

	uc := usecase.NewUserInfoUseCase(userRepo, inventoryRepo, transactionRepo)

	username := "testuser"
	userID := 1
	balance := 100

	user := &entities.User{
		ID:      userID,
		Balance: balance,
	}

	inventoryItems := []*entities.Inventory{
		{Type: "item1", Quantity: 2},
		{Type: "item2", Quantity: 1},
	}

	receivedTransactions := []*entities.Transaction{
		{SenderID: 2, Amount: 50},
		{SenderID: 3, Amount: 30},
	}

	sentTransactions := []*entities.Transaction{
		{ReceiverID: 2, Amount: 20},
		{ReceiverID: 3, Amount: 10},
	}

	userRepo.EXPECT().
		GetByUsername(gomock.Any(), username).
		Return(user, nil)

	inventoryRepo.EXPECT().
		GetByUser(gomock.Any(), userID).
		Return(inventoryItems, nil)

	transactionRepo.EXPECT().
		GetReceivedTransactions(gomock.Any(), userID).
		Return(receivedTransactions, nil)

	transactionRepo.EXPECT().
		GetSentTransactions(gomock.Any(), userID).
		Return(sentTransactions, nil)

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
	assert.Equal(t, 2, info.CoinHistory.Received[0].FromUser)
	assert.Equal(t, 50.0, info.CoinHistory.Received[0].Amount)
	assert.Equal(t, 3, info.CoinHistory.Received[1].FromUser)
	assert.Equal(t, 30.0, info.CoinHistory.Received[1].Amount)

	assert.Len(t, info.CoinHistory.Sent, 2)
	assert.Equal(t, 2, info.CoinHistory.Sent[0].ToUser)
	assert.Equal(t, 20.0, info.CoinHistory.Sent[0].Amount)
	assert.Equal(t, 3, info.CoinHistory.Sent[1].ToUser)
	assert.Equal(t, 10.0, info.CoinHistory.Sent[1].Amount)
}

func TestUserInfoUseCase_GetInfo_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mocks.NewMockUserRepository(ctrl)
	inventoryRepo := mocks.NewMockInventoryRepository(ctrl)
	transactionRepo := mocks.NewMockTransactionRepository(ctrl)

	uc := usecase.NewUserInfoUseCase(userRepo, inventoryRepo, transactionRepo)

	username := "unknownuser"

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

	uc := usecase.NewUserInfoUseCase(userRepo, inventoryRepo, transactionRepo)

	username := "testuser"
	userID := 1
	balance := 100

	user := &entities.User{
		ID:      userID,
		Balance: balance,
	}

	userRepo.EXPECT().
		GetByUsername(gomock.Any(), username).
		Return(user, nil)

	inventoryRepo.EXPECT().
		GetByUser(gomock.Any(), userID).
		Return(nil, fmt.Errorf("database error"))

	info, err := uc.GetInfo(context.Background(), username)

	assert.Error(t, err)
	assert.Nil(t, info)
	assert.Contains(t, err.Error(), "failed to get inventory")
}
