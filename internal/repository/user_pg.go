package repository

import (
	"context"
	"fmt"
	"startGo/internal/data"

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
		"INSERT INTO users (login, full_name, hash_pwd) VALUES (coalesce(TRIM($1),'') != '', $2, $3) RETURNING id",
		data.Login, data.FullName, data.Pwd)

	if err := row.Scan(&id); err != nil {

		fmt.Println(err)

		return 0, err
	}

	return id, nil
}
