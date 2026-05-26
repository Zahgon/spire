package credentialcomposer

type Repository struct {
	CredentialComposers []CredentialComposer
}

func (repo *Repository) GetCredentialComposers() []CredentialComposer {
	_ = "STUB: not implemented"
	return nil
}

func (repo *Repository) AddCredentialComposer(credentialComposer CredentialComposer) {
	_ = "STUB: not implemented"
	return
}

func (repo *Repository) Clear() { _ = "STUB: not implemented"; return }
