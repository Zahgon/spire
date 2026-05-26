package catalog

import "context"

type pluginNameKey struct{}

func PluginNameFromHostServiceContext(ctx context.Context) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func WithPluginName(ctx context.Context, name string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
