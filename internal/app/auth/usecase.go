package auth

import (
	"avitoshop/internal/app/entity"
	"context"
)

type UseCase interface {
	Auth(ctx context.Context, user *entity.User) (string, error)
}
