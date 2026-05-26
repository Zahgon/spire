package workloadattestor

type Repository struct {
	WorkloadAttestors []WorkloadAttestor
}

func (repo *Repository) GetWorkloadAttestors() []WorkloadAttestor {
	_ = "STUB: not implemented"
	return nil
}

func (repo *Repository) AddWorkloadAttestor(workloadattestor WorkloadAttestor) {
	_ = "STUB: not implemented"
	return
}

func (repo *Repository) SetWorkloadAttestors(workloadAttestors ...WorkloadAttestor) {
	_ = "STUB: not implemented"
	return
}

func (repo *Repository) Clear() { _ = "STUB: not implemented"; return }
