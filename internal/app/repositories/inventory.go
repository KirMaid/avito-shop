package repositories

import (
	"avitoshop/internal/app/entities"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type inventoryRepository struct {
	db *pgxpool.Pool
}

// TODO Реализовать
func (i inventoryRepository) Get(ctx context.Context, userID int) ([]entities.Inventory, error) {
	//TODO implement me
	panic("implement me")
}

type InventoryRepository interface {
	Get(ctx context.Context, userID int) ([]entities.Inventory, error)
}

func NewInventoryRepository(db *pgxpool.Pool) InventoryRepository {
	return &inventoryRepository{db: db}
}
