package notifier

type Repository struct {
	Notifiers []Notifier
}

func (repo *Repository) GetNotifiers() []Notifier { _ = "STUB: not implemented"; return nil }

func (repo *Repository) AddNotifier(notifier Notifier) { _ = "STUB: not implemented"; return }

func (repo *Repository) Clear() { _ = "STUB: not implemented"; return }
