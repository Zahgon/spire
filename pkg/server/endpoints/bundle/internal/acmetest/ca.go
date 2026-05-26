// Copyright (c) 2018 The Go Authors. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// Package acmetest provides types for testing acme and autocert packages.
//
// SPIRE modifications:
// - Verifies signatures on incoming requests to ensure requests are signed
//   appropriately by the SPIRE KeyManager signers.
// - Fails new-reg requests if the terms-of-service has not been accepted

// nolint // forked code
package acmetest

import (
	"crypto"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/go-jose/go-jose/v4"
)

var allowedJWTSignatureAlgorithms = []jose.SignatureAlgorithm{
	jose.RS256,
	jose.RS384,
	jose.RS512,
	jose.ES256,
	jose.ES384,
	jose.ES512,
	jose.PS256,
	jose.PS384,
	jose.PS512,
}

// CAServer is a simple test server which implements ACME spec bits needed for testing.
type CAServer struct {
	rootKey      crypto.Signer
	rootCert     []byte // DER encoding
	rootTemplate *x509.Certificate

	t              *testing.T
	server         *httptest.Server
	issuer         pkix.Name
	challengeTypes []string
	url            string
	roots          *x509.CertPool
	eabRequired    bool

	mu             sync.Mutex
	certCount      int                           // number of issued certs
	acctRegistered bool                          // set once an account has been registered
	domainAddr     map[string]string             // domain name to addr:port resolution
	domainGetCert  map[string]getCertificateFunc // domain name to GetCertificate function
	domainHandler  map[string]http.Handler       // domain name to Handle function
	validAuthz     map[string]*authorization     // valid authz, keyed by domain name
	authorizations []*authorization              // all authz, index is used as ID
	orders         []*order                      // index is used as order ID
	errors         []error                       // encountered client errors

	accountKeysMu sync.Mutex
	accountKeys   map[string]any
}

type getCertificateFunc func(hello *tls.ClientHelloInfo) (*tls.Certificate, error)

// NewCAServer creates a new ACME test server. The returned CAServer issues
// certs signed with the CA roots available in the Roots field.
func NewCAServer(t *testing.T) *CAServer { _ = "STUB: not implemented"; return nil }

func (ca *CAServer) generateRoot() { _ = "STUB: not implemented"; return }

// IssuerName sets the name of the issuing CA.
func (ca *CAServer) IssuerName(name pkix.Name) *CAServer { _ = "STUB: not implemented"; return nil }

// ChallengeTypes sets the supported challenge types.
func (ca *CAServer) ChallengeTypes(types ...string) *CAServer {
	_ = "STUB: not implemented"
	return nil
}

// URL returns the server address, after Start has been called.
func (ca *CAServer) URL() string { _ = "STUB: not implemented"; return "" }

// Roots returns a pool cointaining the CA root.
func (ca *CAServer) Roots() *x509.CertPool { _ = "STUB: not implemented"; return nil }

// ExternalAccountRequired makes an EAB JWS required for account registration.
func (ca *CAServer) ExternalAccountRequired() *CAServer { _ = "STUB: not implemented"; return nil }

// Start starts serving requests. The server address becomes available in the
// URL field.
func (ca *CAServer) Start() *CAServer { _ = "STUB: not implemented"; return nil }

func (ca *CAServer) serverURL(format string, arg ...interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

func (ca *CAServer) addr(domain string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func (ca *CAServer) getCert(domain string) (getCertificateFunc, bool) {
	_ = "STUB: not implemented"
	return *new(getCertificateFunc), false
}

func (ca *CAServer) getHandler(domain string) (http.Handler, bool) {
	_ = "STUB: not implemented"
	return *new(http.Handler), false
}

func (ca *CAServer) httpErrorf(w http.ResponseWriter, code int, format string, a ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// FORK DEVIATION FROM ORIGINAL CODE
// We intentionally comment out this line because
// TestACMEAuth/new-account-tos-not-accepted in pkg/server/endpoints/bundle/server_test.go
// tests a condition where an error is sent back to the client,
// and we don't want to fail the test prematurely before we can assert on the error condition.
// ca.t.Errorf(format, a...)

// Resolve adds a domain to address resolution for the ca to dial to
// when validating challenges for the domain authorization.
func (ca *CAServer) Resolve(domain, addr string) { _ = "STUB: not implemented"; return }

// ResolveGetCertificate redirects TLS connections for domain to f when
// validating challenges for the domain authorization.
func (ca *CAServer) ResolveGetCertificate(domain string, f getCertificateFunc) {
	_ = "STUB: not implemented"
	return
}

// ResolveHandler redirects HTTP requests for domain to f when
// validating challenges for the domain authorization.
func (ca *CAServer) ResolveHandler(domain string, h http.Handler) {
	_ = "STUB: not implemented"
	return
}

type discovery struct {
	NewNonce   string `json:"newNonce"`
	NewAccount string `json:"newAccount"`
	NewOrder   string `json:"newOrder"`
	NewAuthz   string `json:"newAuthz"`

	Meta discoveryMeta `json:"meta,omitempty"`
}

type discoveryMeta struct {
	TermsOfService          string `json:"termsOfService,omitempty"`
	ExternalAccountRequired bool   `json:"externalAccountRequired,omitempty"`
}

type challenge struct {
	URI   string `json:"uri"`
	Type  string `json:"type"`
	Token string `json:"token"`
}

type authorization struct {
	Status     string      `json:"status"`
	Challenges []challenge `json:"challenges"`

	domain string
	id     int
}

type order struct {
	Status      string   `json:"status"`
	AuthzURLs   []string `json:"authorizations"`
	FinalizeURL string   `json:"finalize"`    // CSR submit URL
	CertURL     string   `json:"certificate"` // already issued cert

	leaf []byte // issued cert in DER format
}

func (ca *CAServer) handle(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// TODO: Verify nonce header for all POST requests.

// Discovery request.

// Nonce requests.

// Nonce values are always set. Nothing else to do.

// Client key registration request.

// TODO: Check the user account key against a ca.accountKeys?

// New order request.

// Existing order status requests.

// Accept challenge requests.

// Get authorization status requests.

// Note we don't invalidate authorized orders as we should.

// Certificate issuance request.

// Validate CSR request.

// Issue the certificate.

// Already issued cert download requests.

// storedOrder retrieves a previously created order at index i.
// It requires ca.mu to be locked.
func (ca *CAServer) storedOrder(i string) (*order, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// storedAuthz retrieves a previously created authz at index i.
// It requires ca.mu to be locked.
func (ca *CAServer) storedAuthz(i string) (*authorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// authz returns an existing valid authorization for the identifier or creates a
// new one. It requires ca.mu to be locked.
func (ca *CAServer) authz(identifier string) *authorization { _ = "STUB: not implemented"; return nil }

// leafCert issues a new certificate.
// It requires ca.mu to be locked.
func (ca *CAServer) leafCert(csr *x509.CertificateRequest) (der []byte, err error) {
	_ = "STUB: not implemented"
	// next leaf cert serial number
	return nil, nil
}

// LeafCert issues a leaf certificate.
func (ca *CAServer) LeafCert(name, keyType string, notBefore, notAfter time.Time) *tls.Certificate {
	_ = "STUB: not implemented"
	return nil
}

// next leaf cert serial number

func (ca *CAServer) validateChallenge(authz *authorization, typ string) {
	_ = "STUB: not implemented"
	return
}

func (ca *CAServer) updatePendingOrders() {
	_ = "STUB: not implemented"
	// Update all pending orders.
	// An order becomes "ready" if all authorizations are "valid".
	// An order becomes "invalid" if any authorization is "invalid".
	// Status changes: https://tools.ietf.org/html/rfc8555#section-7.1.6
	return
}

func (ca *CAServer) validateAuthzURLs(urls []string, orderNum int) (countValid, countInvalid int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (ca *CAServer) verifyALPNChallenge(a *authorization) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: support selecting ECDSA.

// See RFC 8737, Section 6.1.

// TODO: check the token.

func (ca *CAServer) verifyHTTPChallenge(a *authorization) error {
	_ = "STUB: not implemented"
	return nil
}

func (ca *CAServer) decodePayload(v any, r io.Reader) error { _ = "STUB: not implemented"; return nil }

// TODO: strict validation of keyid

// payload := jws.UnsafePayloadWithoutVerification()

// TODO: calculate per-account key id

func (ca *CAServer) lookupAccountKey(kid string) any { _ = "STUB: not implemented"; return *new(any) }

func (ca *CAServer) setAccountKey(kid string, key any) { _ = "STUB: not implemented"; return }

func challengeToken(domain, challType string, authzID int) string {
	_ = "STUB: not implemented"
	return ""
}
