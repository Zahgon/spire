package structpretty

import (
	"io"
	"reflect"
)

// Print prints a struct prettily.
// It will print only easily printable types, and only to one
// level of depth. It will print arrays, slices, and maps if
// their keys and elements are also easily printable types.
func Print(msgs []any, stdout, stderr io.Writer) error { _ = "STUB: not implemented"; return nil }

func printStruct(msg any, stdout, stderr io.Writer) error { _ = "STUB: not implemented"; return nil }

// We also want to accept pointers to structs

func isFieldTypePrintable(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

func isArrayPrintable(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

func isMapPrintable(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

func isCompositeTypePrintable(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

func isUnprintableType(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

func isListType(t reflect.Type) bool { _ = "STUB: not implemented"; return false }
