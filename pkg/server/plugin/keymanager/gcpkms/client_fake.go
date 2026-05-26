package gcpkms

import (
	"context"
	"crypto"
	"fmt"
	"regexp"
	"sync"
	"testing"

	"cloud.google.com/go/iam"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/googleapis/gax-go/v2"
	"github.com/spiffe/spire/test/clock"
	"google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var lastUpdateFilterRegexp = regexp.MustCompile(fmt.Sprintf(`labels.%s < ([[:digit:]]+)`, labelNameLastUpdate))

type fakeCryptoKeyIterator struct {
	mu sync.RWMutex

	index      int
	cryptoKeys []*kmspb.CryptoKey
	nextErr    error
}

func (i *fakeCryptoKeyIterator) Next() (cryptoKey *kmspb.CryptoKey, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type fakeCryptoKeyVersionIterator struct {
	mu sync.RWMutex

	index             int
	cryptoKeyVersions []*kmspb.CryptoKeyVersion
	nextErr           error
}

func (i *fakeCryptoKeyVersionIterator) Next() (cryptoKeyVersion *kmspb.CryptoKeyVersion, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type fakeCryptoKey struct {
	mu sync.RWMutex
	*kmspb.CryptoKey
	fakeCryptoKeyVersions map[string]*fakeCryptoKeyVersion
}

func (fck *fakeCryptoKey) fetchFakeCryptoKeyVersions() map[string]*fakeCryptoKeyVersion {
	_ = "STUB: not implemented"
	return nil
}

func (fck *fakeCryptoKey) getLabelValue(key string) string { _ = "STUB: not implemented"; return "" }

func (fck *fakeCryptoKey) getName() string { _ = "STUB: not implemented"; return "" }

func (fck *fakeCryptoKey) putFakeCryptoKeyVersion(fckv *fakeCryptoKeyVersion) {
	_ = "STUB: not implemented"
	return
}

type fakeCryptoKeyVersion struct {
	*kmspb.CryptoKeyVersion

	privateKey crypto.Signer
	publicKey  *kmspb.PublicKey
}

type fakeStore struct {
	mu             sync.RWMutex
	fakeCryptoKeys map[string]*fakeCryptoKey

	clk *clock.Mock
}

func (fs *fakeStore) fetchFakeCryptoKey(name string) (*fakeCryptoKey, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (fs *fakeStore) fetchFakeCryptoKeys() map[string]*fakeCryptoKey {
	_ = "STUB: not implemented"
	return nil
}

func (fs *fakeStore) fetchFakeCryptoKeyVersion(name string) (fakeCryptoKeyVersion, error) {
	_ = "STUB: not implemented"
	return *new(fakeCryptoKeyVersion), nil
}

func (fs *fakeStore) putFakeCryptoKey(fck *fakeCryptoKey) { _ = "STUB: not implemented"; return }

type fakeIAMHandle struct {
	mu             sync.RWMutex
	expectedPolicy *iam.Policy3
	policyErr      error
	setPolicyErr   error
}

func (h *fakeIAMHandle) V3() iamHandler3 { _ = "STUB: not implemented"; return *new(iamHandler3) }

func (h *fakeIAMHandle) setExpectedPolicy(expectedPolicy *iam.Policy3) {
	_ = "STUB: not implemented"
	return
}

func (h *fakeIAMHandle) setPolicyError(fakeError error) { _ = "STUB: not implemented"; return }

func (h *fakeIAMHandle) setSetPolicyErr(fakeError error) { _ = "STUB: not implemented"; return }

type fakeIAMHandle3 struct {
	mu             sync.RWMutex
	expectedPolicy *iam.Policy3
	policyErr      error
	setPolicyErr   error
}

func (h3 *fakeIAMHandle3) Policy(context.Context) (*iam.Policy3, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h3 *fakeIAMHandle3) SetPolicy(_ context.Context, policy *iam.Policy3) error {
	_ = "STUB: not implemented"
	return nil
}

type fakeKMSClient struct {
	t *testing.T

	mu                           sync.RWMutex
	asymmetricSignErr            error
	closeErr                     error
	createCryptoKeyErr           error
	initialCryptoKeyVersionState kmspb.CryptoKeyVersion_CryptoKeyVersionState
	destroyCryptoKeyVersionErr   error
	destroyTime                  *timestamppb.Timestamp
	fakeIAMHandle                *fakeIAMHandle
	getCryptoKeyVersionErr       error
	getPublicKeyErrs             []error
	getTokeninfoErr              error
	listCryptoKeysErr            error
	listCryptoKeyVersionsErr     error
	opts                         []option.ClientOption
	pemCrc32C                    *wrapperspb.Int64Value
	signatureCrc32C              *wrapperspb.Int64Value
	store                        fakeStore
	tokeninfo                    *oauth2.Tokeninfo
	updateCryptoKeyErr           error
	keyIsDisabled                bool
}

func (k *fakeKMSClient) setAsymmetricSignErr(fakeError error) { _ = "STUB: not implemented"; return }

func (k *fakeKMSClient) setCreateCryptoKeyErr(fakeError error) { _ = "STUB: not implemented"; return }

func (k *fakeKMSClient) setInitialCryptoKeyVersionState(state kmspb.CryptoKeyVersion_CryptoKeyVersionState) {
	_ = "STUB: not implemented"
	return
}

func (k *fakeKMSClient) setDestroyCryptoKeyVersionErr(fakeError error) {
	_ = "STUB: not implemented"
	return
}

func (k *fakeKMSClient) setDestroyTime(fakeDestroyTime *timestamppb.Timestamp) {
	_ = "STUB: not implemented"
	return
}

func (k *fakeKMSClient) setGetCryptoKeyVersionErr(fakeError error) {
	_ = "STUB: not implemented"
	return
}

func (k *fakeKMSClient) setIsKeyDisabled(ok bool) { _ = "STUB: not implemented"; return }

func (k *fakeKMSClient) setGetPublicKeySequentialErrs(fakeError error, count int) {
	_ = "STUB: not implemented"
	return
}

func (k *fakeKMSClient) nextGetPublicKeySequentialErr() error {
	_ = "STUB: not implemented"
	return nil
}

func (k *fakeKMSClient) setGetTokeninfoErr(fakeError error) { _ = "STUB: not implemented"; return }

func (k *fakeKMSClient) setListCryptoKeysErr(fakeError error) { _ = "STUB: not implemented"; return }

func (k *fakeKMSClient) setPEMCrc32C(pemCrc32C *wrapperspb.Int64Value) {
	_ = "STUB: not implemented"
	return
}

func (k *fakeKMSClient) setSignatureCrc32C(signatureCrc32C *wrapperspb.Int64Value) {
	_ = "STUB: not implemented"
	return
}

func (k *fakeKMSClient) setUpdateCryptoKeyErr(fakeError error) { _ = "STUB: not implemented"; return }

func (k *fakeKMSClient) AsymmetricSign(_ context.Context, signReq *kmspb.AsymmetricSignRequest, _ ...gax.CallOption) (*kmspb.AsymmetricSignResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Override the SignatureCrc32C value

func (k *fakeKMSClient) Close() error { _ = "STUB: not implemented"; return nil }

func (k *fakeKMSClient) CreateCryptoKey(_ context.Context, req *kmspb.CreateCryptoKeyRequest, _ ...gax.CallOption) (*kmspb.CryptoKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Return a clone to decouple the caller's copy from the fake
// store's copy, mimicking the behavior of a real KMS service
// where no memory is shared between client and server.

func (k *fakeKMSClient) CreateCryptoKeyVersion(_ context.Context, req *kmspb.CreateCryptoKeyVersionRequest, _ ...gax.CallOption) (*kmspb.CryptoKeyVersion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *fakeKMSClient) DestroyCryptoKeyVersion(_ context.Context, req *kmspb.DestroyCryptoKeyVersionRequest, _ ...gax.CallOption) (*kmspb.CryptoKeyVersion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *fakeKMSClient) GetCryptoKeyVersion(_ context.Context, req *kmspb.GetCryptoKeyVersionRequest, _ ...gax.CallOption) (*kmspb.CryptoKeyVersion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *fakeKMSClient) GetPublicKey(_ context.Context, req *kmspb.GetPublicKeyRequest, _ ...gax.CallOption) (*kmspb.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Override pemCrc32C

func (k *fakeKMSClient) GetTokeninfo() (*oauth2.Tokeninfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *fakeKMSClient) ListCryptoKeys(_ context.Context, req *kmspb.ListCryptoKeysRequest, _ ...gax.CallOption) cryptoKeyIterator {
	_ = "STUB: not implemented"
	return *new(cryptoKeyIterator)
}

// Make sure that it's within the same Key Ring.
// The Key Ring name es specified in req.Parent.
// The Key Ring name is three levels up from the CryptoKey name.

// Key Ring doesn't match.

// We Have a simplified filtering logic in this fake implementation,
// where we only care about the spire-active and spire-last-update labels.

func (k *fakeKMSClient) ListCryptoKeyVersions(_ context.Context, req *kmspb.ListCryptoKeyVersionsRequest, _ ...gax.CallOption) cryptoKeyVersionIterator {
	_ = "STUB: not implemented"
	return *new(cryptoKeyVersionIterator)
}

// We Have a simplified filtering logic in this fake implementation,
// where we only support filtering by enabled status.

func (k *fakeKMSClient) ResourceIAM(string) iamHandler {
	_ = "STUB: not implemented"
	return *new(iamHandler)
}

func (k *fakeKMSClient) UpdateCryptoKey(_ context.Context, req *kmspb.UpdateCryptoKeyRequest, _ ...gax.CallOption) (*kmspb.CryptoKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Clone to decouple the fake store's copy from the caller's
// copy, preventing shared-memory data races.

func (k *fakeKMSClient) createFakeCryptoKeyVersion(cryptoKey *kmspb.CryptoKey, version string) (*fakeCryptoKeyVersion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *fakeKMSClient) getDefaultPolicy() *iam.Policy3 { _ = "STUB: not implemented"; return nil }

func (k *fakeKMSClient) putFakeCryptoKeys(fakeCryptoKeys []*fakeCryptoKey) {
	_ = "STUB: not implemented"
	return
}

func newKMSClientFake(t *testing.T, c *clock.Mock) *fakeKMSClient {
	_ = "STUB: not implemented"
	return nil
}

func newFakeStore(c *clock.Mock) fakeStore { _ = "STUB: not implemented"; return *new(fakeStore) }
