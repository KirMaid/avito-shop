package repositories

import (
	"avitoshop/internal/app/entities"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrUserDoesNotExist = errors.New("user does not exist")

type userRepository struct {
	db *pgxpool.Pool
}

func (r *userRepository) UpdateBalance(ctx context.Context, userID int, balance int) error {
	commandTag, err := r.db.Exec(ctx, "UPDATE users SET balance = $1 WHERE id = $2", balance, userID)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return ErrUserDoesNotExist
	}

	return nil
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Insert(ctx context.Context, user *entities.User) (*entities.User, error) {
	query := `
        INSERT INTO users (username, password) 
        VALUES ($1, $2)
        RETURNING id,balance
    `
	err := r.db.QueryRow(ctx, query, user.Username, user.Password).Scan(
		&user.ID,
		&user.Balance,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*entities.User, error) {
	var user entities.User

	err := r.db.QueryRow(ctx, "SELECT id, username, password, balance FROM users WHERE username = $1", username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Balance,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserDoesNotExist
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetByID(ctx context.Context, userID int) (*entities.User, error) {
	var user entities.User
	err := r.db.QueryRow(ctx, "SELECT * FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Username, &user.Password, &user.Balance)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByIDs(ctx context.Context, userIDs []int) ([]entities.User, error) {
	rows, err := r.db.Query(ctx, "SELECT id, username, password, balance FROM users WHERE id = ANY($1)", userIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to query usernames: %w", err)
	}
	defer rows.Close()

	var users []entities.User

	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Balance); err != nil {
			return nil, fmt.Errorf("failed to scan username: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating rows: %w", err)
	}

	return users, nil
}

func (r *userRepository) GetUsernamesByIDs(ctx context.Context, userIDs []int) (map[int]string, error) {
	rows, err := r.db.Query(ctx, "SELECT id, username FROM users WHERE id = ANY($1)", userIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to query usernames: %w", err)
	}
	defer rows.Close()

	usernames := make(map[int]string)

	for rows.Next() {
		var userID int
		var username string

		if err := rows.Scan(&userID, &username); err != nil {
			return nil, fmt.Errorf("failed to scan username: %w", err)
		}

		usernames[userID] = username
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating rows: %w", err)
	}

	return usernames, nil
}
