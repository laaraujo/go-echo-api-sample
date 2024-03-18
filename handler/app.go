package handler

import (
	db "golang-echo-api/db/sqlc"

	"github.com/jackc/pgx/v5"
)

type App struct {
	ConnPool *pgx.Conn
	Queries  *db.Queries
}

func NewApp(connPool *pgx.Conn) *App {
	return &App{
		ConnPool: connPool,
		Queries:  db.New(connPool),
	}
}
