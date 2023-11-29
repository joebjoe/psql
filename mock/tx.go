package mock

import (
	"context"

	"github.com/joebjoe/psql"
)

type Tx struct {
	CommitCalledTimes int
	CommitCalledWith  []context.Context
	MockCommit        func(ctx context.Context) error

	ExecCalledTimes       int
	ExecCalledWithContext []context.Context
	ExecCalledWithStmt    []string
	ExecCalledWithArgs    [][]any
	MockExec              func(ctx context.Context, stmt string, args ...any) (CommandTag, error)

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

	RollbackCalledTimes int
	RollbackCalledWith  []context.Context
	MockRollback        func(ctx context.Context) error
}

func (tx *Tx) Commit(ctx context.Context) error {
	tx.CommitCalledTimes++

	tx.CommitCalledWith = append(tx.CommitCalledWith, ctx)

	if tx.MockCommit != nil {
		return tx.MockCommit(ctx)
	}

	return nil
}

func (tx *Tx) Exec(ctx context.Context, stmt string, args ...any) (psql.CommandTag, error) {
	tx.ExecCalledTimes++

	tx.ExecCalledWithContext = append(tx.ExecCalledWithContext, ctx)
	tx.ExecCalledWithStmt = append(tx.ExecCalledWithStmt, stmt)
	tx.ExecCalledWithArgs = append(tx.ExecCalledWithArgs, args)

	if tx.MockExec != nil {
		return tx.MockExec(ctx, stmt, args...)
	}

	return nil, nil
}

func (tx *Tx) Query(ctx context.Context, dst any, query string, args ...any) error {
	tx.QueryCalledTimes++

	tx.QueryCalledWithContext = append(tx.QueryCalledWithContext, ctx)
	tx.QueryCalledWithDst = append(tx.QueryCalledWithDst, dst)
	tx.QueryCalledWithQuery = append(tx.QueryCalledWithQuery, query)
	tx.QueryCalledWithArgs = append(tx.QueryCalledWithArgs, args)

	if tx.MockQuery != nil {
		return tx.MockQuery(ctx, dst, query, args...)
	}

	return nil
}

func (tx *Tx) QueryRow(ctx context.Context, dst any, query string, args ...any) error {
	tx.QueryRowCalledTimes++

	tx.QueryRowCalledWithContext = append(tx.QueryRowCalledWithContext, ctx)
	tx.QueryRowCalledWithDst = append(tx.QueryRowCalledWithDst, dst)
	tx.QueryRowCalledWithQuery = append(tx.QueryRowCalledWithQuery, query)
	tx.QueryRowCalledWithArgs = append(tx.QueryRowCalledWithArgs, args)

	if tx.MockQueryRow != nil {
		return tx.MockQueryRow(ctx, dst, query, args...)
	}

	return nil
}

func (tx *Tx) Rollback(ctx context.Context) error {
	tx.RollbackCalledTimes++

	tx.RollbackCalledWith = append(tx.RollbackCalledWith, ctx)

	if tx.MockRollback != nil {
		return tx.MockRollback(ctx)
	}

	return nil
}
