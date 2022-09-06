package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

type PSQL struct {
	Pool *pgxpool.Pool
}

func PoolCreate() (*PSQL, error) {

	pool, poolErr := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if poolErr != nil {
		return nil, poolErr
	}

	pingErr := pool.Ping(context.Background())
	if pingErr != nil {
		return nil, pingErr
	}

	db := &PSQL{
		Pool: pool,
	}

	return db, nil
}
