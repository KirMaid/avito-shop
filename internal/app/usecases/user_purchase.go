package usecases

import "avitoshop/internal/app/repositories"

type UserPurchaseUseCase struct {
	userRepo        repositories.UserRepository
	transactionRepo repositories.TransactionRepository
	coinHistoryRepo repositories.CoinHistoryRepository
}

func NewUserPurchaseUseCase(
	userRepo repositories.UserRepository,
	transactionRepo repositories.TransactionRepository,
	coinHistoryRepo repositories.CoinHistoryRepository,
) *UserPurchaseUseCase {
	return &UserPurchaseUseCase{
		userRepo:        userRepo,
		transactionRepo: transactionRepo,
		coinHistoryRepo: coinHistoryRepo,
	}
}
