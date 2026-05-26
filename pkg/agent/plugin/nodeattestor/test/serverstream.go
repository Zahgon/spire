package nodeattestortest

import (
	"context"

	"github.com/spiffe/spire/pkg/agent/plugin/nodeattestor"
)

// ServerStreamHandler is a function used to handle payloads or challenge
// responses sent to the stream.
type ServerStreamHandler = func(payloadOrChallengeResponse []byte) (challenge []byte, err error)

// ServerStreamBuilder is used to build server streams for testing.
type ServerStreamBuilder struct {
	pluginName string
	handlers   []ServerStreamHandler
}

// ServerStream initializes a new server stream builder for the given plugin
// name. Attestation data received by the stream will have its type validated
// against the plugin name.
func ServerStream(pluginName string) *ServerStreamBuilder { _ = "STUB: not implemented"; return nil }

// Build builds a stream with the configured handlers
func (b *ServerStreamBuilder) Build() nodeattestor.ServerStream {
	_ = "STUB: not implemented"
	return *new(nodeattestor.ServerStream)
}

// Handle adds an arbitrary handler. If the handler returns a challenge then it
// is expected that the stream will be called again.
func (b *ServerStreamBuilder) Handle(handler ServerStreamHandler) *ServerStreamBuilder {
	_ = "STUB: not implemented"
	return nil
}

// ExpectThenChallenge adds an intermediate handler that asserts that the given
// payload or challenge response is received and then issues the given
// challenge. It returns a new builder with that handler added.
func (b *ServerStreamBuilder) ExpectThenChallenge(payloadOrChallengeResponse, challenge []byte) *ServerStreamBuilder {
	_ = "STUB: not implemented"
	return nil
}

// IgnoreThenChallenge adds an intermediate handler that ignores the payload or
// challenge response and then issues the given challenge. It returns a new
// builder with that handler added.
func (b *ServerStreamBuilder) IgnoreThenChallenge(challenge []byte) *ServerStreamBuilder {
	_ = "STUB: not implemented"
	return nil
}

// ExpectAndBuild adds a final handler wherein the server stream expects to
// receive the given payload or challenge response. It returns a built server
// stream, since the stream does not issue another challenge at this point and
// will fail if invoked again.
func (b *ServerStreamBuilder) ExpectAndBuild(payloadOrChallengeResponse []byte) nodeattestor.ServerStream {
	_ = "STUB: not implemented"
	return *new(nodeattestor.ServerStream)
}

// FailAndBuild adds a final handler wherein the server stream fails with the
// given error.  It returns a built server stream, since the stream does not
// issue another challenge at this point and will fail if invoked again.
func (b *ServerStreamBuilder) FailAndBuild(err error) nodeattestor.ServerStream {
	_ = "STUB: not implemented"
	return *new(nodeattestor.ServerStream)
}

func (b *ServerStreamBuilder) addHandler(handler ServerStreamHandler) *ServerStreamBuilder {
	_ = "STUB: not implemented"
	return nil
}

type serverStream struct {
	pluginName string
	handlers   []ServerStreamHandler
}

func (ss *serverStream) SendAttestationData(_ context.Context, attestationData nodeattestor.AttestationData) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ss *serverStream) SendChallengeResponse(_ context.Context, challengeResponse []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ss *serverStream) handle(payloadOrChallengeResponse []byte) (challenge []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
