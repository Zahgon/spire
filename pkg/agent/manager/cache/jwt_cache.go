package cache

import (
	"container/list"
	"context"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/agent/client"
	"github.com/spiffe/spire/pkg/common/telemetry"
)

type JWTSVIDCache struct {
	log     logrus.FieldLogger
	metrics telemetry.Metrics
	mu      sync.RWMutex

	svids   map[string]*list.Element
	lruList *list.List

	// svidCacheMaxSize is a hard limit of max number of SVIDs that would be stored in cache
	svidCacheMaxSize int
}

type jwtSvidElement struct {
	key  string
	svid *client.JWTSVID
}

func (c *JWTSVIDCache) CountJWTSVIDs() int { _ = "STUB: not implemented"; return 0 }

func NewJWTSVIDCache(log logrus.FieldLogger, metrics telemetry.Metrics, svidCacheMaxSize int) *JWTSVIDCache {
	_ = "STUB: not implemented"
	return nil
}

func (c *JWTSVIDCache) GetJWTSVID(spiffeID spiffeid.ID, audience []string) (*client.JWTSVID, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *JWTSVIDCache) SetJWTSVID(spiffeID spiffeid.ID, audience []string, svid *client.JWTSVID) {
	_ = "STUB: not implemented"
	return
}

func (c *JWTSVIDCache) TaintJWTSVIDs(ctx context.Context, taintedJWTAuthorities map[string]struct{}) {
	_ = "STUB: not implemented"
	return
}

func getKeyIDFromSVIDToken(svidToken string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func jwtSVIDKey(spiffeID spiffeid.ID, audience []string) string {
	_ = "STUB: not implemented"

	// Form the cache key as the SHA-256 hash of the SPIFFE ID and all the audiences.
	// In order to avoid ambiguities, we will write a nul byte to the hash function after each data
	// item.
	return ""
}

// duplicate and sort the audience slice
