package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type PgxIface interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag , error)
}

func InitDB() (*pgx.Conn, error) {
	connStr := "user=postgres password=SEg04lsku dbname=inventaris sslmode=disable host=localhost"
	conn, err := pgx.Connect(context.Background(), connStr)
	return conn, err
}
