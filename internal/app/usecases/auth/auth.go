package usecases

import (
	"avitoshop/internal/app/entities"
	"avitoshop/internal/app/repositories"
	"avitoshop/pkg/jwt"
	"context"
	"errors"
	"time"
)

var ErrInvalidPassword = errors.New("invalid password")

type AuthUseCase struct {
	userRepo repositories.UserRepository

	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthUseCase(userRepo repositories.UserRepository, hashSalt string, signingKey []byte, expireDuration time.Duration) *AuthUseCase {
	return &AuthUseCase{
		userRepo:       userRepo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: expireDuration,
	}
}

func (auc *AuthUseCase) Auth(ctx context.Context, authData *entities.Auth) (string, error) {
	hashedPassword := jwt.HashPassword(authData.Password, auc.hashSalt)

	dbUser, err := auc.userRepo.GetByUsername(ctx, authData.Username)
	if err != nil {
		if errors.Is(err, repositories.ErrUserDoesNotExist) {
			return auc.Register(ctx, authData.Username, hashedPassword)
		}
		return "", err
	}

	if dbUser.Password != hashedPassword {
		return "", ErrInvalidPassword
	}

	token, err := jwt.GenerateToken(dbUser.Username, auc.expireDuration, auc.signingKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (auc *AuthUseCase) Register(ctx context.Context, username string, hashedPassword string) (string, error) {
	user := &entities.User{
		Username: username,
		Password: hashedPassword,
	}

	token, err := jwt.GenerateToken(username, auc.expireDuration, auc.signingKey)
	if err != nil {
		return "", err
	}

	if err := auc.userRepo.Insert(ctx, user); err != nil {
		return "", err
	}
	return token, nil
}
