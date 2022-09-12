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
	return 1, nil
}
