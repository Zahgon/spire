package client

import (
	"context"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/server/datastore"
)

type TrustDomainConfigSource interface {
	GetTrustDomainConfigs(ctx context.Context) (map[spiffeid.TrustDomain]TrustDomainConfig, error)
}

type TrustDomainConfigSourceFunc func(ctx context.Context) (map[spiffeid.TrustDomain]TrustDomainConfig, error)

func (fn TrustDomainConfigSourceFunc) GetTrustDomainConfigs(ctx context.Context) (map[spiffeid.TrustDomain]TrustDomainConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type TrustDomainConfigMap = map[spiffeid.TrustDomain]TrustDomainConfig

type TrustDomainConfigSet struct {
	mtx       sync.RWMutex
	configMap TrustDomainConfigMap
}

func NewTrustDomainConfigSet(configs TrustDomainConfigMap) *TrustDomainConfigSet {
	_ = "STUB: not implemented"
	return nil
}

func (s *TrustDomainConfigSet) Set(td spiffeid.TrustDomain, config TrustDomainConfig) {
	_ = "STUB: not implemented"
	return
}

func (s *TrustDomainConfigSet) SetAll(configMap TrustDomainConfigMap) {
	_ = "STUB: not implemented"
	return
}

func (s *TrustDomainConfigSet) GetTrustDomainConfigs(context.Context) (map[spiffeid.TrustDomain]TrustDomainConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func duplicateTrustDomainConfigMap(in TrustDomainConfigMap) TrustDomainConfigMap {
	_ = "STUB: not implemented"
	return *new(TrustDomainConfigMap)
}

func MergeTrustDomainConfigSources(sources ...TrustDomainConfigSource) TrustDomainConfigSource {
	_ = "STUB: not implemented"
	return *new(TrustDomainConfigSource)
}

// merge in reverse order

func DataStoreTrustDomainConfigSource(log logrus.FieldLogger, ds datastore.DataStore) TrustDomainConfigSource {
	_ = "STUB: not implemented"
	return *new(TrustDomainConfigSource)
}
