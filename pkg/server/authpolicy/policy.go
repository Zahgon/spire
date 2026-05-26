package authpolicy

import (
	"context"

	"github.com/open-policy-agent/opa/v1/rego"
	"github.com/open-policy-agent/opa/v1/storage"
	"github.com/sirupsen/logrus"
)

const (
	allowKey             = "allow"
	allowIfAdminKey      = "allow_if_admin"
	allowIfDownstreamKey = "allow_if_downstream"
	allowIfAgentKey      = "allow_if_agent"
	allowIfLocalKey      = "allow_if_local"
)

// Engine drives policy management.
type Engine struct {
	query rego.PreparedEvalQuery
}

type OpaEngineConfig struct {
	LocalOpaProvider *LocalOpaProviderConfig `hcl:"local"`
}

type LocalOpaProviderConfig struct {
	RegoPath       string `hcl:"rego_path"`
	PolicyDataPath string `hcl:"policy_data_path"`
}

// Input represents context associated with an access request.
type Input struct {
	// Caller is the authenticated identity of the actor making a request.
	Caller string `json:"caller"`

	// CallerFilePath is the file path of a local actor making a request.
	CallerFilePath string `json:"caller_file_path"`

	// FullMethod is the fully-qualified name of the proto rpc service method.
	FullMethod string `json:"full_method"`

	// Req represents data received from the request body. It MUST be a
	// protobuf request object with fields that are serializable as JSON,
	// since they will be used in policy definitions.
	Req any `json:"req"`
}

type Result struct {
	Allow             bool `json:"allow"`
	AllowIfAdmin      bool `json:"allow_if_admin"`
	AllowIfLocal      bool `json:"allow_if_local"`
	AllowIfDownstream bool `json:"allow_if_downstream"`
	AllowIfAgent      bool `json:"allow_if_agent"`
}

// NewEngineFromConfigOrDefault returns a new policy engine. Or if no
// config is provided, provides the default policy
func NewEngineFromConfigOrDefault(ctx context.Context, logger logrus.FieldLogger, cfg *OpaEngineConfig) (*Engine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// newEngine returns a new policy engine. Or nil if no
// config is provided.
func newEngine(ctx context.Context, cfg *OpaEngineConfig) (*Engine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If permissions file is defined use it, else provide empty store

// NewEngineFromRego is a helper to create the Engine object
func NewEngineFromRego(ctx context.Context, regoPolicy string, dataStore storage.Store) (*Engine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Test policy with some simple calls to ensure that the
// policy can be evaluated properly.

// Eval determines whether access should be allowed on a resource.
func (e *Engine) Eval(ctx context.Context, input Input) (result Result, err error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}
