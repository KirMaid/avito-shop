package usecases_test

//
//import (
//	"context"
//	"errors"
//	"go.uber.org/mock/gomock"
//	"testing"
//
//	"avitoshop/internal/app/entities"
//	"avitoshop/internal/app/repositories/mocks"
//	usecases "avitoshop/internal/app/usecases/buy_good"
//
//	"github.com/stretchr/testify/assert"
//)
//
//func TestBuyGoodUseCase_BuyGood_Success(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	userRepo := mocks.NewMockUserRepository(ctrl)
//	goodRepo := mocks.NewMockGoodRepository(ctrl)
//	inventoryRepo := mocks.NewMockInventoryRepository(ctrl)
//
//	uc := usecases.NewBuyGoodUseCase(userRepo, goodRepo, nil, inventoryRepo)
//
//	username := "testuser"
//	merchName := "testmerch"
//	user := &entities.User{
//		ID:      1,
//		Balance: 100,
//	}
//	good := &entities.Good{
//		Name:  merchName,
//		Price: 50,
//	}
//
//	userRepo.EXPECT().
//		GetByUsername(gomock.Any(), username).
//		Return(user, nil)
//
//	goodRepo.EXPECT().
//		GetByName(gomock.Any(), merchName).
//		Return(good, nil)
//
//	userRepo.EXPECT().
//		UpdateBalance(gomock.Any(), user.ID, user.Balance-good.Price).
//		Return(nil)
//
//	inventoryRepo.EXPECT().
//		InsertOrUpdate(gomock.Any(), gomock.Any()).
//		Return(nil)
//
//	err := uc.BuyGood(context.Background(), username, merchName)
//
//	assert.NoError(t, err)
//}
//
//func TestBuyMerchUseCase_BuyMerch_UserNotFound(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	userRepo := mocks.NewMockUserRepository(ctrl)
//	goodRepo := mocks.NewMockGoodRepository(ctrl)
//
//	uc := usecases.NewBuyMerchUseCase(userRepo, goodRepo, nil, nil)
//
//	username := "testuser"
//	merchName := "testmerch"
//
//	userRepo.EXPECT().
//		GetByUsername(gomock.Any(), username).
//		Return(nil, errors.New("user not found"))
//
//	err := uc.BuyMerch(context.Background(), username, merchName)
//
//	assert.Error(t, err)
//	assert.Contains(t, err.Error(), "failed to get user")
//}
//
//func TestBuyMerchUseCase_BuyMerch_MerchNotFound(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	userRepo := mocks.NewMockUserRepository(ctrl)
//	goodRepo := mocks.NewMockGoodRepository(ctrl)
//
//	uc := usecases.NewBuyMerchUseCase(userRepo, goodRepo, nil, nil)
//
//	username := "testuser"
//	goodName := "testmerch"
//	user := &entities.User{
//		ID:      1,
//		Balance: 100,
//	}
//
//	userRepo.EXPECT().
//		GetByUsername(gomock.Any(), username).
//		Return(user, nil)
//
//	goodRepo.EXPECT().
//		GetByName(gomock.Any(), goodName).
//		Return(nil, errors.New("good not found"))
//
//	err := uc.BuyMerch(context.Background(), username, goodName)
//
//	assert.Error(t, err)
//	assert.Contains(t, err.Error(), "failed to get merch")
//}
//
//func TestBuyMerchUseCase_BuyMerch_InsufficientFunds(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	userRepo := mocks.NewMockUserRepository(ctrl)
//	merchRepo := mocks.NewMockMerchRepository(ctrl)
//
//	uc := usecases.NewBuyMerchUseCase(userRepo, goodRepo, nil, nil)
//
//	username := "testuser"
//	merchName := "testmerch"
//	user := &entities.User{
//		ID:      1,
//		Balance: 30,
//	}
//	merch := &entities.Good{
//		Name:  merchName,
//		Price: 50,
//	}
//
//	userRepo.EXPECT().
//		GetByUsername(gomock.Any(), username).
//		Return(user, nil)
//
//	merchRepo.EXPECT().
//		GetByName(gomock.Any(), merchName).
//		Return(merch, nil)
//
//	err := uc.BuyMerch(context.Background(), username, merchName)
//
//	assert.Error(t, err)
//	assert.Equal(t, "insufficient funds", err.Error())
//}
//
//func TestBuyMerchUseCase_BuyMerch_UpdateBalanceError(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	userRepo := mocks.NewMockUserRepository(ctrl)
//	goodRepo := mocks.NewMockMerchRepository(ctrl)
//
//	uc := usecases.NewBuyMerchUseCase(userRepo, goodRepo, nil, nil)
//
//	username := "testuser"
//	goodName := "testmerch"
//	user := &entities.User{
//		ID:      1,
//		Balance: 100,
//	}
//	good := &entities.Good{
//		Name:  goodName,
//		Price: 50,
//	}
//
//	userRepo.EXPECT().
//		GetByUsername(gomock.Any(), username).
//		Return(user, nil)
//
//	goodRepo.EXPECT().
//		GetByName(gomock.Any(), goodName).
//		Return(good, nil)
//
//	userRepo.EXPECT().
//		UpdateBalance(gomock.Any(), user.ID, user.Balance-merch.Price).
//		Return(errors.New("update balance error"))
//
//	err := uc.BuyMerch(context.Background(), username, merchName)
//
//	assert.Error(t, err)
//	assert.Contains(t, err.Error(), "failed to update user balance")
//}
//
//func TestBuyMerchUseCase_BuyMerch_InventoryUpdateError(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	userRepo := mocks.NewMockUserRepository(ctrl)
//	merchRepo := mocks.NewMockMerchRepository(ctrl)
//	inventoryRepo := mocks.NewMockInventoryRepository(ctrl)
//
//	uc := usecases.NewBuyMerchUseCase(userRepo, merchRepo, nil, inventoryRepo)
//
//	username := "testuser"
//	merchName := "testmerch"
//	user := &entities.User{
//		ID:      1,
//		Balance: 100,
//	}
//	merch := &entities.Merch{
//		Name:  merchName,
//		Price: 50,
//	}
//
//	userRepo.EXPECT().
//		GetByUsername(gomock.Any(), username).
//		Return(user, nil)
//
//	merchRepo.EXPECT().
//		GetByName(gomock.Any(), merchName).
//		Return(merch, nil)
//
//	userRepo.EXPECT().
//		UpdateBalance(gomock.Any(), user.ID, user.Balance-merch.Price).
//		Return(nil)
//
//	inventoryRepo.EXPECT().
//		InsertOrUpdate(gomock.Any(), gomock.Any()).
//		Return(errors.New("inventory update error"))
//
//	err := uc.BuyMerch(context.Background(), username, merchName)
//
//	assert.Error(t, err)
//	assert.Contains(t, err.Error(), "failed to update inventory")
//}
