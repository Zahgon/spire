package bundlepublisher

type Repository struct {
	BundlePublishers []BundlePublisher
}

func (repo *Repository) GetBundlePublishers() []BundlePublisher {
	_ = "STUB: not implemented"
	return nil
}

func (repo *Repository) AddBundlePublisher(bundlePublisher BundlePublisher) {
	_ = "STUB: not implemented"
	return
}

func (repo *Repository) Clear() { _ = "STUB: not implemented"; return }
