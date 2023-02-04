package db

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// DB is a wrapper for an sqlx DB that adds sentry tracing to each of the query methods in the sqlx
// library.
type DB struct {
	db *sqlx.DB
}

func NewDb(db *sqlx.DB) *DB {
	return &DB{
		db: db,
	}
}

func Open(driverName string, dataSourceName string) (*DB, error) {
	db, err := sqlx.Open(driverName, dataSourceName)
	return NewDb(db), err
}

func Connect(driverName string, dataSourceName string) (*DB, error) {
	db, err := sqlx.Connect(driverName, dataSourceName)
	return NewDb(db), err
}

func (d *DB) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	return d.db.QueryxContext(ctx, query, args...)
}

func (d *DB) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	return d.db.QueryRowxContext(ctx, query, args...)
}

func (d *DB) NamedQueryContext(ctx context.Context, query string, args interface{}) (*sqlx.Rows, error) {
	return d.db.NamedQueryContext(ctx, query, args)
}

func (d *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return d.db.ExecContext(ctx, query, args...)
}

func (d *DB) NamedExecContext(ctx context.Context, query string, args interface{}) (sql.Result, error) {
	return d.db.NamedExecContext(ctx, query, args)
}

func (d *DB) Rebind(query string) string {
	return d.db.Rebind(query)
}

func (d *DB) Stats() sql.DBStats {
	return d.db.Stats()
}
