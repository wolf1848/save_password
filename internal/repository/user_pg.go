package repository

import "github.com/jackc/pgx/v4/pgxpool"

type userPg struct {
	pool *pgxpool.Pool
}

func createUserPg(pool *pgxpool.Pool) *userPg {
	return &userPg{
		pool: pool,
	}
}

func (user *userPg) create() (int, error) {

	//row := userPg.pool.QueryRow(context.Background(),
	//	"INSERT INTO phonebook (name, phone) VALUES ($1, $2) RETURNING id",
	//	rec.Name, rec.Phone)

	return 1, nil
}
