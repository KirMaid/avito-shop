package repositories

import (
	"avitoshop/internal/app/entities"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type goodRepository struct {
	db *pgxpool.Pool
}

func (g goodRepository) GetByID(ctx context.Context, id int) (*entities.Good, error) {
	var good entities.Good
	err := g.db.QueryRow(ctx, "SELECT * FROM goods WHERE id = $1", id).Scan(&good.ID, &good.Name, &good.Price)
	if err != nil {
		return nil, err
	}
	return &good, nil
}

func NewGoodRepository(db *pgxpool.Pool) GoodRepository {
	return &goodRepository{db: db}
}

func (g goodRepository) GetList(ctx context.Context) ([]entities.Good, error) {
	rows, err := g.db.Query(ctx, "SELECT * FROM goods")
	if err != nil {
		return nil, fmt.Errorf("failed to query received goods: %w", err)
	}
	defer rows.Close()
	var goods []entities.Good
	for rows.Next() {
		var good entities.Good
		if err := rows.Scan(
			&good.ID,
			&good.Name,
			&good.Price,
		); err != nil {
			return nil, fmt.Errorf("failed to scan good: %w", err)
		}
		goods = append(goods, good)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating rows: %w", err)
	}

	return goods, nil
}

func (g goodRepository) GetByName(ctx context.Context, name string) (*entities.Good, error) {
	var good entities.Good
	query := "SELECT id, name, price FROM goods WHERE name = $1"
	err := g.db.QueryRow(ctx, query, name).Scan(&good.ID, &good.Name, &good.Price)
	if err != nil {
		return nil, fmt.Errorf("failed to get good by name: %w", err)
	}
	return &good, nil
}
