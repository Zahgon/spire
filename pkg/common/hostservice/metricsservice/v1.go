package metricsservice

import (
	"context"

	metricsv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/hostservice/common/metrics/v1"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"google.golang.org/protobuf/types/known/emptypb"
)

// V1 returns a v1 metrics service server over the provided Metrics interface
func V1(metrics telemetry.Metrics) metricsv1.MetricsServer {
	_ = "STUB: not implemented"
	return *new(metricsv1.MetricsServer)
}

type metricsV1 struct {
	metricsv1.UnsafeMetricsServer
	metrics telemetry.Metrics
}

func (m metricsV1) AddSample(_ context.Context, req *metricsv1.AddSampleRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m metricsV1) EmitKey(_ context.Context, req *metricsv1.EmitKeyRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m metricsV1) IncrCounter(_ context.Context, req *metricsv1.IncrCounterRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m metricsV1) MeasureSince(_ context.Context, req *metricsv1.MeasureSinceRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m metricsV1) SetGauge(_ context.Context, req *metricsv1.SetGaugeRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func v1ConvertToRPCLabels(inLabels []telemetry.Label) []*metricsv1.Label {
	_ = "STUB: not implemented"
	return nil
}

func v1ConvertToTelemetryLabels(inLabels []*metricsv1.Label) []telemetry.Label {
	_ = "STUB: not implemented"
	return nil
}
