package keymanager

type Repository struct {
	KeyManager KeyManager
}

func (repo *Repository) GetKeyManager() KeyManager {
	_ = "STUB: not implemented"
	return *new(KeyManager)
}

func (repo *Repository) SetKeyManager(keyManager KeyManager) { _ = "STUB: not implemented"; return }

func (repo *Repository) Clear() { _ = "STUB: not implemented"; return }
