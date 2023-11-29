package psql

import (
	"context"
	"fmt"
	"io"
)

type DB interface {
	Conn
	io.Closer
	Ping(ctx context.Context) error
	Begin(ctx context.Context) (Tx, error)
}

type Conn interface {
	Queryer
	Exec(ctx context.Context, stmt string, args ...any) (CommandTag, error)
}

type Queryer interface {
	Query(ctx context.Context, dst any, query string, args ...any) error
	QueryRow(ctx context.Context, dst any, query string, args ...any) error
}

type Tx interface {
	Conn
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type CommandTag interface {
	fmt.Stringer
	RowsAffected() int64
	Insert() bool
	Update() bool
	Delete() bool
	Select() bool
}
