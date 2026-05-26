package catalog

import (
	"reflect"
)

type bindablePluginRepo interface {
	PluginRepo
	bindable
}

type bindableServiceRepo interface {
	ServiceRepo
	bindable
}

type bindable interface {
	bind(Facade)
}

func makeBindablePluginRepos(repos map[string]PluginRepo) (map[string]bindablePluginRepo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func makeBindablePluginRepo(repo PluginRepo) (bindablePluginRepo, error) {
	_ = "STUB: not implemented"
	return *new(bindablePluginRepo), nil
}

func makeBindableServiceRepos(repos []ServiceRepo) ([]bindableServiceRepo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func makeBindableServiceRepo(repo ServiceRepo) (bindableServiceRepo, error) {
	_ = "STUB: not implemented"
	return *new(bindableServiceRepo), nil
}

func makeServiceRepoBinder(repo ServiceRepo) (binder, error) {
	_ = "STUB: not implemented"
	return *new(binder), nil
}

type binder struct {
	fnv reflect.Value
}

func makeBinder(fn any) (binder, error) { _ = "STUB: not implemented"; return *new(binder), nil }

func (b binder) canBind(facade Facade) error { _ = "STUB: not implemented"; return nil }

func (b binder) bind(facade Facade) { _ = "STUB: not implemented"; return }
