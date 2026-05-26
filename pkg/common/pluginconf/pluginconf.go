package pluginconf

import (
	"github.com/hashicorp/hcl/hcl/token"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"github.com/spiffe/spire/pkg/common/catalog"
)

// ReportUnusedKeys reports an error on s listing any keys present in
// unused. If unused is empty, no error is reported.
func ReportUnusedKeys(s *Status, unused map[string][]token.Pos) { _ = "STUB: not implemented"; return }

type Status struct {
	notes []string
	err   error
}

func (s *Status) ReportInfo(message string) { _ = "STUB: not implemented"; return }

func (s *Status) ReportInfof(format string, args ...any) { _ = "STUB: not implemented"; return }

func (s *Status) ReportError(message string) { _ = "STUB: not implemented"; return }

func (s *Status) ReportErrorf(format string, args ...any) { _ = "STUB: not implemented"; return }

type Request interface {
	GetCoreConfiguration() *configv1.CoreConfiguration
	GetHclConfiguration() string
}

func Build[C any](req Request, build func(coreConfig catalog.CoreConfig, hclText string, s *Status) *C) (*C, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
