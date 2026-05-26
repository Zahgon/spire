package awsrds

import (
	"database/sql/driver"
	"sync"
	"time"
)

const (
	MySQLDriverName     = "aws-rds-mysql"
	PostgresDriverName  = "aws-rds-postgres"
	getAuthTokenTimeout = time.Second * 30
)

// nowFunc returns the current time and can overridden in tests.
var nowFunc = time.Now

// Config holds the configuration settings to be able to authenticate to a
// database in the AWS RDS service.
type Config struct {
	Region          string `json:"region"`
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	Endpoint        string `json:"endpoint"`
	DbUser          string `json:"dbuser"`
	DriverName      string `json:"driver_name"`
	ConnString      string `json:"conn_string"`
}

func init() {
	registerPostgres()
	registerMySQL()
}

// FormatDSN returns a DSN string based on the configuration.
func (c *Config) FormatDSN() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (c *Config) getConnStringWithPassword(password string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type tokens map[string]*authToken

// sqlDriverWrapper is a wrapper for SQL drivers, adding IAM authentication.
type sqlDriverWrapper struct {
	sqlDriver    driver.Driver
	tokenBuilder authTokenBuilder

	tokensMapMtx sync.Mutex
	tokensMap    tokens
}

// Open is the overridden method for opening a connection, using
// AWS IAM authentication
func (w *sqlDriverWrapper) Open(name string) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

// We need a context for getting the authentication token. Since there is no
// parent context to derive from, we create a context with a timeout to
// get the authentication token.

func addPasswordToPostgresConnString(connString, password string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func addPasswordToMySQLConnString(connString, password string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// escapeSpecialCharsPostgres escapes special characters within a value of a
// keyword/value postgres connection string.
// Single quotes and backslashes within a value must be escaped with a
// backslash, i.e., \' and \\.
func escapeSpecialCharsPostgres(s string) string { _ = "STUB: not implemented"; return "" }

func registerPostgres() { _ = "STUB: not implemented"; return }

func registerMySQL() { _ = "STUB: not implemented"; return }
