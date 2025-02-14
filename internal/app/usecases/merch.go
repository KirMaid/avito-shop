package usecases

import "avitoshop/internal/app/repositories"

type MerchUseCase struct {
	merchRepo repositories.MerchRepository
	userRepo  repositories.UserRepository
}

func NewMerchUseCase(merchRepo repositories.MerchRepository, userRepo repositories.UserRepository) *MerchUseCase {
	return &MerchUseCase{merchRepo: merchRepo, userRepo: userRepo}
}
