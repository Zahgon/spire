package catalog

import (
	"io"

	"google.golang.org/grpc"
)

type closerGroup []io.Closer

func (cs closerGroup) Close() error {
	_ = "STUB: not implemented"
	// Close in reverse order.
	return nil
}

type closerFunc func()

func closerFuncs(fns ...func()) closerGroup { _ = "STUB: not implemented"; return *new(closerGroup) }

func (fn closerFunc) Close() error { _ = "STUB: not implemented"; return nil }

func gracefulStopWithTimeout(s *grpc.Server) bool { _ = "STUB: not implemented"; return false }
