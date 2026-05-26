package upstreamauthority

type Repository struct {
	UpstreamAuthority UpstreamAuthority
}

func (repo *Repository) GetUpstreamAuthority() (UpstreamAuthority, bool) {
	_ = "STUB: not implemented"
	return *new(UpstreamAuthority), false
}

func (repo *Repository) SetUpstreamAuthority(upstreamAuthority UpstreamAuthority) {
	_ = "STUB: not implemented"
	return
}

func (repo *Repository) ClearUpstreamAuthority() { _ = "STUB: not implemented"; return }

func (repo *Repository) Clear() { _ = "STUB: not implemented"; return }
