package awskms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/hashicorp/go-hclog"
)

type keyFetcher struct {
	log         hclog.Logger
	kmsClient   kmsClient
	serverID    string
	trustDomain string
}

func (kf *keyFetcher) fetchKeyEntries(ctx context.Context) ([]*keyEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ensure the alias has a name. This check is purely defensive
// since aliases should always have a name.

// ignore aliases/keys not belonging to this server

// The following checks are purely defensive, but we want to ensure
// we don't try and handle an alias with a malformed shape.

// this means something external to the plugin created the alias, without associating it to a key.
// it should never happen with CMKs.

// trigger a goroutine to get the details of the key

// wait for all the detail gathering routines to finish

func (kf *keyFetcher) fetchKeyEntryDetails(ctx context.Context, alias types.AliasListEntry, spireKeyID string) (*keyEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// this means something external to the plugin, deleted or disabled the key without removing the alias
// returning an error provides the opportunity or reverting this in KMS

func (kf *keyFetcher) spireKeyIDFromAlias(aliasName string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
