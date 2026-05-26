package k8sconfigmap

import (
	"context"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// kubernetesClient defines the interface for Kubernetes operations.
type kubernetesClient interface {
	// ApplyConfigMap applies a ConfigMap, creating it if it does not exist or updating it if it does.
	// If the ConfigMap already exists, it will be updated with the provided data.
	// If it does not exist, it will be created with the provided data.
	// This function uses the Apply method to ensure idempotency.
	ApplyConfigMap(ctx context.Context, cluster *Cluster, data []byte) error
}

// k8sClient implements the kubernetesClient interface.
type k8sClient struct {
	clientset kubernetes.Interface
}

func (c *k8sClient) ApplyConfigMap(ctx context.Context, cluster *Cluster, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// newK8sClient creates a new Kubernetes client based on the provided configuration.
func newK8sClient(kubeConfigPath string) (kubernetesClient, error) {
	_ = "STUB: not implemented"
	return *new(kubernetesClient), nil
}

// getKubeConfig returns a Kubernetes configuration based on the provided path.
// If the path is empty, it uses the in-cluster configuration.
func getKubeConfig(configPath string) (*rest.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
