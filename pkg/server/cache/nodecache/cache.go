package nodecache

import (
	"context"
	"sync"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire/pkg/server/datastore"
	"github.com/spiffe/spire/proto/spire/common"
)

const (
	rebuildInterval = 5 * time.Second
)

type Cache struct {
	log              logrus.FieldLogger
	ds               datastore.DataStore
	clk              clock.Clock
	automaticRefresh bool
	enableCache      bool
	buildTime        time.Time
	mtx              sync.RWMutex
	nodes            map[string]*common.AttestedNode
	nodeRefreshTime  map[string]time.Time
}

func New(ctx context.Context, log logrus.FieldLogger, ds datastore.DataStore, clk clock.Clock, automaticRefresh, enableCache bool) (*Cache, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Cache) LookupAttestedNode(id string) (*common.AttestedNode, time.Time) {
	_ = "STUB: not implemented"
	return nil, *new(time.Time)
}

func (c *Cache) FetchAttestedNode(ctx context.Context, id string) (*common.AttestedNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Cache) Rebuild(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *Cache) UpdateAttestedNode(node *common.AttestedNode) { _ = "STUB: not implemented"; return }

func (c *Cache) RemoveAttestedNode(spiffeId string) { _ = "STUB: not implemented"; return }

func (c *Cache) PeriodicRebuild(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
