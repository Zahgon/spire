package awsiid

import (
	"context"
	"sync"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/hashicorp/go-hclog"
)

const (
	orgAccountID             = "management_account_id"
	orgAccountRole           = "assume_org_role"
	orgAccRegion             = "management_account_region" // required for cache key
	orgAccountStatus         = "ACTIVE"                    // Only allow node account id's with status ACTIVE
	orgAccountListTTL        = "org_account_map_ttl"       // Cache the list of account for specific time, if not sent default will be used.
	orgAccountDefaultListTTL = "3m"                        // pull account list after 3 minutes
	orgAccountMinListTTL     = "1m"                        // Minimum TTL configuration to pull the org account list
	orgAccountRetries        = 5
	orgDefaultAccRegion      = "us-west-2"
)

var (
	orgAccountDefaultListDuration, _ = time.ParseDuration(orgAccountDefaultListTTL)
	orgAccountMinTTL, _              = time.ParseDuration(orgAccountMinListTTL)
)

type orgValidationConfig struct {
	AccountID      string `hcl:"management_account_id"`
	AccountRole    string `hcl:"assume_org_role"`
	AccountRegion  string `hcl:"management_account_region"`
	AccountListTTL string `hcl:"org_account_map_ttl"`
}

type orgValidator struct {
	orgAccountList              map[string]any
	orgAccountListValidDuration time.Time
	orgConfig                   *orgValidationConfig
	mutex                       sync.RWMutex
	// orgAccountListCacheTTL holds the cache ttl from configuration; otherwise, it will be set to the default value.
	orgAccountListCacheTTL time.Duration
	log                    hclog.Logger
	// retries fix number of retries before ttl is expired.
	retries int
	// require for testing
	clk clock.Clock
}

func newOrganizationValidationBase(config *orgValidationConfig) *orgValidator {
	_ = "STUB: not implemented"
	return nil
}

func (o *orgValidator) getRetries() int { _ = "STUB: not implemented"; return 0 }

func (o *orgValidator) decrRetries() int { _ = "STUB: not implemented"; return 0 }

func (o *orgValidator) configure(config *orgValidationConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// While doing configuration invalidate the map so we don't keep using old one.

func (o *orgValidator) setLogger(log hclog.Logger) {
	_ = "STUB: not implemented"

	// IsMemberAccount method checks if the Account ID attached on the node is part of the organization.
	// If it is part of the organization then validation should be successful if not attestation should fail, on enabling this verification method.
	// This could be alternative for not explicitly maintaining allowed list of account ids.
	// Method pulls the list of accounts from the organization and caches it for certain time, cache time can be configured.
	return
}

func (o *orgValidator) IsMemberAccount(ctx context.Context, orgClient organizations.ListAccountsAPIClient, accountIDOfNode string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// validateCache validates cache and refresh if its stale
func (o *orgValidator) validateCache(ctx context.Context, orgClient organizations.ListAccountsAPIClient) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// cache is stale, reload the account map

func (o *orgValidator) lookupCache(ctx context.Context, orgClient organizations.ListAccountsAPIClient, accountIDOfNode string, reValidatedCache bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Retry if it doesn't exist in cache and cache was not revalidated

// refreshCache refreshes list with new cache if cache miss happens and check if element exist
func (o *orgValidator) refreshCache(ctx context.Context, orgClient organizations.ListAccountsAPIClient) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// checkIfOrgAccountListIsStale checks if the cached org account list is stale.
func (o *orgValidator) checkIfOrgAccountListIsStale() bool { _ = "STUB: not implemented"; return false }

// Map is empty that means this is first time plugin is being initialised

// reloadAccountList gets the list of accounts belonging to organization and catch them
func (o *orgValidator) reloadAccountList(ctx context.Context, orgClient organizations.ListAccountsAPIClient, catchBurst bool) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Make sure: we are not doing cache burst and account map is not updated recently from different go routine.

// Avoid if other thread has already updated the map

// Get the list of accounts

// Build new org accounts list

// Update the org account list cache with ACTIVE accounts & handle pagination

// Update timestamp, if it was not invoked as part of cache miss.

// Also reset the retries

// Overwrite the cache/list

// checkIFTTLIsExpire check if the creation time is pass defined ttl
func (o *orgValidator) checkIfTTLIsExpired(ttl time.Time) bool {
	_ = "STUB: not implemented"
	return false
}
