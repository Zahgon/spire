//go:build !windows

package main

import (
	"context"
	"sync"

	"github.com/spiffe/spire-plugin-sdk/pluginmain"
	svidstorev1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/agent/svidstore/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
)

type Config struct {
	SVIDsPath string `hcl:"svids_path"`
}

type Plugin struct {
	svidstorev1.UnimplementedSVIDStoreServer
	configv1.UnimplementedConfigServer

	config *Config
	mtx    sync.RWMutex
	svids  map[string]*svidstorev1.X509SVID
}

func (p *Plugin) DeleteX509SVID(_ context.Context, req *svidstorev1.DeleteX509SVIDRequest) (*svidstorev1.DeleteX509SVIDResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) PutX509SVID(_ context.Context, req *svidstorev1.PutX509SVIDRequest) (*svidstorev1.PutX509SVIDResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) Configure(_ context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint // file used for testing

func (p *Plugin) putSVID(secretName string, svid *svidstorev1.X509SVID) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) deleteSVID(secretName string) error { _ = "STUB: not implemented"; return nil }

func (p *Plugin) updateFile(op func(map[string]*svidstorev1.X509SVID)) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // G117: complaint about marshaling PrivateKey field

func getSecretName(metadata []string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func main() {
	plugin := new(Plugin)

	pluginmain.Serve(
		svidstorev1.SVIDStorePluginServer(plugin),
		configv1.ConfigServiceServer(plugin),
	)
}
