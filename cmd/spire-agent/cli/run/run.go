package run

import (
	"context"
	"io"
	"time"

	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/hashicorp/hcl/hcl/token"
	"github.com/mitchellh/cli"
	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire/pkg/agent"
	common_cli "github.com/spiffe/spire/pkg/common/cli"
	"github.com/spiffe/spire/pkg/common/fflag"
	"github.com/spiffe/spire/pkg/common/health"
	"github.com/spiffe/spire/pkg/common/log"
	"github.com/spiffe/spire/pkg/common/telemetry"
)

const (
	commandName = "run"

	defaultConfigPath = "conf/agent/agent.conf"

	// TODO: Make my defaults sane
	defaultDataDir                     = "."
	defaultLogLevel                    = "INFO"
	defaultDefaultSVIDName             = "default"
	defaultDefaultBundleName           = "ROOTCA"
	defaultDefaultAllBundlesName       = "ALL"
	defaultDisableSPIFFECertValidation = false

	minimumAvailabilityTarget = 24 * time.Hour
)

// Config contains all available configurables, arranged by section
type Config struct {
	Agent              *agentConfig           `hcl:"agent"`
	Plugins            ast.Node               `hcl:"plugins"`
	Telemetry          telemetry.FileConfig   `hcl:"telemetry"`
	HealthChecks       health.Config          `hcl:"health_checks"`
	UnusedKeyPositions map[string][]token.Pos `hcl:",unusedKeyPositions"`
}

type agentConfig struct {
	DataDir                       string    `hcl:"data_dir"`
	AdminSocketPath               string    `hcl:"admin_socket_path"`
	InsecureBootstrap             bool      `hcl:"insecure_bootstrap"`
	RebootstrapMode               string    `hcl:"rebootstrap_mode"`
	RebootstrapDelay              string    `hcl:"rebootstrap_delay"`
	JoinToken                     string    `hcl:"join_token"`
	JoinTokenFile                 string    `hcl:"join_token_file"`
	LogFile                       string    `hcl:"log_file"`
	LogFormat                     string    `hcl:"log_format"`
	LogLevel                      string    `hcl:"log_level"`
	LogSourceLocation             bool      `hcl:"log_source_location"`
	SDS                           sdsConfig `hcl:"sds"`
	ServerAddress                 string    `hcl:"server_address"`
	ServerPort                    int       `hcl:"server_port"`
	SocketPath                    string    `hcl:"socket_path"`
	WorkloadX509SVIDKeyType       string    `hcl:"workload_x509_svid_key_type"`
	TrustBundleFormat             string    `hcl:"trust_bundle_format"`
	TrustBundlePath               string    `hcl:"trust_bundle_path"`
	TrustBundleUnixSocket         string    `hcl:"trust_bundle_unix_socket"`
	TrustBundleURL                string    `hcl:"trust_bundle_url"`
	TrustDomain                   string    `hcl:"trust_domain"`
	AllowUnauthenticatedVerifiers bool      `hcl:"allow_unauthenticated_verifiers"`
	AllowedForeignJWTClaims       []string  `hcl:"allowed_foreign_jwt_claims"`
	AvailabilityTarget            string    `hcl:"availability_target"`
	X509SVIDCacheMaxSize          int       `hcl:"x509_svid_cache_max_size"`
	JWTSVIDCacheMaxSize           int       `hcl:"jwt_svid_cache_max_size"`

	AuthorizedDelegates []string `hcl:"authorized_delegates"`

	ConfigPath string
	ExpandEnv  bool

	// Undocumented configurables
	ProfilingEnabled bool               `hcl:"profiling_enabled"`
	ProfilingPort    int                `hcl:"profiling_port"`
	ProfilingFreq    int                `hcl:"profiling_freq"`
	ProfilingNames   []string           `hcl:"profiling_names"`
	Experimental     experimentalConfig `hcl:"experimental"`

	UnusedKeyPositions map[string][]token.Pos `hcl:",unusedKeyPositions"`
}

type sdsConfig struct {
	DefaultSVIDName             string `hcl:"default_svid_name"`
	DefaultBundleName           string `hcl:"default_bundle_name"`
	DefaultAllBundlesName       string `hcl:"default_all_bundles_name"`
	DisableSPIFFECertValidation bool   `hcl:"disable_spiffe_cert_validation"`
}

type experimentalConfig struct {
	SyncInterval             string `hcl:"sync_interval"`
	JWTSVIDCacheHitTimeout   string `hcl:"jwt_svid_cache_hit_timeout"`
	NamedPipeName            string `hcl:"named_pipe_name"`
	AdminNamedPipeName       string `hcl:"admin_named_pipe_name"`
	UseSyncAuthorizedEntries *bool  `hcl:"use_sync_authorized_entries"`
	RequirePQKEM             bool   `hcl:"require_pq_kem"`

	Flags fflag.RawConfig `hcl:"feature_flags"`
}

type Command struct {
	ctx                context.Context
	logOptions         []log.Option
	env                *common_cli.Env
	allowUnknownConfig bool
}

func NewRunCommand(ctx context.Context, logOptions []log.Option, allowUnknownConfig bool) cli.Command {
	_ = "STUB: not implemented"
	return *new(cli.Command)
}

func newRunCommand(ctx context.Context, env *common_cli.Env, logOptions []log.Option, allowUnknownConfig bool) *Command {
	_ = "STUB: not implemented"
	return nil
}

// Help prints the agent cmd usage
func (cmd *Command) Help() string { _ = "STUB: not implemented"; return "" }

// Help is a standalone function that prints a help message to writer.
// It is used by both the run and validate commands, so they can share flag usage messages.
func Help(name string, writer io.Writer) string { _ = "STUB: not implemented"; return "" }

// Error is always present because -h is passed

func LoadConfig(name string, args []string, logOptions []log.Option, output io.Writer, allowUnknownConfig bool) (*agent.Config, error) {
	_ = "STUB: not implemented"
	// First parse the CLI flags so we can get the config
	// file path, if set
	return nil, nil
}

// Load and parse the config file using either the default
// path or CLI-specified value

func (cmd *Command) Run(args []string) int { _ = "STUB: not implemented"; return 0 }

func (*Command) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *agentConfig) validate() error { _ = "STUB: not implemented"; return nil }

// Validate join token configuration

// If insecure_bootstrap is set, trust_bundle_path or trust_bundle_url cannot be set
// If trust_bundle_url is set, download the trust bundle using HTTP and parse it from memory
// If trust_bundle_path is set, parse the trust bundle file on disk
// Both cannot be set
// The trust bundle URL must start with HTTPS

func ParseFile(path string, expandEnv bool) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Return a friendly error if the file is missing

// If envTemplate flag is passed, substitute $VARIABLES in configuration file

func parseFlags(name string, args []string, output io.Writer) (*agentConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mergeInput(fileInput *Config, cliInput *agentConfig) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Highest precedence first

func NewAgentConfig(c *Config, logOptions []log.Option, allowUnknownConfig bool) (*agent.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Handle join token - read from file if specified

func validateConfig(c *Config) error { _ = "STUB: not implemented"; return nil }

func checkForUnknownConfig(c *Config, l logrus.FieldLogger) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Re-enable unused key detection for telemetry. See
// https://github.com/spiffe/spire/issues/1101 for more information
//
// if len(c.Telemetry.UnusedKeyPositions) != 0 {
//	detectedUnknown("telemetry", c.Telemetry.UnusedKeyPositions)
// }

func defaultConfig() *Config { _ = "STUB: not implemented"; return nil }
