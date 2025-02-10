package usecase

import (
	"avitoshop/internal/app/auth"
	"avitoshop/internal/app/entity"
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Authorizer struct {
	repo auth.Repository

	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthorizer(repo auth.Repository, hashSalt string, signingKey []byte, expireDuration time.Duration) *Authorizer {
	return &Authorizer{
		repo:           repo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: expireDuration,
	}
}

// TODO Посмотреть по поводу контекста
func (a *Authorizer) Auth(ctx context.Context, user *entity.User) (string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(user.Password))
	pwd.Write([]byte(a.hashSalt))
	hashedPassword := fmt.Sprintf("%x", pwd.Sum(nil))

	dbUser, err := a.repo.Get(ctx, user.Username)

	if err != nil {
		if errors.Is(err, auth.ErrUserDoesNotExist) {
			user.Password = hashedPassword
			if err := a.repo.Insert(ctx, user); err != nil {
				return "", err
			}
			dbUser = user
		} else {
			return "", err
		}
	}

	if dbUser.Password != hashedPassword {
		return "", auth.ErrInvalidPassword
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &auth.Claims{
		//TODO Посмотреть по поводу остальных полей
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(a.expireDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Username: dbUser.Username,
	})

	return token.SignedString(a.signingKey)
}
