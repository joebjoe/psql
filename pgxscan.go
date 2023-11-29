package psql

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
)

func selectRows(ctx context.Context, conn pgxscan.Querier, dst any, query string, args ...any) error {
	return pgxscan.Select(ctx, conn, dst, query, args...)
}

func getRow(ctx context.Context, conn pgxscan.Querier, dst any, query string, args ...any) error {
	return pgxscan.Get(ctx, conn, dst, query, args...)
}
