package repositories

import (
	"avitoshop/internal/app/entities"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type coinHistoryRepository struct {
	db *pgxpool.Pool
}

type CoinHistoryRepository interface {
	GetByUser(ctx context.Context, userID int) ([]entities.CoinHistory, error)
	Insert(ctx context.Context, coinHistory *entities.CoinHistory) (int, error)
}

func NewCoinHistoryRepository(db *pgxpool.Pool) CoinHistoryRepository {
	return &coinHistoryRepository{db: db}
}

func (r *coinHistoryRepository) Insert(ctx context.Context, coinHistory *entities.CoinHistory) (int, error) {
	query := "INSERT INTO coin_history (user_id, change_amount, operation_type) VALUES ($1, $2, $3) RETURNING id"

	var id int
	err := r.db.QueryRow(ctx, query, coinHistory.UserID, coinHistory.ChangeAmount, coinHistory.OperationType).Scan(&id)
	if err != nil {
		// log.Errorf("error on inserting transaction: %s", err.Error())
		return 0, err
	}

	return id, nil
}

func (r *coinHistoryRepository) GetByUser(ctx context.Context, userID int) ([]entities.CoinHistory, error) {
	rows, err := r.db.Query(ctx, "SELECT * FROM coin_history WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var coinHistories []entities.CoinHistory

	for rows.Next() {
		var coinHistory entities.CoinHistory
		err := rows.Scan(
			&coinHistory.ID,
			&coinHistory.UserID,
			&coinHistory.ChangeAmount,
			&coinHistory.CreatedAt,
			&coinHistory.OperationType,
		)
		if err != nil {
			return nil, err
		}
		coinHistories = append(coinHistories, coinHistory)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return coinHistories, nil
}
