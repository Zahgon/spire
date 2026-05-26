package authoritycommontest

import (
	"bytes"
	"context"
	"testing"

	"github.com/mitchellh/cli"
	localauthorityv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/localauthority/v1"
	commoncli "github.com/spiffe/spire/pkg/common/cli"
)

var AvailableFormats = []string{"pretty", "json"}

type localAuthorityTest struct {
	Stdin  *bytes.Buffer
	Stdout *bytes.Buffer
	Stderr *bytes.Buffer
	Args   []string
	Server *fakeLocalAuthorityServer
	Client cli.Command
}

func (s *localAuthorityTest) afterTest(t *testing.T) { _ = "STUB: not implemented"; return }

func SetupTest(t *testing.T, newClient func(*commoncli.Env) cli.Command) *localAuthorityTest {
	_ = "STUB: not implemented"
	return nil
}

type fakeLocalAuthorityServer struct {
	localauthorityv1.UnsafeLocalAuthorityServer

	ActiveJWT,
	PreparedJWT,
	OldJWT,
	ActiveX509,
	PreparedX509,
	OldX509,
	TaintedX509,
	RevokedX509,
	TaintedJWT,
	RevokedJWT *localauthorityv1.AuthorityState

	TaintedUpstreamAuthoritySubjectKeyId,
	RevokedUpstreamAuthoritySubjectKeyId string
	Err error
}

func (s *fakeLocalAuthorityServer) GetJWTAuthorityState(context.Context, *localauthorityv1.GetJWTAuthorityStateRequest) (*localauthorityv1.GetJWTAuthorityStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *fakeLocalAuthorityServer) PrepareJWTAuthority(context.Context, *localauthorityv1.PrepareJWTAuthorityRequest) (*localauthorityv1.PrepareJWTAuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *fakeLocalAuthorityServer) ActivateJWTAuthority(context.Context, *localauthorityv1.ActivateJWTAuthorityRequest) (*localauthorityv1.ActivateJWTAuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *fakeLocalAuthorityServer) TaintJWTAuthority(context.Context, *localauthorityv1.TaintJWTAuthorityRequest) (*localauthorityv1.TaintJWTAuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *fakeLocalAuthorityServer) RevokeJWTAuthority(context.Context, *localauthorityv1.RevokeJWTAuthorityRequest) (*localauthorityv1.RevokeJWTAuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *fakeLocalAuthorityServer) GetX509AuthorityState(context.Context, *localauthorityv1.GetX509AuthorityStateRequest) (*localauthorityv1.GetX509AuthorityStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *fakeLocalAuthorityServer) PrepareX509Authority(context.Context, *localauthorityv1.PrepareX509AuthorityRequest) (*localauthorityv1.PrepareX509AuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *fakeLocalAuthorityServer) ActivateX509Authority(context.Context, *localauthorityv1.ActivateX509AuthorityRequest) (*localauthorityv1.ActivateX509AuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *fakeLocalAuthorityServer) TaintX509Authority(context.Context, *localauthorityv1.TaintX509AuthorityRequest) (*localauthorityv1.TaintX509AuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *fakeLocalAuthorityServer) TaintX509UpstreamAuthority(context.Context, *localauthorityv1.TaintX509UpstreamAuthorityRequest) (*localauthorityv1.TaintX509UpstreamAuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *fakeLocalAuthorityServer) RevokeX509Authority(context.Context, *localauthorityv1.RevokeX509AuthorityRequest) (*localauthorityv1.RevokeX509AuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *fakeLocalAuthorityServer) RevokeX509UpstreamAuthority(context.Context, *localauthorityv1.RevokeX509UpstreamAuthorityRequest) (*localauthorityv1.RevokeX509UpstreamAuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *fakeLocalAuthorityServer) GetWITAuthorityState(ctx context.Context, _ *localauthorityv1.GetWITAuthorityStateRequest) (*localauthorityv1.GetWITAuthorityStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *fakeLocalAuthorityServer) PrepareWITAuthority(ctx context.Context, _ *localauthorityv1.PrepareWITAuthorityRequest) (*localauthorityv1.PrepareWITAuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *fakeLocalAuthorityServer) ActivateWITAuthority(ctx context.Context, req *localauthorityv1.ActivateWITAuthorityRequest) (*localauthorityv1.ActivateWITAuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *fakeLocalAuthorityServer) TaintWITAuthority(ctx context.Context, req *localauthorityv1.TaintWITAuthorityRequest) (*localauthorityv1.TaintWITAuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *fakeLocalAuthorityServer) RevokeWITAuthority(ctx context.Context, req *localauthorityv1.RevokeWITAuthorityRequest) (*localauthorityv1.RevokeWITAuthorityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func RequireOutputBasedOnFormat(t *testing.T, format, stdoutString string, expectedStdoutPretty, expectedStdoutJSON string) {
	_ = "STUB: not implemented"
	return
}
