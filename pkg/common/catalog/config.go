package catalog

import (
	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/hashicorp/hcl/hcl/token"
)

type PluginConfigs []PluginConfig

func (cs PluginConfigs) FilterByType(pluginType string) (matching PluginConfigs, remaining PluginConfigs) {
	_ = "STUB: not implemented"
	return *new(PluginConfigs), *new(PluginConfigs)
}

func (cs PluginConfigs) Find(pluginType, pluginName string) (PluginConfig, bool) {
	_ = "STUB: not implemented"
	return *new(PluginConfig), false
}

type PluginConfig struct {
	Type       string
	Name       string
	Path       string
	Args       []string
	Checksum   string
	DataSource DataSource
	Disabled   bool
}

func (c PluginConfig) IsEnabled() bool { _ = "STUB: not implemented"; return false }

func (c *PluginConfig) IsExternal() bool { _ = "STUB: not implemented"; return false }

type DataSource interface {
	Load() (string, error)
	IsDynamic() bool
}

type FixedData string

func (d FixedData) Load() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (d FixedData) IsDynamic() bool { _ = "STUB: not implemented"; return false }

type FileData string

func (d FileData) Load() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (d FileData) IsDynamic() bool { _ = "STUB: not implemented"; return false }

type hclPluginConfig struct {
	PluginCmd      string   `hcl:"plugin_cmd"`
	PluginArgs     []string `hcl:"plugin_args"`
	PluginChecksum string   `hcl:"plugin_checksum"`
	PluginData     ast.Node `hcl:"plugin_data"`
	PluginDataFile *string  `hcl:"plugin_data_file"`
	Enabled        *bool    `hcl:"enabled"`
}

func (c hclPluginConfig) IsEnabled() bool { _ = "STUB: not implemented"; return false }

func (c hclPluginConfig) IsExternal() bool { _ = "STUB: not implemented"; return false }

func PluginConfigsFromHCLNode(pluginsNode ast.Node) (PluginConfigs, error) {
	_ = "STUB: not implemented"
	return *new(PluginConfigs), nil
}

// Sanity check the length of the pluginsMapList and those found when
// determining order. If this mismatches, it's a bug.

// This would be a programmer error. We should always be able to
// locate the plugin configuration in one of the maps.

type pluginIdent struct {
	Type string
	Name string
}

func determinePluginOrder(pluginsList *ast.ObjectList) ([]pluginIdent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Walk the AST, pushing and popping nodes from an "object" stack. At
// each step, determine if we've accumulated object keys at least 2 deep.
// If so, we've found a plugin definition and add the plugin identifier
// to the ordering.
//
// This accommodates nesting of all shapes and sizes, for example:
//
// "NodeAttestor" {
//     "k8s_psat" {
//         plugin_data {
//         }
//     }
// }
//
// "NodeAttestor" "k8s_psat" {
//     plugin_data {
//     }
// }
//
// "NodeAttestor" "k8s_psat" plugin_data {
// }
//
//

// Since we've found an object item for the plugin, pop it from
// the stack and do not recurse.

// Check for duplicates

type pluginsMapList []map[string]map[string]hclPluginConfig

func (m pluginsMapList) FindPluginConfig(pluginType, pluginName string) (hclPluginConfig, bool) {
	_ = "STUB: not implemented"
	return *new(hclPluginConfig), false
}

func (m pluginsMapList) Len() int { _ = "STUB: not implemented"; return 0 }

func pluginConfigFromHCL(pluginType, pluginName string, hclPluginConfig hclPluginConfig) (PluginConfig, error) {
	_ = "STUB: not implemented"
	return *new(PluginConfig), nil
}

func stringFromToken(keyToken token.Token) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
