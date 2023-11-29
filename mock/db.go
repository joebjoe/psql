package mock

import (
	"context"

	"github.com/joebjoe/psql"
)

type DB struct {
	txs []*Tx

	BeginCalledTimes int
	BeginCalledWith  []context.Context
	MockBegin        func(ctx context.Context) (*Tx, error)

	CloseCalledTimes int
	MockClose        func() error

	ExecCalledTimes       int
	ExecCalledWithContext []context.Context
	ExecCalledWithStmt    []string
	ExecCalledWithArgs    [][]any
	MockExec              func(ctx context.Context, stmt string, args ...any) (CommandTag, error)

	PingCalledTimes int
	PingCalledWith  []context.Context
	MockPing        func(ctx context.Context) error

	QueryCalledTimes       int
	QueryCalledWithContext []context.Context
	QueryCalledWithDst     []any
	QueryCalledWithQuery   []string
	QueryCalledWithArgs    [][]any
	MockQuery              func(ctx context.Context, dst any, query string, args ...any) error

	QueryRowCalledTimes       int
	QueryRowCalledWithContext []context.Context
	QueryRowCalledWithDst     []any
	QueryRowCalledWithQuery   []string
	QueryRowCalledWithArgs    [][]any
	MockQueryRow              func(ctx context.Context, dst any, query string, args ...any) error
}

// TxHistory returns every Tx created by Begin, scrubbed of any non-assertable fields
// (mock functions) so as to be able to perform deep-equal, expected-to-actual assertionss.
func (db *DB) TxHistory() []*Tx {
	txs := make([]*Tx, len(db.txs))

	for i, tx := range db.txs {
		txs[i] = &Tx{
			CommitCalledTimes:         tx.CommitCalledTimes,
			CommitCalledWith:          tx.CommitCalledWith,
			ExecCalledTimes:           tx.ExecCalledTimes,
			ExecCalledWithContext:     tx.ExecCalledWithContext,
			ExecCalledWithStmt:        tx.ExecCalledWithStmt,
			ExecCalledWithArgs:        tx.ExecCalledWithArgs,
			QueryCalledTimes:          tx.QueryCalledTimes,
			QueryCalledWithContext:    tx.QueryCalledWithContext,
			QueryCalledWithDst:        tx.QueryCalledWithDst,
			QueryCalledWithQuery:      tx.QueryCalledWithQuery,
			QueryCalledWithArgs:       tx.QueryCalledWithArgs,
			QueryRowCalledTimes:       tx.QueryRowCalledTimes,
			QueryRowCalledWithContext: tx.QueryRowCalledWithContext,
			QueryRowCalledWithDst:     tx.QueryRowCalledWithDst,
			QueryRowCalledWithQuery:   tx.QueryRowCalledWithQuery,
			QueryRowCalledWithArgs:    tx.QueryRowCalledWithArgs,
			RollbackCalledTimes:       tx.RollbackCalledTimes,
			RollbackCalledWith:        tx.RollbackCalledWith,
		}
	}

	return txs
}

func (db *DB) Begin(ctx context.Context) (psql.Tx, error) {
	db.BeginCalledTimes++
	db.BeginCalledWith = append(db.BeginCalledWith, ctx)

	tx := &Tx{}
	if db.MockBegin != nil {
		var err error
		tx, err = db.MockBegin(ctx)
		if err != nil {
			return nil, err
		}
	}

	db.txs = append(db.txs, tx)

	return tx, nil
}
func (db *DB) Close() error {
	db.CloseCalledTimes++

	if db.MockClose != nil {
		return db.MockClose()
	}

	return nil
}
func (db *DB) Exec(ctx context.Context, stmt string, args ...any) (psql.CommandTag, error) {
	db.ExecCalledTimes++

	db.ExecCalledWithContext = append(db.ExecCalledWithContext, ctx)
	db.ExecCalledWithStmt = append(db.ExecCalledWithStmt, stmt)
	db.ExecCalledWithArgs = append(db.ExecCalledWithArgs, args)

	if db.MockExec != nil {
		return db.MockExec(ctx, stmt, args...)
	}

	return nil, nil
}
func (db *DB) Ping(ctx context.Context) error {
	db.PingCalledTimes++
	db.PingCalledWith = append(db.PingCalledWith, ctx)

	if db.MockPing != nil {
		return db.MockPing(ctx)
	}

	return nil
}
func (db *DB) Query(ctx context.Context, dst any, query string, args ...any) error {
	db.QueryCalledTimes++

	db.QueryCalledWithContext = append(db.QueryCalledWithContext, ctx)
	db.QueryCalledWithDst = append(db.QueryCalledWithDst, dst)
	db.QueryCalledWithQuery = append(db.QueryCalledWithQuery, query)
	db.QueryCalledWithArgs = append(db.QueryCalledWithArgs, args)

	if db.MockQuery != nil {
		return db.MockQuery(ctx, dst, query, args...)
	}

	return nil
}
func (db *DB) QueryRow(ctx context.Context, dst any, query string, args ...any) error {
	db.QueryRowCalledTimes++

	db.QueryRowCalledWithContext = append(db.QueryRowCalledWithContext, ctx)
	db.QueryRowCalledWithDst = append(db.QueryRowCalledWithDst, dst)
	db.QueryRowCalledWithQuery = append(db.QueryRowCalledWithQuery, query)
	db.QueryRowCalledWithArgs = append(db.QueryRowCalledWithArgs, args)

	if db.MockQueryRow != nil {
		return db.MockQueryRow(ctx, dst, query, args...)
	}

	return nil
}
