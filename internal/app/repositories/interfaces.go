package repositories

import (
	"avitoshop/internal/app/entities"
	"context"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks/mocks.go -package=mocks

type InventoryRepository interface {
	GetByUser(ctx context.Context, userID int) ([]entities.Inventory, error)
	InsertOrUpdate(ctx context.Context, inventory *entities.Inventory) error
}

type RedisInventoryRepository interface {
	GetByUser(ctx context.Context, userID int) ([]entities.Inventory, error)
	SetByUser(ctx context.Context, userID int, inventory []entities.Inventory) error
	DeleteByUser(ctx context.Context, userID int) error
}

type MerchRepository interface {
	GetList(ctx context.Context) ([]entities.Merch, error)
	GetByName(ctx context.Context, name string) (*entities.Merch, error)
}

type RedisMerchRepository interface {
	GetByName(ctx context.Context, name string) (*entities.Merch, error)
	SetByName(ctx context.Context, name string, merch *entities.Merch) error
}

type TransactionRepository interface {
	Insert(ctx context.Context, transaction *entities.Transaction) (*entities.Transaction, error)
	GetReceivedTransactions(ctx context.Context, userID int) ([]entities.Transaction, error)
	GetSentTransactions(ctx context.Context, userID int) ([]entities.Transaction, error)
}

type RedisTransactionRepository interface {
	GetReceivedTransactions(ctx context.Context, userID int) ([]entities.Transaction, error)
	SetReceivedTransactions(ctx context.Context, userID int, transactions []entities.Transaction) error
	DeleteReceivedTransactions(ctx context.Context, userID int) error

	GetSentTransactions(ctx context.Context, userID int) ([]entities.Transaction, error)
	SetSentTransactions(ctx context.Context, userID int, transactions []entities.Transaction) error
	DeleteSentTransactions(ctx context.Context, userID int) error
}

type UserRepository interface {
	Insert(ctx context.Context, user *entities.User) (*entities.User, error)
	GetByUsername(ctx context.Context, username string) (*entities.User, error)
	GetByID(ctx context.Context, userID int) (*entities.User, error)
	GetUsernamesByIDs(ctx context.Context, userIDs []int) (map[int]string, error)
	UpdateBalance(ctx context.Context, userID int, balance int) error
}

type RedisUserRepository interface {
	GetByUsername(ctx context.Context, username string) (*entities.User, error)
	SetByUsername(ctx context.Context, username string, user *entities.User) error
	DeleteByUsername(ctx context.Context, username string) error
	GetUsernamesByIDs(ctx context.Context, userIDs []int) (map[int]string, error)
}
