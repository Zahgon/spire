package credentialcomposer

import (
	"context"
	"crypto"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	credentialcomposerv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/credentialcomposer/v1"
	"github.com/spiffe/spire/pkg/common/plugin"
)

var _ CredentialComposer = (*V1)(nil)

type V1 struct {
	plugin.Facade
	credentialcomposerv1.CredentialComposerPluginClient
}

func (v1 V1) ComposeServerX509CA(ctx context.Context, attributes X509CAAttributes) (X509CAAttributes, error) {
	_ = "STUB: not implemented"
	return *new(X509CAAttributes), nil
}

func (v1 V1) ComposeServerX509SVID(ctx context.Context, attributes X509SVIDAttributes) (X509SVIDAttributes, error) {
	_ = "STUB: not implemented"
	return *new(X509SVIDAttributes), nil
}

func (v1 V1) ComposeAgentX509SVID(ctx context.Context, id spiffeid.ID, publicKey crypto.PublicKey, attributes X509SVIDAttributes) (X509SVIDAttributes, error) {
	_ = "STUB: not implemented"
	return *new(X509SVIDAttributes), nil
}

func (v1 V1) ComposeWorkloadX509SVID(ctx context.Context, id spiffeid.ID, publicKey crypto.PublicKey, attributes X509SVIDAttributes) (X509SVIDAttributes, error) {
	_ = "STUB: not implemented"
	return *new(X509SVIDAttributes), nil
}

func (v1 V1) ComposeWorkloadJWTSVID(ctx context.Context, id spiffeid.ID, attributes JWTSVIDAttributes) (JWTSVIDAttributes, error) {
	_ = "STUB: not implemented"
	return *new(JWTSVIDAttributes), nil
}

func (v1 V1) handleX509CAAttributesResponse(attributes X509CAAttributes, resp x509CAAttributesResponseV1, respErr error) (_ X509CAAttributes, err error) {
	_ = "STUB: not implemented"
	return *new(X509CAAttributes), nil
}

func (v1 V1) handleX509SVIDAttributesResponse(attributes X509SVIDAttributes, resp x509SVIDAttributesResponseV1, respErr error) (_ X509SVIDAttributes, err error) {
	_ = "STUB: not implemented"
	return *new(X509SVIDAttributes), nil
}

func (v1 V1) handleJWTSVIDAttributesResponse(attributes JWTSVIDAttributes, resp jwtSVIDAttributesResponseV1, respErr error) (_ JWTSVIDAttributes, err error) {
	_ = "STUB: not implemented"
	return *new(JWTSVIDAttributes), nil
}

func x509CAAttributesToV1(attributes X509CAAttributes) (*credentialcomposerv1.X509CAAttributes, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type x509CAAttributesResponseV1 interface {
	GetAttributes() *credentialcomposerv1.X509CAAttributes
}

func x509CAAttributesFromV1(pb *credentialcomposerv1.X509CAAttributes) (attributes X509CAAttributes, err error) {
	_ = "STUB: not implemented"
	return *new(X509CAAttributes), nil
}

type x509SVIDAttributesResponseV1 interface {
	GetAttributes() *credentialcomposerv1.X509SVIDAttributes
}

func x509SVIDAttributesToV1(attributes X509SVIDAttributes) (*credentialcomposerv1.X509SVIDAttributes, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func x509SVIDAttributesFromV1(pb *credentialcomposerv1.X509SVIDAttributes) (attributes X509SVIDAttributes, err error) {
	_ = "STUB: not implemented"
	return *new(X509SVIDAttributes), nil
}

type jwtSVIDAttributesResponseV1 interface {
	GetAttributes() *credentialcomposerv1.JWTSVIDAttributes
}

func jwtSVIDAttributesToV1(attributes JWTSVIDAttributes) (*credentialcomposerv1.JWTSVIDAttributes, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// structpb.NewValue cannot handle Go types such as jwt.NumericDate so we marshal them into their JSON representation first

func jwtSVIDAttributesFromV1(pb *credentialcomposerv1.JWTSVIDAttributes) JWTSVIDAttributes {
	_ = "STUB: not implemented"
	return *new(JWTSVIDAttributes)
}

func subjectFromV1(in *credentialcomposerv1.DistinguishedName) (pkix.Name, error) {
	_ = "STUB: not implemented"
	return *new(pkix.Name), nil
}

func subjectToV1(in pkix.Name) (*credentialcomposerv1.DistinguishedName, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func policyIdentifiersFromV1(ins []string) ([]x509.OID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func policyIdentifiersToV1(ins []x509.OID) []string { _ = "STUB: not implemented"; return nil }

func extraExtensionsFromV1(ins []*credentialcomposerv1.X509Extension) ([]pkix.Extension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extraExtensionsToV1(ins []pkix.Extension) []*credentialcomposerv1.X509Extension {
	_ = "STUB: not implemented"
	return nil
}

func extraNamesToV1(ins []pkix.AttributeTypeAndValue) ([]*credentialcomposerv1.AttributeTypeAndValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extraNamesFromV1(ins []*credentialcomposerv1.AttributeTypeAndValue) ([]pkix.AttributeTypeAndValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseOID(s string) (_ asn1.ObjectIdentifier, err error) {
	_ = "STUB: not implemented"
	return *new(asn1.ObjectIdentifier), nil
}
