//go:build !cgo

package sqlstore

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type sqliteDB struct {
	log logrus.FieldLogger
}

func (s sqliteDB) connect(ctx context.Context, cfg *configuration, isReadOnly bool) (db *gorm.DB, version string, supportsCTE bool, err error) {
	_ = "STUB: not implemented"
	return nil, "", false, nil
}

func (s sqliteDB) isConstraintViolation(err error) bool { _ = "STUB: not implemented"; return false }
