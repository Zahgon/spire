package k8s

import (
	"github.com/go-jose/go-jose/v4/jwt"
	authv1 "k8s.io/api/authentication/v1"
)

const (
	k8sPodNameKey = "authentication.kubernetes.io/pod-name"
	k8sPodUIDKey  = "authentication.kubernetes.io/pod-uid"
)

// SATClaims represents claims in a service account token, for example:
//
//	{
//	  "iss": "kubernetes/serviceaccount",
//	  "kubernetes.io/serviceaccount/namespace": "spire",
//	  "kubernetes.io/serviceaccount/secret.name": "spire-agent-token-zjr8v",
//	  "kubernetes.io/serviceaccount/service-account.name": "spire-agent",
//	  "kubernetes.io/serviceaccount/service-account.uid": "1881e84f-b612-11e8-a543-0800272c6e42",
//	  "sub": "system:serviceaccount:spire:spire-agent"
//	}
type SATClaims struct {
	jwt.Claims
	Namespace          string `json:"kubernetes.io/serviceaccount/namespace"`
	ServiceAccountName string `json:"kubernetes.io/serviceaccount/service-account.name"`

	// This struct is included in case that a projected service account token is
	// parsed as a regular service account token
	K8s struct {
		Namespace      string `json:"namespace"`
		ServiceAccount struct {
			Name string `json:"name"`
		} `json:"serviceaccount"`
	} `json:"kubernetes.io"`
}

// PSATClaims represents claims in a projected service account token, for example:
//
//	{
//		 "aud": [
//		   "spire-server"
//		 ],
//		 "exp": 1550850854,
//		 "iat": 1550843654,
//		 "iss": "api",
//		 "kubernetes.io": {
//		   "namespace": "spire",
//		   "pod": {
//		 	"name": "spire-agent-5d84p",
//		 	"uid": "56857f33-36a9-11e9-860c-080027b25557"
//		   },
//		   "serviceaccount": {
//		 	"name": "spire-agent",
//		 	"uid": "ca29bd95-36a8-11e9-b8af-080027b25557"
//		   }
//		 },
//		 "nbf": 1550843654,
//		 "sub": "system:serviceaccount:spire:spire-agent"
//	}
type PSATClaims struct {
	jwt.Claims
	K8s struct {
		Namespace string `json:"namespace"`

		Pod struct {
			Name string `json:"name"`
			UID  string `json:"uid"`
		} `json:"pod"`

		ServiceAccount struct {
			Name string `json:"name"`
			UID  string `json:"uid"`
		} `json:"serviceaccount"`
	} `json:"kubernetes.io"`
}

type SATAttestationData struct {
	Cluster string `json:"cluster"`
	Token   string `json:"token"`
}

type PSATAttestationData struct {
	Cluster string `json:"cluster"`
	Token   string `json:"token"`
}

func AgentID(pluginName, trustDomain, cluster, uuid string) string {
	_ = "STUB: not implemented"
	return ""
}

func MakeSelectorValue(kind string, values ...string) string { _ = "STUB: not implemented"; return "" }

// GetNamesFromTokenStatus parses a fully qualified k8s username like: 'system:serviceaccount:spire:spire-agent'
// from tokenStatus. The string is split and the last two names are returned: namespace and service account name
func GetNamesFromTokenStatus(tokenStatus *authv1.TokenReviewStatus) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// GetPodNameFromTokenStatus extracts pod name from a tokenReviewStatus type
func GetPodNameFromTokenStatus(tokenStatus *authv1.TokenReviewStatus) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetPodUIDFromTokenStatus extracts pod UID from a tokenReviewStatus type
func GetPodUIDFromTokenStatus(tokenStatus *authv1.TokenReviewStatus) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
