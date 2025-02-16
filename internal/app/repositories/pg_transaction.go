package repositories

import (
	"avitoshop/internal/app/entities"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type transactionRepository struct {
	db *pgxpool.Pool
}

func NewTransactionRepository(db *pgxpool.Pool) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Insert(ctx context.Context, transaction *entities.Transaction) (*entities.Transaction, error) {
	query := `
       INSERT INTO transactions (amount, sender_id, receiver_id) 
       VALUES ($1, $2, $3) 
       RETURNING id
    `

	err := r.db.QueryRow(ctx, query, transaction.Amount, transaction.SenderID, transaction.ReceiverID).Scan(
		&transaction.ID,
	)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *transactionRepository) GetReceivedTransactions(ctx context.Context, userID int) ([]entities.Transaction, error) {
	query := `
		SELECT id, amount, sender_id, receiver_id, created_at
		FROM transactions
		WHERE receiver_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query received transactions: %w", err)
	}
	defer rows.Close()

	var transactions []entities.Transaction
	for rows.Next() {
		var tx entities.Transaction
		if err := rows.Scan(&tx.ID, &tx.Amount, &tx.SenderID, &tx.ReceiverID, &tx.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan transaction: %w", err)
		}
		transactions = append(transactions, tx)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating rows: %w", err)
	}

	return transactions, nil
}

func (r *transactionRepository) GetSentTransactions(ctx context.Context, userID int) ([]entities.Transaction, error) {
	query := `
		SELECT id, amount, sender_id, receiver_id, created_at
		FROM transactions
		WHERE sender_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query sent transactions: %w", err)
	}
	defer rows.Close()

	var transactions []entities.Transaction
	for rows.Next() {
		var tx entities.Transaction
		if err := rows.Scan(&tx.ID, &tx.Amount, &tx.SenderID, &tx.ReceiverID, &tx.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan transaction: %w", err)
		}
		transactions = append(transactions, tx)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating rows: %w", err)
	}

	return transactions, nil
}
