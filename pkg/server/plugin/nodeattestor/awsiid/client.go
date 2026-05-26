package awsiid

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
)

var (
	defaultNewClientCallback = newClient
)

type Client interface {
	ec2.DescribeInstancesAPIClient
	iam.GetInstanceProfileAPIClient
	organizations.ListAccountsAPIClient
	autoscaling.DescribeAutoScalingGroupsAPIClient
	eks.ListNodegroupsAPIClient
	eks.DescribeNodegroupAPIClient
}

type clientsCache struct {
	mtx       sync.RWMutex
	config    *SessionConfig
	orgConfig *orgValidationConfig
	clients   map[string]*cacheEntry
	newClient newClientCallback
}

type cacheEntry struct {
	lock   chan struct{}
	client Client
}

type newClientCallback func(ctx context.Context, config *SessionConfig, region string, assumeRoleARN string, orgRoleARN string) (Client, error)

func newClientsCache(newClient newClientCallback) *clientsCache {
	_ = "STUB: not implemented"
	return nil
}

func (cc *clientsCache) configure(config SessionConfig, orgConfig orgValidationConfig) {
	_ = "STUB: not implemented"
	return
}

func (cc *clientsCache) getClient(ctx context.Context, region, accountID string) (Client, error) {
	_ = "STUB: not implemented"
	// Do an initial check to see if p client for this region already exists
	return *new(Client), nil
}

// Grab (or create) the cache for the region

// Obtain the "lock" to the region cache

// "clear" the lock when the function is complete

// If the client is populated, return it.

// If organization attestation feature is enabled, assume org role

func (cc *clientsCache) getCachedClient(cacheKey string) *cacheEntry {
	_ = "STUB: not implemented"
	return nil
}

func newClient(ctx context.Context, config *SessionConfig, region string, assumeRoleARN string, orgRoleArn string) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

// If the organizationAttestation feature is enabled, use the role configured for feature.
