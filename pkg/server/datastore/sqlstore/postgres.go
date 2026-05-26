package sqlstore

import (
	"context"

	"github.com/jinzhu/gorm"

	// gorm postgres `cloudsql` dialect, for GCP Cloud SQL Proxy
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	// gorm postgres dialect init registration
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type postgresDB struct{}

func (p postgresDB) connect(ctx context.Context, cfg *configuration, isReadOnly bool) (db *gorm.DB, version string, supportsCTE bool, err error) {
	_ = "STUB: not implemented"
	return nil, "", false, nil
}

// Supported versions of PostgreSQL all support CTE so unconditionally
// return true.

func (p postgresDB) isConstraintViolation(err error) bool { _ = "STUB: not implemented"; return false }

// "23xxx" is the constraint violation class for PostgreSQL
