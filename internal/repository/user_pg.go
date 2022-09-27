package repository

import (
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

	fmt.Println("test sql method")

	//row := userPg.pool.QueryRow(context.Background(),
	//	"INSERT INTO phonebook (name, phone) VALUES ($1, $2) RETURNING id",
	//	rec.Name, rec.Phone)

	return 1, nil
}
