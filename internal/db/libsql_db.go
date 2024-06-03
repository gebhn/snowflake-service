package db

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type libsqlDb struct {
	conn *sql.DB
}

func NewLibsqlConn(connectionString, authToken string) *libsqlDb {
	conn, err := sql.Open("libsql", connectionString+"?authToken="+authToken)
	if err != nil {
		log.Fatal(err)
	}
	return &libsqlDb{
		conn: conn,
	}
}

func (c *libsqlDb) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return c.conn.ExecContext(ctx, query, args...)
}

func (c *libsqlDb) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return c.conn.PrepareContext(ctx, query)
}

func (c *libsqlDb) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return c.conn.QueryContext(ctx, query, args...)
}

func (c *libsqlDb) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return c.conn.QueryRowContext(ctx, query, args...)
}
