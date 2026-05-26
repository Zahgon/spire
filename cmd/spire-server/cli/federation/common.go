package federation

import (
	"encoding/json"
	"flag"

	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
)

// FederationRelationships type is used for parsing federation relationships from file
type federationRelationships struct {
	FederationRelationships []*federationRelationshipConfig `json:"federationRelationships"`
}

// federationRelationshipConfig is the configuration for the federation relationship provided either by CLI flags or a JSON file.
type federationRelationshipConfig struct {
	TrustDomain             string          `json:"trustDomain,omitempty"`
	BundleEndpointURL       string          `json:"bundleEndpointURL,omitempty"`
	BundleEndpointProfile   string          `json:"bundleEndpointProfile,omitempty"`
	EndpointSPIFFEID        string          `json:"endpointSPIFFEID,omitempty"`
	TrustDomainBundle       json.RawMessage `json:"trustDomainBundle,omitempty"`
	TrustDomainBundleFormat string          `json:"trustDomainBundleFormat,omitempty"`
	// TrustDomainBundlePath is the path to the bundle on disk. It is only set via CLI flags. JSON config uses the embedded `Bundle` field instead.
	TrustDomainBundlePath string `json:"-"`
}

func (c federationRelationshipConfig) isEmpty() bool { _ = "STUB: not implemented"; return false }

// federationRelationshipsFromFile parse a json file into types FederationRelationships
func federationRelationshipsFromFile(path string) ([]*types.FederationRelationship, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func jsonToProto(fr *federationRelationshipConfig) (*types.FederationRelationship, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// bundleFromPath get a bundle from a file
func bundleFromPath(bundlePath string, bundleFormat string, endpointTrustDomain string) (*types.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// bundleFromRawMessage get a bundle for a raw message
func bundleFromRawMessage(raw json.RawMessage, bundleFormat string, endpointTrustDomain string) (*types.Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func printFederationRelationship(fr *types.FederationRelationship, printf func(format string, args ...any) error) {
	_ = "STUB: not implemented"
	return
}

func appendConfigFlags(config *federationRelationshipConfig, f *flag.FlagSet) {
	_ = "STUB: not implemented"
	return
}

func getRelationships(config *federationRelationshipConfig, path string) ([]*types.FederationRelationship, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
