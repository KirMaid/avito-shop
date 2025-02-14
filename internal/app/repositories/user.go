package repositories

import (
	"avitoshop/internal/app/entities"
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrUserDoesNotExist = errors.New("user does not exist")

type UserRepository interface {
	Insert(ctx context.Context, user *entities.User) error
	GetByUsername(ctx context.Context, username string) (*entities.User, error)
	GetByID(ctx context.Context, userID int) (*entities.User, error)
	//TODO Определится всё таки что прокидывать
	UpdateBalance(ctx context.Context, userID int, balance int) error
	//UpdateToken(ctx context.Context, user *entities.User) error
}

type userRepository struct {
	db *pgxpool.Pool
}

// TODO Реализовать
func (r *userRepository) UpdateBalance(ctx context.Context, userID int, balance int) error {
	//TODO implement me
	panic("implement me")
}

//func (r *userRepository) UpdateToken(ctx context.Context, user *entities.User) error {
//	_, err := r.db.Exec(ctx, `
//		UPDATE users
//		SET token = $1, token_expires_at = $2
//		WHERE id = $3
//	`, user.Token, user.TokenExpiresAt, user.ID)
//	return err
//}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

//func (r *userRepository) Insert(ctx context.Context, user *entities.User) error {
//	query := "INSERT INTO users (username, password, token, token_expires_at) VALUES ($1, $2, $3, $4)"
//	_, err := r.db.Exec(ctx, query, user.Username, user.Password, user.Token, user.TokenExpiresAt)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

func (r *userRepository) Insert(ctx context.Context, user *entities.User) error {
	query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err := r.db.Exec(ctx, query, user.Username, user.Password)
	if err != nil {
		return err
	}

	return nil
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
