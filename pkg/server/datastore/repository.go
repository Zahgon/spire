package datastore

type Repository struct {
	DataStore DataStore
}

func (repo *Repository) GetDataStore() DataStore { _ = "STUB: not implemented"; return *new(DataStore) }

func (repo *Repository) SetDataStore(dataStore DataStore) { _ = "STUB: not implemented"; return }
