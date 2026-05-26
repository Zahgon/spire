package tpmdevid

import (
	"context"
	"sync"

	"github.com/hashicorp/go-hclog"
	nodeattestorv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/agent/nodeattestor/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/agent/plugin/nodeattestor/tpmdevid/tpmutil"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
)

const BaseTPMDir = "/dev"

// Functions defined here are overridden in test files to facilitate unit testing
var (
	AutoDetectTPMPath func(string) (string, error)                           = tpmutil.AutoDetectTPMPath
	NewSession        func(*tpmutil.SessionConfig) (*tpmutil.Session, error) = tpmutil.NewSession
)

func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *Plugin) catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

type Config struct {
	DevIDPrivPath string `hcl:"devid_priv_path"`
	DevIDPubPath  string `hcl:"devid_pub_path"`
	DevIDCertPath string `hcl:"devid_cert_path"`

	DevIDKeyPassword             string `hcl:"devid_password"`
	OwnerHierarchyPassword       string `hcl:"owner_hierarchy_password"`
	EndorsementHierarchyPassword string `hcl:"endorsement_hierarchy_password"`

	DevicePath string `hcl:"tpm_device_path"`
	Autodetect bool
}

func buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *Config {
	_ = "STUB: not implemented"
	return nil
}

type config struct {
	devicePath string
	devIDCert  [][]byte
	devIDPub   []byte
	devIDPriv  []byte
	passwords  tpmutil.TPMPasswords
}

type Plugin struct {
	nodeattestorv1.UnsafeNodeAttestorServer
	configv1.UnsafeConfigServer
	log hclog.Logger

	m sync.Mutex
	c *config
}

func New() *Plugin { _ = "STUB: not implemented"; return nil }

func (p *Plugin) AidAttestation(stream nodeattestorv1.NodeAttestor_AidAttestationServer) error {
	_ = "STUB: not implemented"
	return nil
}

// Open TPM connection and load DevID keys

// Get endorsement certificate from TPM NV index

// Get regenerated endorsement public key

// Certify DevID is in the same TPM than AK

// Marshal attestation data

// Send attestation request

// Receive challenges

// Solve DevID challenge (verify the possession of the DevID private key)

// Solve Credential Activation challenge

// Marshal challenges responses

// Send challenge response back to the server

func (p *Plugin) Configure(_ context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) Validate(_ context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) SetLogger(log hclog.Logger) { _ = "STUB: not implemented"; return }

func (p *Plugin) getConfig() *config { _ = "STUB: not implemented"; return nil }

func (p *Plugin) loadDevIDFiles(c *Config) error { _ = "STUB: not implemented"; return nil }
