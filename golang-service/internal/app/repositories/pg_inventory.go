package repositories

import (
	"avitoshop/internal/app/entities"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type inventoryRepository struct {
	db *pgxpool.Pool
}

func NewInventoryRepository(db *pgxpool.Pool) InventoryRepository {
	return &inventoryRepository{db: db}
}

func (i *inventoryRepository) GetByUser(ctx context.Context, userID int) ([]entities.Inventory, error) {
	rows, err := i.db.Query(ctx, "SELECT * FROM inventories WHERE user_id = $1", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query received inventory: %w", err)
	}
	defer rows.Close()

	var inventory []entities.Inventory
	for rows.Next() {
		var inventoryItem entities.Inventory
		if err := rows.Scan(
			&inventoryItem.UserID,
			&inventoryItem.GoodID,
			&inventoryItem.Quantity,
		); err != nil {
			return nil, fmt.Errorf("failed to scan inventoryItem: %w", err)
		}
		inventory = append(inventory, inventoryItem)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating rows: %w", err)
	}

	return inventory, nil
}

func (i *inventoryRepository) InsertOrUpdate(ctx context.Context, inventory *entities.Inventory) (*entities.Inventory, error) {
	query := `
        INSERT INTO inventories (user_id, good_id, quantity)
        VALUES ($1, $2, $3)
        ON CONFLICT (user_id, good_id) DO UPDATE
        SET quantity = inventories.quantity + EXCLUDED.quantity
        WHERE inventories.quantity + EXCLUDED.quantity >= 0
	 	RETURNING quantity
	`

	err := i.db.QueryRow(ctx, query, inventory.UserID, inventory.GoodID, inventory.Quantity).Scan(
		&inventory.Quantity,
	)
	if err != nil {
		return nil, err
	}

	return inventory, nil
}
