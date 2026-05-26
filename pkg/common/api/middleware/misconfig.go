package middleware

import (
	"context"
	"sync"
	"time"

	"github.com/andres-erbsen/clock"
)

var (
	misconfigLogMtx   sync.Mutex
	misconfigLogTimes = make(map[string]time.Time)
	misconfigClk      = clock.New()
)

const misconfigLogEvery = time.Minute

// LogMisconfiguration logs a misconfiguration for the RPC. It assumes that the
// context has been embellished with the names for the RPC. This method should
// not be called under normal operation and only when there is an
// implementation bug. As such there is no attempt at a time/space efficient
// implementation. In any case, the number of distinct misconfiguration
// messages intersected with the number of RPCs should not produce any amount
// of real memory use. Contention on the global mutex should also be
// reasonable.
func LogMisconfiguration(ctx context.Context, msg string) { _ = "STUB: not implemented"; return }

func shouldLogMisconfiguration(ctx context.Context, msg string) bool {
	_ = "STUB: not implemented"
	return false
}
