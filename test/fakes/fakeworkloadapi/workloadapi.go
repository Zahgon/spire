package fakeworkloadapi

import (
	"context"
	"net"
	"testing"

	"github.com/spiffe/go-spiffe/v2/proto/spiffe/workload"
	"google.golang.org/protobuf/proto"
)

type FakeRequest struct {
	Req  proto.Message
	Resp proto.Message
	Err  error
}

type WorkloadAPI struct {
	workload.UnimplementedSpiffeWorkloadAPIServer
	addr net.Addr
	t    *testing.T

	ExpFetchJWTSVIDReq    *workload.JWTSVIDRequest
	ExpFetchJWTBundlesReq *workload.JWTBundlesRequest

	fetchX509SVIDRequest   FakeRequest
	fetchJWTSVIDRequest    FakeRequest
	fetchJWTBundlesRequest FakeRequest
	validateJWTRequest     FakeRequest
}

func New(t *testing.T, responses ...*FakeRequest) *WorkloadAPI {
	_ = "STUB: not implemented"
	return nil
}

func (w *WorkloadAPI) Addr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (w *WorkloadAPI) FetchX509SVID(req *workload.X509SVIDRequest, stream workload.SpiffeWorkloadAPI_FetchX509SVIDServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *WorkloadAPI) FetchJWTSVID(_ context.Context, req *workload.JWTSVIDRequest) (*workload.JWTSVIDResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *WorkloadAPI) FetchJWTBundles(req *workload.JWTBundlesRequest, stream workload.SpiffeWorkloadAPI_FetchJWTBundlesServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *WorkloadAPI) ValidateJWTSVID(_ context.Context, req *workload.ValidateJWTSVIDRequest) (*workload.ValidateJWTSVIDResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkSecurityHeader(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Ensure security header is sent
	return nil
}
