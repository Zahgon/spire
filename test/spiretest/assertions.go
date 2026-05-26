package spiretest

import (
	"reflect"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

var (
	protoMessageType = reflect.TypeFor[proto.Message]()
)

func RequireErrorContains(tb testing.TB, err error, contains string) {
	_ = "STUB: not implemented"
	return
}

func AssertErrorContains(tb testing.TB, err error, contains string) bool {
	_ = "STUB: not implemented"
	return false
}

func RequireGRPCStatus(tb testing.TB, err error, code codes.Code, message string) {
	_ = "STUB: not implemented"
	return
}

func AssertGRPCStatus(tb testing.TB, err error, code codes.Code, message string) bool {
	_ = "STUB: not implemented"
	return false
}

func RequireGRPCStatusContains(tb testing.TB, err error, code codes.Code, contains string, msgAndArgs ...any) {
	_ = "STUB: not implemented"
	return
}

func AssertGRPCStatusContains(tb testing.TB, err error, code codes.Code, contains string, msgAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

func RequireGRPCStatusHasPrefix(tb testing.TB, err error, code codes.Code, prefix string) {
	_ = "STUB: not implemented"
	return
}

func AssertGRPCStatusHasPrefix(tb testing.TB, err error, code codes.Code, prefix string) bool {
	_ = "STUB: not implemented"
	return false
}

func RequireProtoListEqual(tb testing.TB, expected, actual any) { _ = "STUB: not implemented"; return }

func AssertProtoListEqual(tb testing.TB, expected, actual any) bool {
	_ = "STUB: not implemented"
	return false
}

func CheckProtoListEqual(tb testing.TB, expected, actual any) bool {
	_ = "STUB: not implemented"
	return false
}

func RequireProtoEqual(tb testing.TB, expected, actual proto.Message, msgAndArgs ...any) {
	_ = "STUB: not implemented"
	return
}

func AssertProtoEqual(tb testing.TB, expected, actual proto.Message, msgAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

func CheckProtoEqual(tb testing.TB, expected, actual proto.Message) bool {
	_ = "STUB: not implemented"
	return false
}

func RequireErrorPrefix(tb testing.TB, err error, prefix string) { _ = "STUB: not implemented"; return }

func AssertErrorPrefix(tb testing.TB, err error, prefix string) bool {
	_ = "STUB: not implemented"
	return false
}

func AssertHasPrefix(tb testing.TB, msg string, prefix string) bool {
	_ = "STUB: not implemented"
	return false
}
