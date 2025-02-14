package repositories

import "github.com/jackc/pgx/v5/pgxpool"

type MerchRepository struct {
	db *pgxpool.Pool
}

func NewMerchRepository(db *pgxpool.Pool) *MerchRepository {
	return &MerchRepository{db: db}
}
