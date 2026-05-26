package sqlstore

import (
	"context"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	// gorm mysql `cloudsql` dialect, for GCP
	// Cloud SQL Proxy
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	// gorm mysql dialect init registration
	// also needed for GCP Cloud SQL Proxy
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type mysqlDB struct {
	logger logrus.FieldLogger
}

const (
	tlsConfigName = "spireCustomTLS"
)

func (my mysqlDB) connect(ctx context.Context, cfg *configuration, isReadOnly bool) (db *gorm.DB, version string, supportsCTE bool, err error) {
	_ = "STUB: not implemented"
	return nil, "", false, nil
}

func (my mysqlDB) supportsCTE(ctx context.Context, gormDB *gorm.DB) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (my mysqlDB) isParseError(err error) bool { _ = "STUB: not implemented"; return false }

// ER_PARSE_ERROR

func (my mysqlDB) isConstraintViolation(err error) bool { _ = "STUB: not implemented"; return false }

// ER_DUP_ENTRY

// configureConnection modifies the connection string to support features that
// normally require code changes, like custom Root CAs or client certificates
func configureConnection(cfg *configuration, isReadOnly bool) (*mysql.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// the connection string should have already been validated by now
// (in validateMySQLConfig)

// connection string doesn't have to be modified

// load and configure Root CA if it exists

// load and configure client certificate if it exists

// register a custom TLS config that uses custom Root CAs with the MySQL driver

// instruct MySQL driver to use the custom TLS config

func hasTLSConfig(cfg *configuration) bool { _ = "STUB: not implemented"; return false }

func validateMySQLConfig(cfg *configuration, isReadOnly bool) error {
	_ = "STUB: not implemented"
	return nil
}
