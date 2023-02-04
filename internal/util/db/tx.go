package db

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Tx struct {
	tx *sqlx.Tx
}

func (t *Tx) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	return t.tx.QueryxContext(ctx, query, args...)
}

func (t *Tx) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	return t.tx.QueryRowxContext(ctx, query, args...)
}

func (t *Tx) NamedQueryContext(ctx context.Context, query string, args interface{}) (*sqlx.Rows, error) {
	return sqlx.NamedQueryContext(ctx, t.tx, query, args)
}

func (t *Tx) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return t.tx.ExecContext(ctx, query, args...)
}

func (t *Tx) NamedExecContext(ctx context.Context, query string, args interface{}) (sql.Result, error) {
	return t.tx.NamedExecContext(ctx, query, args)
}

// func (t *Tx) PreparexContext(ctx context.Context, query string) (*Stmt, error) {
// 	stmt, err := t.tx.PreparexContext(ctx, query)
// 	return stmtWithTracing(stmt, query, t.parentSpan), err
// }

// func (t *Tx) PrepareNamedContext(ctx context.Context, query string) (*NamedStmt, error) {
// 	stmt, err := t.tx.PrepareNamedContext(ctx, query)
// 	return namedStmtWithTracing(stmt, query), err
// }

// func (t *Tx) Rollback() error {
// 	defer func() {
// 		t.parentSpan.Finish()
// 	}()
// 	return t.tx.Rollback()
// }

func (t *Tx) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return t.tx.GetContext(ctx, dest, query, args...)
}

func (t *Tx) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return t.tx.SelectContext(ctx, dest, query, args...)
}

// func (t *Tx) Commit() error {
// 	defer func() {
// 		t.parentSpan.Finish()
// 	}()
// 	return t.tx.Commit()
// }

func (t *Tx) Rebind(query string) string {
	return t.tx.Rebind(query)
}
