//go:build windows

package peertracker

import (
	"io"
	"net"
	"os/exec"
	"testing"

	"github.com/sirupsen/logrus"
)

const (
	childSource = "peertracker_test_child_windows.go"
)

type fakePeer struct {
	grandchildPID int
	conn          net.Conn
	childStdin    io.Closer
	childCmd      *exec.Cmd
	t             *testing.T
}

func (f *fakePeer) killGrandchild() { _ = "STUB: not implemented"; return }

// Wait for the process to exit, so we are sure that we can
// clean up the directory containing the executable

func addr(*testing.T) net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func listener(t *testing.T, log *logrus.Logger, addr net.Addr) *Listener {
	_ = "STUB: not implemented"
	return nil
}

func childExecCommand(childPath string, addr net.Addr) *exec.Cmd {
	_ = "STUB: not implemented"
	// #nosec G204 test code
	return nil
}

func dial(addr net.Addr) (net.Conn, error) { _ = "STUB: not implemented"; return *new(net.Conn), nil }
