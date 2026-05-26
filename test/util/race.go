package util

import (
	"testing"
)

var (
	raceTestNumThreads = 2
	raceTestNumLoops   = 2
)

func init() {
	raceTestNumThreads = getEnvInt("SPIRE_TEST_RACE_NUM_THREADS", raceTestNumThreads)
	raceTestNumLoops = getEnvInt("SPIRE_TEST_RACE_NUM_LOOPS", raceTestNumLoops)
}

func RaceTest(t *testing.T, fn func(*testing.T)) {
	_ = "STUB: not implemented"
	// wrap in a top level group to ensure all subtests
	// complete before this method returns. All subtests
	// will be run in parallel
	return
}

func getEnvInt(name string, fallback int) int { _ = "STUB: not implemented"; return 0 }
