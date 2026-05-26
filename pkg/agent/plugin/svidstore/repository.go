package svidstore

type Repository struct {
	SVIDStores map[string]SVIDStore
}

func (repo *Repository) GetSVIDStoreNamed(name string) (SVIDStore, bool) {
	_ = "STUB: not implemented"
	return *new(SVIDStore), false
}

func (repo *Repository) SetSVIDStore(svidStore SVIDStore) { _ = "STUB: not implemented"; return }

func (repo *Repository) Clear() { _ = "STUB: not implemented"; return }
