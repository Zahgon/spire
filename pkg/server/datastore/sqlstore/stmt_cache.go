package sqlstore

import (
	"context"
	"database/sql"
	"sync"
)

type stmtCache struct {
	db    *sql.DB
	stmts sync.Map
}

func newStmtCache(db *sql.DB) *stmtCache { _ = "STUB: not implemented"; return nil }

func (cache *stmtCache) get(ctx context.Context, query string) (*sql.Stmt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Somebody beat us to it. Close the statement we prepared.
