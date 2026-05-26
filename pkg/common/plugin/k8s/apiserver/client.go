package apiserver

import (
	"context"

	authv1 "k8s.io/api/authentication/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

// Client is a client for querying k8s API server
type Client interface {
	// GetNode returns the node object for the given node name
	GetNode(ctx context.Context, nodeName string) (*v1.Node, error)

	// GetPod returns the pod object for the given pod name and namespace
	GetPod(ctx context.Context, namespace, podName string) (*v1.Pod, error)

	// ValidateToken queries k8s token review API and returns information about the given token
	ValidateToken(ctx context.Context, token string, audiences []string) (*authv1.TokenReviewStatus, error)
}

type client struct {
	kubeConfigFilePath string

	// loadClientHook is used to inject a fake loadClient on tests
	loadClientHook func(string) (kubernetes.Interface, error)
}

// New creates a new Client.
// There are two cases:
// - If a kubeConfigFilePath is provided, config is taken from that file -> use for clients running out of a k8s cluster
// - If not (empty kubeConfigFilePath), InClusterConfig is used          -> use for clients running in a k8s cluster
func New(kubeConfigFilePath string) Client { _ = "STUB: not implemented"; return *new(Client) }

func (c *client) GetPod(ctx context.Context, namespace, podName string) (*v1.Pod, error) {
	_ = "STUB: not implemented"
	// Validate inputs
	return nil, nil
}

// Reload config

// Get pod

func (c *client) GetNode(ctx context.Context, nodeName string) (*v1.Node, error) {
	_ = "STUB: not implemented"
	// Validate inputs
	return nil, nil
}

// Reload config

// Get node

func (c *client) ValidateToken(ctx context.Context, token string, audiences []string) (*authv1.TokenReviewStatus, error) {
	_ = "STUB: not implemented"
	// Reload config
	return nil, nil
}

// Create token review request

// Do request

// Evaluate token review response (review server will populate TokenReview.Status field)

// Ensure the audiences returned in the status are compatible with those requested
// in the TokenReviewSpec (if any). This is to ensure the validator is
// audience aware.
// See the documentation on the Status Audiences field.

func loadClient(kubeConfigFilePath string) (kubernetes.Interface, error) {
	_ = "STUB: not implemented"
	return *new(kubernetes.Interface), nil
}
