package psql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type pgxTx struct {
	conn pgx.Tx
}

func (tx *pgxTx) Commit(ctx context.Context) error {
	return tx.conn.Commit(ctx)
}

func (tx *pgxTx) Exec(ctx context.Context, stmt string, args ...any) (CommandTag, error) {
	return tx.conn.Exec(ctx, stmt, args...)
}

func (tx *pgxTx) Query(ctx context.Context, dst any, query string, args ...any) error {
	return selectRows(ctx, tx.conn, dst, query, args...)
}

func (tx *pgxTx) QueryRow(ctx context.Context, dst any, query string, args ...any) error {
	return getRow(ctx, tx.conn, dst, query, args...)
}

func (tx *pgxTx) Rollback(ctx context.Context) error {
	return tx.conn.Rollback(ctx)
}
