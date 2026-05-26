//go:build cgo

package sqlstore

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	// gorm sqlite dialect init registration
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type sqliteDB struct {
	log logrus.FieldLogger
}

func (s sqliteDB) connect(ctx context.Context, cfg *configuration, isReadOnly bool) (db *gorm.DB, version string, supportsCTE bool, err error) {
	_ = "STUB: not implemented"
	return nil, "", false, nil
}

// The embedded version of SQLite3 unconditionally supports CTE.

func (s sqliteDB) isConstraintViolation(err error) bool { _ = "STUB: not implemented"; return false }

func openSQLite3(connString string) (*gorm.DB, error) { _ = "STUB: not implemented"; return nil, nil }

// embellishSQLite3ConnString adds query values supported by
// github.com/mattn/go-sqlite3 to enable journal mode and foreign key support.
// These query values MUST be part of the connection string in order to be
// enabled for *each* connection opened by db/sql. If the connection string is
// not already a file: URI, it is converted first.
func embellishSQLite3ConnString(connectionString string) (string, error) {
	_ = "STUB: not implemented"
	// On Windows, when parsing an absolute path like "c:\tmp\lite",
	// "c" is parsed as the URL scheme
	return "", nil
}

// connection string is a path. move the path section into the
// opaque section so it renders property for sqlite3, for example:
// data.db = file:data.db
// ./data.db = file:./data.db
// /data.db = file:/data.db

// only no scheme (i.e. file path) or file scheme is supported
