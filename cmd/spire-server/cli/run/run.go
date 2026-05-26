package run

import (
	"context"
	"crypto/x509/pkix"
	"io"
	"time"

	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/hashicorp/hcl/hcl/token"
	"github.com/mitchellh/cli"
	"github.com/sirupsen/logrus"
	common_cli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/diskcertmanager"
	"github.com/spiffe/spire/pkg/common/fflag"
	"github.com/spiffe/spire/pkg/common/health"
	"github.com/spiffe/spire/pkg/common/log"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/pkg/server"
	"github.com/spiffe/spire/pkg/server/authpolicy"
	bundleClient "github.com/spiffe/spire/pkg/server/bundle/client"
	"github.com/spiffe/spire/pkg/server/endpoints/bundle"
)

const (
	commandName = "run"

	defaultConfigPath = "conf/server/server.conf"
	defaultLogLevel   = "INFO"
)

var defaultRateLimit = true

// Config contains all available configurables, arranged by section
type Config struct {
	Server             *serverConfig          `hcl:"server"`
	Plugins            ast.Node               `hcl:"plugins"`
	Telemetry          telemetry.FileConfig   `hcl:"telemetry"`
	HealthChecks       health.Config          `hcl:"health_checks"`
	UnusedKeyPositions map[string][]token.Pos `hcl:",unusedKeyPositions"`
}

type serverConfig struct {
	AdminIDs                     []string           `hcl:"admin_ids"`
	AgentTTL                     string             `hcl:"agent_ttl"`
	AuditLogEnabled              bool               `hcl:"audit_log_enabled"`
	BindAddress                  string             `hcl:"bind_address"`
	BindPort                     int                `hcl:"bind_port"`
	CAKeyType                    string             `hcl:"ca_key_type"`
	CASubject                    *caSubjectConfig   `hcl:"ca_subject"`
	CATTL                        string             `hcl:"ca_ttl"`
	DataDir                      string             `hcl:"data_dir"`
	DefaultX509SVIDTTL           string             `hcl:"default_x509_svid_ttl"`
	DefaultJWTSVIDTTL            string             `hcl:"default_jwt_svid_ttl"`
	Experimental                 experimentalConfig `hcl:"experimental"`
	Federation                   *federationConfig  `hcl:"federation"`
	DisableJWTSVIDs              bool               `hcl:"disable_jwt_svids"`
	JWTIssuer                    string             `hcl:"jwt_issuer"`
	JWTKeyType                   string             `hcl:"jwt_key_type"`
	LogFile                      string             `hcl:"log_file"`
	LogLevel                     string             `hcl:"log_level"`
	LogFormat                    string             `hcl:"log_format"`
	LogSourceLocation            bool               `hcl:"log_source_location"`
	PruneAttestedNodesExpiredFor string             `hcl:"prune_attested_nodes_expired_for"`
	PruneNonReattestableNodes    bool               `hcl:"prune_tofu_nodes"`
	ProxyProtocolTrustedCIDRs    []string           `hcl:"proxy_protocol_trusted_cidrs"`
	RateLimit                    rateLimitConfig    `hcl:"ratelimit"`
	SocketPath                   string             `hcl:"socket_path"`
	TrustDomain                  string             `hcl:"trust_domain"`
	MaxAttestedNodeInfoStaleness *string            `hcl:"max_attested_node_info_staleness"`

	ConfigPath string
	ExpandEnv  bool

	// Undocumented configurables
	ProfilingEnabled bool     `hcl:"profiling_enabled"`
	ProfilingPort    int      `hcl:"profiling_port"`
	ProfilingFreq    int      `hcl:"profiling_freq"`
	ProfilingNames   []string `hcl:"profiling_names"`

	UnusedKeyPositions map[string][]token.Pos `hcl:",unusedKeyPositions"`
}

type experimentalConfig struct {
	AgentSpiffeIdAsSelector bool                        `hcl:"agent_spiffe_id_as_selector"`
	AuthOpaPolicyEngine     *authpolicy.OpaEngineConfig `hcl:"auth_opa_policy_engine"`
	CacheReloadInterval     string                      `hcl:"cache_reload_interval"`
	FullCacheReloadInterval string                      `hcl:"full_cache_reload_interval"`
	EventsBasedCache        bool                        `hcl:"events_based_cache"`
	PruneEventsOlderThan    string                      `hcl:"prune_events_older_than"`
	EventTimeout            string                      `hcl:"event_timeout"`
	SQLTransactionTimeout   string                      `hcl:"sql_transaction_timeout"`
	RequirePQKEM            bool                        `hcl:"require_pq_kem"`
	WITKeyType              string                      `hcl:"wit_key_type"`
	WITIssuer               string                      `hcl:"wit_issuer"`

	Flags fflag.RawConfig `hcl:"feature_flags"`

	NamedPipeName string `hcl:"named_pipe_name"`

	UnusedKeyPositions map[string][]token.Pos `hcl:",unusedKeyPositions"`
}

type caSubjectConfig struct {
	Country            []string               `hcl:"country"`
	Organization       []string               `hcl:"organization"`
	CommonName         string                 `hcl:"common_name"`
	UnusedKeyPositions map[string][]token.Pos `hcl:",unusedKeyPositions"`
}

type federationConfig struct {
	BundleEndpoint     *bundleEndpointConfig          `hcl:"bundle_endpoint"`
	FederatesWith      map[string]federatesWithConfig `hcl:"federates_with"`
	UnusedKeyPositions map[string][]token.Pos         `hcl:",unusedKeyPositions"`
}

type bundleEndpointConfig struct {
	Address     string `hcl:"address"`
	Port        int    `hcl:"port"`
	RefreshHint string `hcl:"refresh_hint"`

	ACME    *bundleEndpointACMEConfig `hcl:"acme"`
	Profile ast.Node                  `hcl:"profile"`

	UnusedKeyPositions map[string][]token.Pos `hcl:",unusedKeyPositions"`
}

type bundleEndpointConfigProfile struct {
	HTTPSSPIFFE        *bundleEndpointProfileHTTPSSPIFFEConfig `hcl:"https_spiffe"`
	HTTPSWeb           *bundleEndpointProfileHTTPSWebConfig    `hcl:"https_web"`
	UnusedKeyPositions map[string][]token.Pos                  `hcl:",unusedKeyPositions"`
}

type bundleEndpointProfileHTTPSWebConfig struct {
	ACME            *bundleEndpointACMEConfig      `hcl:"acme"`
	ServingCertFile *bundleEndpointServingCertFile `hcl:"serving_cert_file"`
}

type bundleEndpointProfileHTTPSSPIFFEConfig struct{}

type bundleEndpointServingCertFile struct {
	CertFilePath        string        `hcl:"cert_file_path"`
	KeyFilePath         string        `hcl:"key_file_path"`
	FileSyncInterval    time.Duration `hcl:"-"`
	RawFileSyncInterval string        `hcl:"file_sync_interval"`
}

type bundleEndpointACMEConfig struct {
	DirectoryURL       string                 `hcl:"directory_url"`
	DomainName         string                 `hcl:"domain_name"`
	Email              string                 `hcl:"email"`
	ToSAccepted        bool                   `hcl:"tos_accepted"`
	UnusedKeyPositions map[string][]token.Pos `hcl:",unusedKeyPositions"`
}

type federatesWithConfig struct {
	BundleEndpointURL     string                 `hcl:"bundle_endpoint_url"`
	BundleEndpointProfile ast.Node               `hcl:"bundle_endpoint_profile"`
	UnusedKeyPositions    map[string][]token.Pos `hcl:",unusedKeyPositions"`
}

type bundleEndpointProfileConfig struct {
	HTTPSSPIFFE        *httpsSPIFFEProfileConfig `hcl:"https_spiffe"`
	HTTPSWeb           *httpsWebProfileConfig    `hcl:"https_web"`
	UnusedKeyPositions map[string][]token.Pos    `hcl:",unusedKeyPositions"`
}

type httpsSPIFFEProfileConfig struct {
	EndpointSPIFFEID   string                 `hcl:"endpoint_spiffe_id"`
	UnusedKeyPositions map[string][]token.Pos `hcl:",unusedKeyPositions"`
}

type httpsWebProfileConfig struct{}

type rateLimitConfig struct {
	Attestation        *bool                  `hcl:"attestation"`
	Signing            *bool                  `hcl:"signing"`
	UnusedKeyPositions map[string][]token.Pos `hcl:",unusedKeyPositions"`
}

func NewRunCommand(ctx context.Context, logOptions []log.Option, allowUnknownConfig bool) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

func newRunCommand(ctx context.Context, env *common_cli.Env, logOptions []log.Option, allowUnknownConfig bool) *Command {
	_ = "STUB: not implemented"
	return nil
}

// Run Command struct
type Command struct {
	ctx                context.Context
	logOptions         []log.Option
	env                *common_cli.Env
	allowUnknownConfig bool
}

// Help prints the server cmd usage
func (cmd *Command) Help() string { _ = "STUB: not implemented"; return "" }

// Help is a standalone function that prints a help message to writer.
// It is used by both the run and validate commands, so they can share flag usage messages.
func Help(name string, writer io.Writer) string { _ = "STUB: not implemented"; return "" }

// Error is always present because -h is passed

func LoadConfig(name string, args []string, logOptions []log.Option, output io.Writer, allowUnknownConfig bool) (*server.Config, error) {
	_ = "STUB: not implemented"
	// First parse the CLI flags so we can get the config
	// file path, if set
	return nil, nil
}

// Load and parse the config file using either the default
// path or CLI-specified value

// Run the SPIFFE Server
func (cmd *Command) Run(args []string) int { _ = "STUB: not implemented"; return 0 }

// Set umask before starting up the server

// Synopsis of the command
func (*Command) Synopsis() string { _ = "STUB: not implemented"; return "" }

func ParseFile(path string, expandEnv bool) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Return a friendly error if the file is missing

// If envTemplate flag is passed, substitute $VARIABLES in configuration file

func parseFlags(name string, args []string, output io.Writer) (*serverConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mergeInput(fileInput *Config, cliInput *serverConfig) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Highest precedence first

func NewServerConfig(c *Config, logOptions []log.Option, allowUnknownConfig bool) (*server.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If neither new nor deprecated config value is set, then use hard-coded default TTL
// Note, due to back-compat issues we cannot set this default inside defaultConfig() function

// If not set using new field then use hard-coded default TTL
// Note, due to back-compat issues we cannot set this default inside defaultConfig() function

// If the configured TTLs can lead to surprises, then do our best to log an
// accurate message and guide the user to resolution

// TTL is smaller than our cap, but the CA TTL
// is not large enough to accommodate it

// TTL is larger than our cap, it needs to be
// decreased no matter what. Additionally, the CA TTL is
// too small to accommodate the maximum SVID TTL.

// TTL is larger than our cap and needs to be
// decreased.

// RFC3280(4.1.2.4) requires the issuer DN be set.

func setBundleEndpointConfigProfile(config *bundleEndpointConfig, dataDir string, log logrus.FieldLogger, federationConfig *server.FederationConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Profile is set, parse it

// For now ignore SPIFFE configuration

func configToACMEConfig(acme *bundleEndpointACMEConfig, dataDir string) *bundle.ACMEConfig {
	_ = "STUB: not implemented"
	return nil
}

func configToDiskCertManager(serviceCertFile *bundleEndpointServingCertFile, log logrus.FieldLogger) (*diskcertmanager.DiskCertManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseBundleEndpointProfile(config federatesWithConfig) (trustDomainConfig *bundleClient.TrustDomainConfig, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseBundleEndpointProfileASTNode(node ast.Node) (string, error) {
	_ = "STUB: not implemented"
	// First check the number of bundle endpoint profiles in the config
	return "", nil
}

func validateConfig(c *Config) error { _ = "STUB: not implemented"; return nil }

func checkForUnknownConfig(c *Config, l logrus.FieldLogger) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Re-enable unused key detection for experimental config. See
// https://github.com/spiffe/spire/issues/1101 for more information
//
// if len(c.Server.Experimental.UnusedKeyPositions) != 0 {
//	detectedUnknown("experimental", c.Server.Experimental.UnusedKeyPositions)
// }

// TODO: Re-enable unused key detection for federation config. See
// https://github.com/spiffe/spire/issues/1101 for more information
//
// if len(c.Server.Federation.UnusedKeyPositions) != 0 {
//	detectedUnknown("federation", c.Server.Federation.UnusedKeyPositions)
// }

// TODO: Re-enable unused key detection for bundle endpoint profile config. See
// https://github.com/spiffe/spire/issues/1101 for more information
//
// for k, v := range c.Server.Federation.FederatesWith {
//	if len(v.UnusedKeyPositions) != 0 {
//		detectedUnknown(fmt.Sprintf("federates_with %q", k), v.UnusedKeyPositions)
//	}
// }

// TODO: Re-enable unused key detection for telemetry. See
// https://github.com/spiffe/spire/issues/1101 for more information
//
// if len(c.Telemetry.UnusedKeyPositions) != 0 {
//	detectedUnknown("telemetry", c.Telemetry.UnusedKeyPositions)
// }

func defaultConfig() *Config { _ = "STUB: not implemented"; return nil }

// hasCompatibleTTL checks if we can guarantee the configured SVID TTL given the
// configured CA TTL. If we detect that a new SVID TTL may be cut short due to
// a scheduled CA rotation, this function will return false. This method should
// be called for each SVID TTL we may use
func hasCompatibleTTL(caTTL time.Duration, svidTTL time.Duration) bool {
	_ = "STUB: not implemented"
	return false
}

// printMaxSVIDTTL calculates the display string for a sufficiently short SVID TTL
func printMaxSVIDTTL(caTTL time.Duration) string { _ = "STUB: not implemented"; return "" }

// printMinCATTL calculates the display string for a sufficiently large CA TTL
func printMinCATTL(svidTTL time.Duration) string { _ = "STUB: not implemented"; return "" }

func printDuration(d time.Duration) string { _ = "STUB: not implemented"; return "" }

func isPKIXNameEmpty(name pkix.Name) bool {
	_ = "STUB: not implemented"
	// pkix.Name contains slices which make it directly incomparable. We could
	// do a field by field check since it is unlikely that pkix.Name will grow,
	// but reflect.DeepEqual is more convenient and safe for this particular
	// use.
	return false
}
