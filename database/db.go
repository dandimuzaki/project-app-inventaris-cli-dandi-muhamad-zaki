package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type PgxIface interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}

func InitDB() (*pgx.Conn, error) {
	connStr := "user=postgres password=SEg04lsku dbname=Ojek-Online sslmode=disable host=localhost"
	conn, err := pgx.Connect(context.Background(), connStr)
	return conn, err
}
