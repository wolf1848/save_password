package repository

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	UserPg
}

func CreateRepository() (*Repository, error) {

	pool, poolErr := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if poolErr != nil {
		return nil, poolErr
	}

	pingErr := pool.Ping(context.Background())
	if pingErr != nil {
		return nil, pingErr
	}

	db := &Repository{
		UserPg: createUserPg(pool),
	}

	return db, nil
}
