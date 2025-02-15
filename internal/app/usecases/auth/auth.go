package usecases

import (
	"avitoshop/internal/app/entities"
	"avitoshop/internal/app/repositories"
	"avitoshop/pkg/jwt"
	"context"
	"errors"
	"fmt"
	"time"
)

var ErrInvalidPassword = errors.New("invalid password")

type AuthUseCase struct {
	userRepo       repositories.UserRepository
	redisUserRepo  repositories.RedisUserRepository
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthUseCase(
	userRepo repositories.UserRepository,
	redisUserRepo repositories.RedisUserRepository,
	hashSalt string,
	signingKey []byte,
	expireDuration time.Duration,
) *AuthUseCase {
	return &AuthUseCase{
		userRepo:       userRepo,
		redisUserRepo:  redisUserRepo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: expireDuration,
	}
}

func (auc *AuthUseCase) Auth(ctx context.Context, authData *entities.Auth) (string, error) {
	hashedPassword := jwt.HashPassword(authData.Password, auc.hashSalt)

	dbUser, err := auc.redisUserRepo.GetByUsername(ctx, authData.Username)

	if err != nil {
		if !errors.Is(err, repositories.ErrCacheMiss) && !errors.Is(err, repositories.ErrEmptyCacheData) {
			return "", fmt.Errorf("failed to get user from Redis: %w", err)
		}

		dbUser, err = auc.userRepo.GetByUsername(ctx, authData.Username)

		if err != nil {
			if errors.Is(err, repositories.ErrUserDoesNotExist) {
				return auc.Register(ctx, authData.Username, hashedPassword)
			}
		}

		if err := auc.redisUserRepo.SetByUsername(ctx, authData.Username, dbUser); err != nil {
			return "", fmt.Errorf("failed to set user in Redis: %w", err)
		}
	}

	fmt.Printf("Пароль жи есть из БД %s\n\n", dbUser.Password)
	fmt.Printf("Пароль жи есть хэшированный %s\n\n", hashedPassword)

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

	user, err = auc.userRepo.Insert(ctx, user)
	if err != nil {
		return "", err
	}

	if err := auc.redisUserRepo.SetByUsername(ctx, username, user); err != nil {
		return "", fmt.Errorf("failed to set user in Redis: %w", err)
	}
	return token, nil
}
