package datastore

import (
	"github.com/spiffe/spire/pkg/common/health"
)

type Health struct {
	DataStore DataStore
}

func (h *Health) CheckHealth() health.State { _ = "STUB: not implemented"; return *new(health.State) }

// Both liveness and readiness are determined by the datastore's
// ability to list all the bundles.

type HealthDetails struct {
	ListBundleErr string `json:"list_bundle_err,omitempty"`
}

func errString(err error) string { _ = "STUB: not implemented"; return "" }
