package main

import (
	"flag"
	"log"
)

var (
	socketPathFlag = flag.String("socket", "unix:///tmp/spire-server/private/api.sock", "server socket path")
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Debug server client fails: %v", err)
	}
}

func run() error { _ = "STUB: not implemented"; return nil }
