package spiretest

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

func Run(t *testing.T, s suite.TestingSuite) { _ = "STUB: not implemented"; return }

type Suite struct {
	suite.Suite
}

func (s *Suite) Cleanup(cleanup func()) { _ = "STUB: not implemented"; return }

func (s *Suite) TempDir() string { _ = "STUB: not implemented"; return "" }

func (s *Suite) RequireErrorContains(err error, contains string) { _ = "STUB: not implemented"; return }

func (s *Suite) RequireGRPCStatus(err error, code codes.Code, message string) {
	_ = "STUB: not implemented"
	return
}

func (s *Suite) RequireGRPCStatusContains(err error, code codes.Code, contains string) {
	_ = "STUB: not implemented"
	return
}

func (s *Suite) RequireProtoListEqual(expected, actual any) { _ = "STUB: not implemented"; return }

func (s *Suite) RequireProtoEqual(expected, actual proto.Message) {
	_ = "STUB: not implemented"
	return
}

func (s *Suite) AssertErrorContains(err error, contains string) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *Suite) AssertGRPCStatus(err error, code codes.Code, message string) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *Suite) AssertGRPCStatusContains(err error, code codes.Code, contains string) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *Suite) AssertProtoListEqual(expected, actual any) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *Suite) AssertProtoEqual(expected, actual proto.Message, msgAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *Suite) CheckProtoListEqual(expected, actual any) bool {
	_ = "STUB: not implemented"
	return false
}
