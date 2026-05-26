package cli

import (
	"io"
	"os"
)

var (
	// DefaultEnv is the default environment used by commands
	DefaultEnv = &Env{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
)

// Env provides a pluggable environment for CLI commands that facilitates easy
// testing.
type Env struct {
	Stdin   io.Reader
	Stdout  io.Writer
	Stderr  io.Writer
	BaseDir string
}

func (e *Env) Printf(format string, args ...any) error { _ = "STUB: not implemented"; return nil }

func (e *Env) Println(args ...any) error { _ = "STUB: not implemented"; return nil }

func (e *Env) ErrPrintf(format string, args ...any) error { _ = "STUB: not implemented"; return nil }

func (e *Env) ErrPrintln(args ...any) error { _ = "STUB: not implemented"; return nil }

func (e *Env) JoinPath(parts ...string) string { _ = "STUB: not implemented"; return "" }
