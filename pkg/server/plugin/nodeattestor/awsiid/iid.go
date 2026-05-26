package awsiid

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"regexp"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	iamtypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/fullsailor/pkcs7"
	"github.com/hashicorp/go-hclog"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	nodeattestorv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/nodeattestor/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/agentpathtemplate"
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/common/pluginconf"
	nodeattestorbase "github.com/spiffe/spire/pkg/server/plugin/nodeattestor/base"
)

var (
	awsTimeout      = 20 * time.Second
	instanceFilters = []ec2types.Filter{
		{
			Name: aws.String("instance-state-name"),
			Values: []string{
				"pending",
				"running",
			},
		},
	}

	defaultPartition = "aws"
	// No constant was found in the sdk, using the list of partitions defined on
	// the page https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html
	partitions = []string{
		defaultPartition,
		"aws-cn",
		"aws-us-gov",
	}
)

const (
	maxSecondsBetweenDeviceAttachments int64 = 60
	// accessKeyIDVarName env var name for AWS access key ID
	accessKeyIDVarName = "AWS_ACCESS_KEY_ID"
	// secretAccessKeyVarName env car name for AWS secret access key
	secretAccessKeyVarName   = "AWS_SECRET_ACCESS_KEY" //nolint: gosec // false positive
	accountIDSelectorPrefix  = "account_id"
	azSelectorPrefix         = "az"
	imageIDSelectorPrefix    = "image:id"
	instanceIDSelectorPrefix = "instance:id"
	regionSelectorPrefix     = "region"
	sgIDSelectorPrefix       = "sg:id"
	sgNameSelectorPrefix     = "sg:name"
	tagSelectorPrefix        = "tag"
	iamRoleSelectorPrefix    = "iamrole"
)

// BuiltIn creates a new built-in plugin
func BuiltIn() catalog.BuiltIn { _ = "STUB: not implemented"; return *new(catalog.BuiltIn) }

func builtin(p *IIDAttestorPlugin) catalog.BuiltIn {
	_ = "STUB: not implemented"
	return *new(catalog.BuiltIn)
}

// IIDAttestorPlugin implements node attestation for agents running in aws.
type IIDAttestorPlugin struct {
	nodeattestorbase.Base
	nodeattestorv1.UnsafeNodeAttestorServer
	configv1.UnsafeConfigServer

	config  *IIDAttestorConfig
	mtx     sync.RWMutex
	clients *clientsCache

	orgValidation *orgValidator
	eksValidation *eksValidator

	// test hooks
	hooks struct {
		getAWSCACertificate func(string, PublicKeyType) (*x509.Certificate, error)
		getenv              func(string) string
	}

	log hclog.Logger
}

// IIDAttestorConfig holds hcl configuration for IID attestor plugin
type IIDAttestorConfig struct {
	SessionConfig                   `hcl:",squash"`
	SkipBlockDevice                 bool                 `hcl:"skip_block_device"`
	DisableInstanceProfileSelectors bool                 `hcl:"disable_instance_profile_selectors"`
	LocalValidAcctIDs               []string             `hcl:"account_ids_for_local_validation"`
	AgentPathTemplate               string               `hcl:"agent_path_template"`
	AssumeRole                      string               `hcl:"assume_role"`
	Partition                       string               `hcl:"partition"`
	ValidateOrgAccountID            *orgValidationConfig `hcl:"verify_organization"`
	ValidateEKSClusterMembership    *eksValidationConfig `hcl:"validate_eks_cluster_membership"`
	pathTemplate                    *agentpathtemplate.Template
	trustDomain                     spiffeid.TrustDomain
	getAWSCACertificate             func(string, PublicKeyType) (*x509.Certificate, error)
}

func (p *IIDAttestorPlugin) buildConfig(coreConfig catalog.CoreConfig, hclText string, status *pluginconf.Status) *IIDAttestorConfig {
	_ = "STUB: not implemented"
	return nil
}

// Function to get the AWS CA certificate. We do this lazily on configure so deployments
// not using this plugin don't pay for parsing it on startup. This
// operation should not fail, but we check the return value just in case.

// Check if Feature flag for account belongs to organization is enabled.

// New creates a new IIDAttestorPlugin.
func New() *IIDAttestorPlugin { _ = "STUB: not implemented"; return nil }

// Attest implements the server side logic for the aws iid node attestation plugin.
func (p *IIDAttestorPlugin) Attest(stream nodeattestorv1.NodeAttestor_AttestServer) error {
	_ = "STUB: not implemented"
	return nil
}

// Feature account belongs to organization
// Get the account id of the node from attestation and then check if respective account belongs to organization

// Feature node belongs to EKS cluster

// Ideally we wouldn't do this work at all if the agent has already attested
// e.g. do it after the call to `p.AssessTOFU`, however, we may need
// the instance to construct tags used in the agent ID.
//
// This overhead will only affect agents attempting to re-attest which
// should be a very small portion of the overall server workload. This
// is a potential DoS vector.

// Configure configures the IIDAttestorPlugin.
func (p *IIDAttestorPlugin) Configure(_ context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// unconfigure existing clients

// Setup required config, for validation and for bootstrapping org client

func (p *IIDAttestorPlugin) Validate(_ context.Context, req *configv1.ValidateRequest) (*configv1.ValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetLogger sets this plugin's logger
func (p *IIDAttestorPlugin) SetLogger(log hclog.Logger) { _ = "STUB: not implemented"; return }

func (p *IIDAttestorPlugin) checkBlockDevice(instance ec2types.Instance) error {
	_ = "STUB: not implemented"
	return nil
}

// skip anti-tampering mechanism when RootDeviceType is instance-store
// specifically, if device type is persistent, and the device was attached past
// a threshold time after instance boot, fail attestation

func (p *IIDAttestorPlugin) getConfig() (*IIDAttestorConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *IIDAttestorPlugin) getEC2Instance(instancesDesc *ec2.DescribeInstancesOutput) (ec2types.Instance, error) {
	_ = "STUB: not implemented"
	return *new(ec2types.Instance), nil
}

func tagsFromInstance(instance ec2types.Instance) instanceTags {
	_ = "STUB: not implemented"
	return *new(instanceTags)
}

func unmarshalAndValidateIdentityDocument(data []byte, getAWSCACertificate func(string, PublicKeyType) (*x509.Certificate, error)) (imds.InstanceIdentityDocument, error) {
	_ = "STUB: not implemented"
	return *new(imds.InstanceIdentityDocument), nil
}

// Use the RSA-2048 signature if present, otherwise use the RSA-1024 signature
// This enables the support of new and old SPIRE agents, maintaining backwards compatibility.

// Verify the PKCS7 content matches the Document field
// to prevent substitution attacks where an attacker
// provides a legitimate PKCS7 signature from their
// own instance alongside a forged identity document.

func verifyRSASignature(pubKey *rsa.PublicKey, doc string, signature string) error {
	_ = "STUB: not implemented"
	return nil
}

func decodeAndParsePKCS7Signature(signature string, caCert *x509.Certificate) (*pkcs7.PKCS7, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// add the CA certificate to the PKCS7 signature to verify it

// AWS returns the PKCS7 signature without the header and footer. This function adds them to be able to parse
// the signature as a PEM block.
func addPKCS7HeaderAndFooter(signature string) string { _ = "STUB: not implemented"; return "" }

func (p *IIDAttestorPlugin) resolveSelectors(parent context.Context, instancesDesc *ec2.DescribeInstancesOutput, iiDoc imds.InstanceIdentityDocument, client Client) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// build and sort selectors

func resolveIIDocSelectors(selectorSet map[string]bool, iiDoc imds.InstanceIdentityDocument) {
	_ = "STUB: not implemented"
	return
}

func resolveTags(tags []ec2types.Tag) []string { _ = "STUB: not implemented"; return nil }

func resolveSecurityGroups(sgs []ec2types.GroupIdentifier) []string {
	_ = "STUB: not implemented"
	return nil
}

func resolveInstanceProfile(instanceProfile *iamtypes.InstanceProfile) []string {
	_ = "STUB: not implemented"
	return nil
}

var reInstanceProfileARNResource = regexp.MustCompile(`instance-profile[/:](.+)`)

func instanceProfileNameFromArn(profileArn string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// only the last element is the profile name

func isValidAWSPartition(partition string) bool { _ = "STUB: not implemented"; return false }

func validateOrganizationConfig(config *IIDAttestorConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// check TTL if specified

// Assign default ttl if ttl doesnt exist.
