package auth

import (
	"avitoshop/internal/app/entity"
	"context"
)

type Repository interface {
	Insert(ctx context.Context, user *entity.User) error
	Get(ctx context.Context, username string) (*entity.User, error)
}
