package cliprinter

import (
	"flag"
	"fmt"

	commoncli "github.com/spiffe/spire/pkg/common/cli"
)

const defaultFlagName = "output"

var flagDescription = fmt.Sprintf(
	"Desired output format (%s, %s); default: %s.",
	formatTypeToStr(pretty),
	formatTypeToStr(json),
	formatTypeToStr(defaultFormatType),
)

// AppendFlag adds the -format flag to the provided flagset, and populates
// the referenced Printer interface with a properly configured printer.
func AppendFlag(p *Printer, fs *flag.FlagSet, env *commoncli.Env) *FormatterFlag {
	_ = "STUB: not implemented"
	return nil
}

// AppendFlagWithCustomPretty is the same as AppendFlag, however it also allows
// a custom pretty function to be specified. A custom pretty function can be used
// to override the pretty print logic that normally ships with this package. Its
// intended use is to allow for the adoption of cliprinter while still retaining
// backwards compatibility with the legacy/bespoke pretty print output.
func AppendFlagWithCustomPretty(p *Printer, fs *flag.FlagSet, env *commoncli.Env, cp CustomPrettyFunc) *FormatterFlag {
	_ = "STUB: not implemented"
	// Set the default
	return nil
}

type FormatterFlag struct {
	customPretty CustomPrettyFunc

	// A pointer to our consumer's Printer interface, along with
	// its format type
	p     *Printer
	f     formatType
	env   *commoncli.Env
	isSet bool
}

func (f *FormatterFlag) String() string { _ = "STUB: not implemented"; return "" }

func (f *FormatterFlag) Set(formatStr string) error { _ = "STUB: not implemented"; return nil }
