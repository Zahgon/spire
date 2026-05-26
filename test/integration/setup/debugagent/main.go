package main

import (
	"context"
	"flag"
	"log"
)

var (
	socketPathFlag = flag.String("debugSocketPath", "unix:///opt/debug.sock", "agent socket path")

	testCaseFlag = flag.String("testCase", "agentEndpoints", "running test case")
)

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatalf("Debug client failed: %v", err)
	}
}

func run() error { _ = "STUB: not implemented"; return nil }

func agentEndpoints(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// printDebugPage allows integration tests to easily parse debug page with jq
func printDebugPage(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func retrieveDebugPage(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func serverWithWorkload(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func serverWithInsecure(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func validateError(err error) error { _ = "STUB: not implemented"; return nil }
