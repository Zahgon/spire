package awskms

import (
	"context"
	"crypto"
	"regexp"
	"sync"
	"testing"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/spiffe/spire/test/testkey"
)

type kmsClientFake struct {
	t                      *testing.T
	store                  fakeStore
	mu                     sync.RWMutex
	testKeys               testkey.Keys
	validAliasName         *regexp.Regexp
	createKeyErr           error
	describeKeyErr         error
	describeKeyMalformed   bool
	getPublicKeyErr        error
	listAliasesErr         error
	createAliasErr         error
	updateAliasErr         error
	scheduleKeyDeletionErr error
	signErr                error
	listKeysErr            error
	deleteAliasErr         error

	expectedKeyPolicy *string
}

type stsClientFake struct {
	account string
	arn     string
	err     string
}

func newKMSClientFake(t *testing.T, c *clock.Mock) *kmsClientFake {
	_ = "STUB: not implemented"
	return nil
}

// Valid KMS alias name must match the expression below:
// https://docs.aws.amazon.com/kms/latest/APIReference/API_CreateAlias.html#API_CreateAlias_RequestSyntax

func newSTSClientFake() *stsClientFake { _ = "STUB: not implemented"; return nil }

func (s *stsClientFake) GetCallerIdentity(context.Context, *sts.GetCallerIdentityInput, ...func(*sts.Options)) (*sts.GetCallerIdentityOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *stsClientFake) setGetCallerIdentityErr(err string) { _ = "STUB: not implemented"; return }

func (s *stsClientFake) setGetCallerIdentityAccount(account string) {
	_ = "STUB: not implemented"
	return
}

func (s *stsClientFake) setGetCallerIdentityArn(arn string) { _ = "STUB: not implemented"; return }

func (k *kmsClientFake) setExpectedKeyPolicy(keyPolicy *string) { _ = "STUB: not implemented"; return }

func (k *kmsClientFake) CreateKey(_ context.Context, input *kms.CreateKeyInput, _ ...func(*kms.Options)) (*kms.CreateKeyOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *kmsClientFake) DescribeKey(_ context.Context, input *kms.DescribeKeyInput, _ ...func(*kms.Options)) (*kms.DescribeKeyOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *kmsClientFake) GetPublicKey(_ context.Context, input *kms.GetPublicKeyInput, _ ...func(*kms.Options)) (*kms.GetPublicKeyOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *kmsClientFake) ListAliases(_ context.Context, input *kms.ListAliasesInput, _ ...func(*kms.Options)) (*kms.ListAliasesOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *kmsClientFake) ScheduleKeyDeletion(_ context.Context, input *kms.ScheduleKeyDeletionInput, _ ...func(*kms.Options)) (*kms.ScheduleKeyDeletionOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *kmsClientFake) Sign(_ context.Context, input *kms.SignInput, _ ...func(*kms.Options)) (*kms.SignOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *kmsClientFake) CreateAlias(_ context.Context, input *kms.CreateAliasInput, _ ...func(*kms.Options)) (*kms.CreateAliasOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *kmsClientFake) UpdateAlias(_ context.Context, input *kms.UpdateAliasInput, _ ...func(*kms.Options)) (*kms.UpdateAliasOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *kmsClientFake) ListKeys(context.Context, *kms.ListKeysInput, ...func(*kms.Options)) (*kms.ListKeysOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *kmsClientFake) DeleteAlias(_ context.Context, params *kms.DeleteAliasInput, _ ...func(*kms.Options)) (*kms.DeleteAliasOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *kmsClientFake) setEntries(entries []fakeKeyEntry) { _ = "STUB: not implemented"; return }

func (k *kmsClientFake) setCreateKeyErr(fakeError string) { _ = "STUB: not implemented"; return }

func (k *kmsClientFake) setDescribeKeyErr(fakeError string) { _ = "STUB: not implemented"; return }

func (k *kmsClientFake) setDescribeKeyMalformed(malformed bool) { _ = "STUB: not implemented"; return }

func (k *kmsClientFake) setgetPublicKeyErr(fakeError string) { _ = "STUB: not implemented"; return }

func (k *kmsClientFake) setListAliasesErr(fakeError string) { _ = "STUB: not implemented"; return }

func (k *kmsClientFake) setCreateAliasesErr(fakeError string) { _ = "STUB: not implemented"; return }

func (k *kmsClientFake) setUpdateAliasErr(fakeError string) { _ = "STUB: not implemented"; return }

func (k *kmsClientFake) setScheduleKeyDeletionErr(fakeError error) {
	_ = "STUB: not implemented"
	return
}

func (k *kmsClientFake) setSignDataErr(fakeError string) { _ = "STUB: not implemented"; return }

func (k *kmsClientFake) setListKeysErr(fakeError string) { _ = "STUB: not implemented"; return }

func (k *kmsClientFake) setDeleteAliasErr(fakeError string) { _ = "STUB: not implemented"; return }

const (
	fakeKeyArnPrefix   = "arn:aws:kms:region:1234:key/"
	fakeAliasArnPrefix = "arn:aws:kms:region:1234:"
)

type fakeStore struct {
	keyEntries map[string]*fakeKeyEntry // don't user ara for key
	aliases    map[string]fakeAlias     // don't user ara for key
	mu         sync.RWMutex
	nextID     int
	clk        *clock.Mock
}

func newFakeStore(c *clock.Mock) fakeStore { _ = "STUB: not implemented"; return *new(fakeStore) }

type fakeKeyEntry struct {
	KeyID                *string
	Arn                  *string
	Description          *string
	CreationDate         *time.Time
	AliasName            *string // Only one alias per key. "Real" KMS supports many aliases per key
	AliasLastUpdatedDate *time.Time
	PublicKey            []byte
	privateKey           crypto.Signer
	Enabled              bool
	KeySpec              types.KeySpec
}

type fakeAlias struct {
	AliasName *string
	AliasArn  *string
	KeyEntry  *fakeKeyEntry
}

func (fs *fakeStore) SaveKeyEntry(input *fakeKeyEntry) { _ = "STUB: not implemented"; return }

func (fs *fakeStore) DeleteKeyEntry(keyID string) { _ = "STUB: not implemented"; return }

func (fs *fakeStore) SaveAlias(targetKeyID, aliasName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (fs *fakeStore) DeleteAlias(aliasName string) { _ = "STUB: not implemented"; return }

func (fs *fakeStore) ListKeyEntries() []fakeKeyEntry { _ = "STUB: not implemented"; return nil }

func (fs *fakeStore) ListAliases() []fakeAlias { _ = "STUB: not implemented"; return nil }

func (fs *fakeStore) FetchKeyEntry(id string) (*fakeKeyEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fs *fakeStore) fetchKeyEntry(id string) (*fakeKeyEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func aliasArnFromAliasName(aliasName string) string { _ = "STUB: not implemented"; return "" }

func aliasNameFromArn(arn string) string { _ = "STUB: not implemented"; return "" }

func arnFromKeyID(keyID string) string { _ = "STUB: not implemented"; return "" }

func keyIDFromArn(arn string) string { _ = "STUB: not implemented"; return "" }
