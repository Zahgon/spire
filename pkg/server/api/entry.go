package api

import (
	"context"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/proto/spire/common"
)

const (
	hintMaximumLength = 1024
)

type ReadOnlyEntry struct {
	entry *types.Entry
}

func NewReadOnlyEntry(entry *types.Entry) ReadOnlyEntry {
	_ = "STUB: not implemented"
	return *new(ReadOnlyEntry)
}

func (e ReadOnlyEntry) GetId() string { _ = "STUB: not implemented"; return "" }

func (e *ReadOnlyEntry) GetSpiffeId() *types.SPIFFEID { _ = "STUB: not implemented"; return nil }

func (e *ReadOnlyEntry) GetX509SvidTtl() int32 { _ = "STUB: not implemented"; return 0 }

func (e *ReadOnlyEntry) GetJwtSvidTtl() int32 { _ = "STUB: not implemented"; return 0 }

func (e *ReadOnlyEntry) GetDnsNames() []string { _ = "STUB: not implemented"; return nil }

func (e *ReadOnlyEntry) GetRevisionNumber() int64 { _ = "STUB: not implemented"; return 0 }

func (e *ReadOnlyEntry) GetCreatedAt() int64 { _ = "STUB: not implemented"; return 0 }

func (e *ReadOnlyEntry) GetAdditionalAttributes() *types.Entry_AdditionalAttributes {
	_ = "STUB: not implemented"
	return nil
}

// Manually clone the entry instead of using the protobuf helpers
// since those are two times slower.
func (e *ReadOnlyEntry) Clone(mask *types.EntryMask) *types.Entry {
	_ = "STUB: not implemented"
	return nil
}

// RegistrationEntriesToProto converts RegistrationEntry's into Entry's
func RegistrationEntriesToProto(es []*common.RegistrationEntry) ([]*types.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RegistrationEntryToProto converts RegistrationEntry into types Entry
func RegistrationEntryToProto(e *common.RegistrationEntry) (*types.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ProtoFromAdditionalAttributes(in *common.RegistrationEntry_AdditionalAttributes) *types.Entry_AdditionalAttributes {
	_ = "STUB: not implemented"
	return nil
}

func AdditionalAttributesFromProto(in *types.Entry_AdditionalAttributes) *common.RegistrationEntry_AdditionalAttributes {
	_ = "STUB: not implemented"
	return nil
}

// ProtoToRegistrationEntry converts and validate entry into common registration entry
func ProtoToRegistrationEntry(ctx context.Context, td spiffeid.TrustDomain, e *types.Entry) (*common.RegistrationEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProtoToRegistrationEntryWithMask converts and validate entry into common registration entry,
// while allowing empty values for SpiffeId, ParentId, and Selectors IF their corresponding values
// in the mask are false.
// This allows the user to not specify these fields while updating using a mask.
// All other fields are allowed to be empty (with or without a mask).
func ProtoToRegistrationEntryWithMask(ctx context.Context, td spiffeid.TrustDomain, e *types.Entry, mask *types.EntryMask) (_ *common.RegistrationEntry, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
