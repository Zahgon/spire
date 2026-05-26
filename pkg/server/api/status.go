package api

import (
	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"google.golang.org/grpc/codes"
)

// CreateStatus creates a proto Status
func CreateStatus(code codes.Code, msg string) *types.Status { _ = "STUB: not implemented"; return nil }

// CreateStatus creates a proto Status
func CreateStatusf(code codes.Code, format string, a ...any) *types.Status {
	_ = "STUB: not implemented"
	return nil
}

// OK creates a success proto status
func OK() *types.Status { _ = "STUB: not implemented"; return nil }

// MakeStatus logs and returns a status composed of: msg, err and code.
// Errors are treated differently according to its gRPC code.
func MakeStatus(log logrus.FieldLogger, code codes.Code, msg string, err error) *types.Status {
	_ = "STUB: not implemented"
	return nil
}

// MakeErr logs and returns an error composed of: msg, err and code.
// Errors are treated differently according to its gRPC code.
func MakeErr(log logrus.FieldLogger, code codes.Code, msg string, err error) error {
	_ = "STUB: not implemented"
	return nil
}

// It is not expected for MakeErr to be called with nil
// but we make a case for it in the switch to prevent it to
// go to the default case

// Add the prefix 'Invalid argument' for InvalidArgument errors

// Do not log nor return the inner error for NotFound errors

// Concat message with provided error and avoid "status.Code"
func concatErr(msg string, err error) string { _ = "STUB: not implemented"; return "" }

// Proto will be nil "only" when err is nil

func capitalize(s string) string { _ = "STUB: not implemented"; return "" }
