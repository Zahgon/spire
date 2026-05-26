package nodeattestor

type Repository struct {
	NodeAttestor NodeAttestor
}

func (repo *Repository) GetNodeAttestor() NodeAttestor {
	_ = "STUB: not implemented"
	return *new(NodeAttestor)
}

func (repo *Repository) SetNodeAttestor(nodeAttestor NodeAttestor) {
	_ = "STUB: not implemented"
	return
}

func (repo *Repository) Clear() { _ = "STUB: not implemented"; return }
