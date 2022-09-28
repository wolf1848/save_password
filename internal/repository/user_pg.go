package repository

import (
	"context"
	"errors"
	"startGo/internal/data"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userPg struct {
	pool *pgxpool.Pool
}

type UserPg interface {
	Create(data *data.User) (int, error)
}

func createUserPg(pool *pgxpool.Pool) *userPg {
	return &userPg{
		pool: pool,
	}
}

func (user *userPg) Create(data *data.User) (int, error) {

	var id int

	row := user.pool.QueryRow(context.Background(),
		"INSERT INTO users (login, full_name, hash_pwd) VALUES (trim($1), $2, trim($3)) RETURNING id",
		data.Login, data.FullName, data.Pwd)

	if err := row.Scan(&id); err != nil {

		if pgErr, ok := err.(*pgconn.PgError); ok {
			var userErr error

			switch pgErr.ConstraintName {
			case "users_login_check":
				userErr = errors.New("поле \"Логин\" обязательно для заполнения")
			case "users_hash_pwd_check":
				userErr = errors.New("поле \"Пароль\" обязательно для заполнения")
			case "users_login_key":
				userErr = errors.New("поле \"Логин\" не уникально")
			default:
				userErr = errors.New("возникла непредвиденная проблема")

			}

			return 0, userErr
		}

	}

	return id, nil
}
