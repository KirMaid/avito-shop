package repository

import (
	"avitoshop/internal/app/auth"
	"avitoshop/internal/app/entity"
	"context"
	//"log"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

// TODO Разобраться с контекстом, зачем он вообще нужен
func (r *UserRepository) Insert(ctx context.Context, user *entity.User) error {
	query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err := r.db.Exec(ctx, query, user.Username, user.Password)
	if err != nil {
		//log.Errorf("error on inserting user: %s", err.Error())
		return auth.ErrUserAlreadyExists
	}

	return nil
}

func (r *UserRepository) Get(ctx context.Context, username string) (*entity.User, error) {
	//TODO Посмотреть механику работы с new
	user := new(entity.User)
	query := "SELECT username, password FROM users WHERE username = $1"
	err := r.db.QueryRow(ctx, query, username).Scan(&user.Username, &user.Password)
	if err != nil {
		//TODO Разобраться с логгированием
		//if err.Error() == sql.ErrNoRows {
		//	log.Errorf("user does not exist: %s", err.Error())
		//	return nil, auth.ErrUserDoesNotExist
		//}
		//log.Errorf("error occurred while getting user from db: %s", err.Error())
		return nil, err
	}
	return nil, nil
}
