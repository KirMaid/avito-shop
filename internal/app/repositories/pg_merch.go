package repositories

import (
	"avitoshop/internal/app/entities"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type merchRepository struct {
	db *pgxpool.Pool
}

func NewMerchRepository(db *pgxpool.Pool) MerchRepository {
	return &merchRepository{db: db}
}

func (m merchRepository) GetList(ctx context.Context) ([]entities.Merch, error) {
	rows, err := m.db.Query(ctx, "SELECT * FROM goods")
	if err != nil {
		return nil, fmt.Errorf("failed to query received goods: %w", err)
	}
	defer rows.Close()
	var goods []entities.Merch
	for rows.Next() {
		var merch entities.Merch
		if err := rows.Scan(
			&merch.ID,
			&merch.Name,
			&merch.Price,
		); err != nil {
			return nil, fmt.Errorf("failed to scan merch: %w", err)
		}
		goods = append(goods, merch)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating rows: %w", err)
	}

	return goods, nil
}

func (m merchRepository) GetByName(ctx context.Context, name string) (*entities.Merch, error) {
	var merch entities.Merch
	query := "SELECT id, name, price FROM goods WHERE name = $1"
	err := m.db.QueryRow(ctx, query, name).Scan(&merch.ID, &merch.Name, &merch.Price)
	if err != nil {
		return nil, fmt.Errorf("failed to get merch by name: %w", err)
	}
	return &merch, nil
}
