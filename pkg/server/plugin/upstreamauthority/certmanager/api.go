package certmanager

import (
	"context"

	upstreamauthorityv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/upstreamauthority/v1"
	cmapi "github.com/spiffe/spire/pkg/server/plugin/upstreamauthority/certmanager/internal/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	scheme = runtime.NewScheme()
)

func init() {
	schemeGroupVersion := schema.GroupVersion{Group: "cert-manager.io", Version: "v1"}
	scheme.AddKnownTypes(schemeGroupVersion,
		&cmapi.CertificateRequest{},
		&cmapi.CertificateRequestList{},
	)
	metav1.AddToGroupVersion(scheme, schemeGroupVersion)
}

func (p *Plugin) buildCertificateRequest(request *upstreamauthorityv1.MintX509CARequest) (*cmapi.CertificateRequest, error) {
	_ = "STUB: not implemented"
	// Build PEM encoded CSR
	return nil, nil
}

// cleanupStaleCertificateRequests will attempt to delete CertificateRequests
// that have been created for this trust domain, and are in a terminal state.
// Terminal states are:
// - The request has been Denied
// - The request is in a Ready state
// - The request is in a Failed state
func (p *Plugin) cleanupStaleCertificateRequests(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// certificateRequestHasCondition will return true and the condition if the
// given CertificateRequest has a condition matching the provided
// CertificateRequestCondition.
// Only the Type and Status field will be used in the comparison, unless the
// given condition has set a Reason.
func certificateRequestHasCondition(cr *cmapi.CertificateRequest, c cmapi.CertificateRequestCondition) (bool, cmapi.CertificateRequestCondition) {
	_ = "STUB: not implemented"
	return false, *new(cmapi.CertificateRequestCondition)
}
