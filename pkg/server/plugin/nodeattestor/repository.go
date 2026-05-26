package nodeattestor

type Repository struct {
	NodeAttestors map[string]NodeAttestor
}

func (repo *Repository) GetNodeAttestorNamed(name string) (NodeAttestor, bool) {
	_ = "STUB: not implemented"
	return *new(NodeAttestor), false
}

func (repo *Repository) SetNodeAttestor(nodeAttestor NodeAttestor) {
	_ = "STUB: not implemented"
	return
}

func (repo *Repository) Clear() { _ = "STUB: not implemented"; return }
