package cliprinter

import (
	"errors"
	"io"

	commoncli "github.com/spiffe/spire/pkg/common/cli"
	"google.golang.org/protobuf/proto"
)

// Printer is an interface for providing a printer implementation to
// a CLI utility.
type Printer interface {
	PrintError(error) error
	PrintProto(...proto.Message) error
	PrintStruct(...any) error
}

// CustomPrettyFunc is used to provide a custom function for pretty
// printing messages. The intent is to provide a migration pathway
// for pre-existing CLI code, such that this code can supply a
// custom pretty printer that mirrors its current behavior, but
// still be able to gain formatter functionality for other outputs.
type CustomPrettyFunc func(*commoncli.Env, ...any) error

// ErrInternalCustomPrettyFunc should be returned by a CustomPrettyFunc when some internal error occurs.
var ErrInternalCustomPrettyFunc = errors.New("internal error: cli printer; please report this bug")

type printer struct {
	format formatType
	env    *commoncli.Env
	cp     CustomPrettyFunc
}

func newPrinter(f formatType, env *commoncli.Env) *printer { _ = "STUB: not implemented"; return nil }

// PrintError prints an error and applies the configured formatting.
func (p *printer) PrintError(err error) error { _ = "STUB: not implemented"; return nil }

// PrintProto prints a protobuf message and applies the configured formatting.
func (p *printer) PrintProto(msg ...proto.Message) error { _ = "STUB: not implemented"; return nil }

// PrintStruct prints a struct and applies the configured formatting.
func (p *printer) PrintStruct(msg ...any) error { _ = "STUB: not implemented"; return nil }

func (p *printer) printError(err error) error { _ = "STUB: not implemented"; return nil }

func (p *printer) printProto(msg ...proto.Message) error { _ = "STUB: not implemented"; return nil }

func (p *printer) printStruct(msg ...any) error { _ = "STUB: not implemented"; return nil }

func (p *printer) getFormat() formatType { _ = "STUB: not implemented"; return *new(formatType) }

func (p *printer) setCustomPrettyPrinter(cp CustomPrettyFunc) { _ = "STUB: not implemented"; return }

func (p *printer) printPrettyError(err error, stdout, stderr io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *printer) printPrettyProto(msgs []proto.Message, stdout, stderr io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *printer) printPrettyStruct(msg []any, stdout, stderr io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}
