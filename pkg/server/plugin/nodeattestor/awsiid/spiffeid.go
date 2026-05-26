package awsiid

import (
	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/common/agentpathtemplate"
)

var defaultAgentPathTemplate = agentpathtemplate.MustParse("/{{ .PluginName}}/{{ .AccountID }}/{{ .Region }}/{{ .InstanceID }}")

type agentPathTemplateData struct {
	InstanceID  string
	AccountID   string
	Region      string
	PluginName  string
	TrustDomain string
	Tags        instanceTags
}

type instanceTags map[string]string

// makeAgentID creates an agent ID from IID data
func makeAgentID(td spiffeid.TrustDomain, agentPathTemplate *agentpathtemplate.Template, doc imds.InstanceIdentityDocument, tags instanceTags) (spiffeid.ID, error) {
	_ = "STUB: not implemented"
	return *new(spiffeid.ID), nil
}
