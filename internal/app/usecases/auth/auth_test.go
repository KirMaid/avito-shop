package usecases_test

import (
	"avitoshop/internal/app/entities"
	"avitoshop/internal/app/repositories"
	"avitoshop/internal/app/repositories/mocks"
	usecases "avitoshop/internal/app/usecases/auth"
	"avitoshop/pkg/jwt"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestAuthUseCase_Auth_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mocks.NewMockUserRepository(ctrl)

	hashSalt := "salt"
	signingKey := []byte("secret")
	expireDuration := time.Hour
	auc := usecases.NewAuthUseCase(userRepo, hashSalt, signingKey, expireDuration)

	username := "testuser"
	password := "testpassword"
	hashedPassword := jwt.HashPassword(password, hashSalt)

	dbUser := &entities.User{
		Username: username,
		Password: hashedPassword,
	}

	userRepo.EXPECT().
		GetByUsername(gomock.Any(), username).
		Return(dbUser, nil)

	token, err := auc.Auth(context.Background(), &entities.Auth{
		Username: username,
		Password: password,
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	validUsername, err := jwt.ValidateToken(token, signingKey)
	assert.NoError(t, err)
	assert.Equal(t, username, validUsername)
}

func TestAuthUseCase_Auth_InvalidPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mocks.NewMockUserRepository(ctrl)

	hashSalt := "salt"
	signingKey := []byte("secret")
	expireDuration := time.Hour
	auc := usecases.NewAuthUseCase(userRepo, hashSalt, signingKey, expireDuration)

	username := "testuser"
	password := "testpassword"
	wrongPassword := "wrongpassword"
	hashedPassword := jwt.HashPassword(password, hashSalt)

	dbUser := &entities.User{
		Username: username,
		Password: hashedPassword,
	}

	userRepo.EXPECT().
		GetByUsername(gomock.Any(), username).
		Return(dbUser, nil)

	token, err := auc.Auth(context.Background(), &entities.Auth{
		Username: username,
		Password: wrongPassword,
	})

	assert.Error(t, err)
	assert.Equal(t, usecases.ErrInvalidPassword, err)
	assert.Empty(t, token)
}

func TestAuthUseCase_Auth_UserNotFound_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mocks.NewMockUserRepository(ctrl)

	hashSalt := "salt"
	signingKey := []byte("secret")
	expireDuration := time.Hour
	auc := usecases.NewAuthUseCase(userRepo, hashSalt, signingKey, expireDuration)

	username := "newuser"
	password := "newpassword"
	//hashedPassword := jwt.HashPassword(password, hashSalt)

	userRepo.EXPECT().
		GetByUsername(gomock.Any(), username).
		Return(nil, repositories.ErrUserDoesNotExist)

	userRepo.EXPECT().
		Insert(gomock.Any(), gomock.Any()).
		Return(nil)

	token, err := auc.Auth(context.Background(), &entities.Auth{
		Username: username,
		Password: password,
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	validUsername, err := jwt.ValidateToken(token, signingKey)
	assert.NoError(t, err)
	assert.Equal(t, username, validUsername)
}

func TestAuthUseCase_Register_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mocks.NewMockUserRepository(ctrl)

	hashSalt := "salt"
	signingKey := []byte("secret")
	expireDuration := time.Hour
	auc := usecases.NewAuthUseCase(userRepo, hashSalt, signingKey, expireDuration)

	username := "newuser"
	password := "newpassword"
	hashedPassword := jwt.HashPassword(password, hashSalt)

	userRepo.EXPECT().
		Insert(gomock.Any(), gomock.Any()).
		Return(nil)

	token, err := auc.Register(context.Background(), username, hashedPassword)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	validUsername, err := jwt.ValidateToken(token, signingKey)
	assert.NoError(t, err)
	assert.Equal(t, username, validUsername)
}

func TestAuthUseCase_Register_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mocks.NewMockUserRepository(ctrl)

	hashSalt := "salt"
	signingKey := []byte("secret")
	expireDuration := time.Hour
	auc := usecases.NewAuthUseCase(userRepo, hashSalt, signingKey, expireDuration)

	username := "newuser"
	password := "newpassword"
	hashedPassword := jwt.HashPassword(password, hashSalt)

	expectedErr := errors.New("database error")
	userRepo.EXPECT().
		Insert(gomock.Any(), gomock.Any()).
		Return(expectedErr)

	token, err := auc.Register(context.Background(), username, hashedPassword)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Empty(t, token)
}
