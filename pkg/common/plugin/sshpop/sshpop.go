// Package sshpop implements ssh proof of possession based node attestation.
package sshpop

import (
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/agentpathtemplate"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
	"golang.org/x/crypto/ssh"
)

const (
	// PluginName is used for identifying this plugin type for protobuf blobs.
	PluginName = "sshpop"

	defaultHostKeyPath  = "/etc/ssh/ssh_host_rsa_key"
	defaultHostCertPath = "/etc/ssh/ssh_host_rsa_key-cert.pub"
	nonceLen            = 32
)

var (
	// DefaultAgentPathTemplate is the default text/template.
	DefaultAgentPathTemplate = agentpathtemplate.MustParse("/{{ .PluginName}}/{{ .Fingerprint }}")
)

// agentPathTemplateData is used to hydrate the agent path template used in generating spiffe ids.
type agentPathTemplateData struct {
	*ssh.Certificate
	PluginName  string
	Fingerprint string
	Hostname    string
}

// Client is a factory for generating client handshake objects.
type Client struct {
	cert   *ssh.Certificate
	signer ssh.Signer
}

// Server is a factory for generating server handshake objects.
type Server struct {
	certChecker       *ssh.CertChecker
	agentPathTemplate *agentpathtemplate.Template
	trustDomain       spiffeid.TrustDomain
	canonicalDomain   string
}

// ClientConfig configures the client.
type ClientConfig struct {
	HostKeyPath  string `hcl:"host_key_path"`
	HostCertPath string `hcl:"host_cert_path"`

	cert   *ssh.Certificate
	signer ssh.Signer
}

type ClientConfigRequest struct {
	coreConfig *configv1.CoreConfiguration
	hclText    string
}

func (ccr *ClientConfigRequest) GetCoreConfiguration() *configv1.CoreConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (ccr *ClientConfigRequest) GetHclConfiguration() string { _ = "STUB: not implemented"; return "" }

type ServerConfigRequest struct {
	coreConfig *configv1.CoreConfiguration
	hclText    string
}

func (scr *ServerConfigRequest) GetCoreConfiguration() *configv1.CoreConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (scr *ServerConfigRequest) GetHclConfiguration() string {
	_ = "STUB: not implemented"

	// ServerConfig configures the server.
	return ""
}

type ServerConfig struct {
	CertAuthorities     []string `hcl:"cert_authorities"`
	CertAuthoritiesPath string   `hcl:"cert_authorities_path"`
	// CanonicalDomain specifies the domain suffix for validating the hostname against
	// the certificate's valid principals. See CanonicalDomains in ssh_config(5).
	CanonicalDomain   string `hcl:"canonical_domain"`
	AgentPathTemplate string `hcl:"agent_path_template"`

	certChecker       *ssh.CertChecker
	agentPathTemplate *agentpathtemplate.Template
	trustDomain       spiffeid.TrustDomain
}

func BuildServerConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *ServerConfig {
	_ = "STUB: not implemented"
	return nil
}

func (sc *ServerConfig) NewServer() *Server { _ = "STUB: not implemented"; return nil }

func BuildClientConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *ClientConfig {
	_ = "STUB: not implemented"
	return nil
}

func (cc *ClientConfig) NewClient() *Client { _ = "STUB: not implemented"; return nil }

func NewClient(trustDomain string, configString string) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func stringOrDefault(configValue, defaultValue string) string { _ = "STUB: not implemented"; return "" }

func getCertAndSignerFromBytes(certBytes, keyBytes []byte) (*ssh.Certificate, ssh.Signer, error) {
	_ = "STUB: not implemented"
	return nil, *new(ssh.Signer), nil
}

func NewServer(trustDomain, configString string) (*Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func pubkeysFromPath(pubkeysPath string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func certCheckerFromPubkeys(certAuthorities []string) (*ssh.CertChecker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) NewHandshake() *ClientHandshake { _ = "STUB: not implemented"; return nil }

func (s *Server) NewHandshake() *ServerHandshake { _ = "STUB: not implemented"; return nil }
