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
	Create(data *data.User) (uint32, error)
	CheckUser(data *data.User) (uint32, string, string, error)
}

func createUserPg(pool *pgxpool.Pool) *userPg {
	return &userPg{
		pool: pool,
	}
}

func (user *userPg) Create(data *data.User) (uint32, error) {

	var id uint32

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

func (user *userPg) CheckUser(data *data.User) (uint32, string, string, error) {

	var uid uint32
	var hash_pwd string
	var full_name string
	row := user.pool.QueryRow(context.Background(),
		"SELECT id, hash_pwd, full_name FROM users WHERE login=trim($1) LIMIT 1",
		data.Login)

	if err := row.Scan(&uid, &hash_pwd, &full_name); err != nil {
		return 0, "", "", err
	}

	return uid, hash_pwd, full_name, nil
}
