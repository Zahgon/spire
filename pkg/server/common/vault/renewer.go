package vault

import (
	"github.com/hashicorp/go-hclog"
	vapi "github.com/hashicorp/vault/api"
)

const (
	defaultRenewBehavior = vapi.RenewBehaviorIgnoreErrors
)

type Renew struct {
	logger  hclog.Logger
	watcher *vapi.LifetimeWatcher
}

func NewRenew(client *vapi.Client, secret *vapi.Secret, logger hclog.Logger) (*Renew, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Renew) Run() { _ = "STUB: not implemented"; return }
