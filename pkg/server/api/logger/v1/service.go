package logger

import (
	"context"

	"github.com/sirupsen/logrus"
	loggerv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/logger/v1"
	apitype "github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"google.golang.org/grpc"
)

type Logger interface {
	logrus.FieldLogger

	GetLevel() logrus.Level
	SetLevel(level logrus.Level)
}

func RegisterService(s grpc.ServiceRegistrar, service *Service) { _ = "STUB: not implemented"; return }

type Config struct {
	Log Logger
}

type Service struct {
	loggerv1.UnsafeLoggerServer

	log         Logger
	launchLevel logrus.Level
}

func New(c Config) *Service { _ = "STUB: not implemented"; return nil }

func (s *Service) GetLogger(ctx context.Context, _ *loggerv1.GetLoggerRequest) (*apitype.Logger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) SetLogLevel(ctx context.Context, req *loggerv1.SetLogLevelRequest) (*apitype.Logger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) ResetLogLevel(ctx context.Context, _ *loggerv1.ResetLogLevelRequest) (*apitype.Logger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) createAPILogger() *apitype.Logger { _ = "STUB: not implemented"; return nil }
