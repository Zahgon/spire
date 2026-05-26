package api

import (
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/spiffe/spire/proto/spire/common"
)

var (
	UpdateAttestedNodeCertificateMask = &common.AttestedNodeMask{
		CertNotAfter:        true,
		CertSerialNumber:    true,
		NewCertNotAfter:     true,
		NewCertSerialNumber: true,
		CanReattest:         true,
	}
)

func ProtoFromAttestedNode(n *common.AttestedNode) (*types.Agent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
