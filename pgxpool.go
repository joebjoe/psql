package psql

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type pgxPool struct {
	conn *pgxpool.Pool
}

type Config = pgxpool.Config

func New(ctx context.Context, cxn string) (DB, error) {
	conn, err := pgxpool.New(ctx, cxn)
	if err != nil {
		return nil, err
	}
	return &pgxPool{conn: conn}, nil
}

func NewWithConfig(ctx context.Context, cfg *Config) (DB, error) {
	conn, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return &pgxPool{conn: conn}, nil
}

func (p *pgxPool) Begin(ctx context.Context) (Tx, error) {
	tx, err := p.conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return &pgxTx{conn: tx}, nil
}

func (p *pgxPool) Close() error {
	p.conn.Close()
	return nil
}

func (p *pgxPool) Exec(ctx context.Context, stmt string, args ...any) (CommandTag, error) {
	return p.conn.Exec(ctx, stmt, args...)
}

func (p *pgxPool) Ping(ctx context.Context) error {
	return p.conn.Ping(ctx)
}

func (p *pgxPool) Query(ctx context.Context, dst any, query string, args ...any) error {
	return selectRows(ctx, p.conn, dst, query, args...)
}

func (p *pgxPool) QueryRow(ctx context.Context, dst any, query string, args ...any) error {
	return getRow(ctx, p.conn, dst, query, args...)
}
