package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func GetConnPool(ctx context.Context) *pgx.Conn {
	connConfig, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	connPool, err := pgx.ConnectConfig(ctx, connConfig)
	if err != nil {
		log.Fatal(err)
	}

	return connPool
}
