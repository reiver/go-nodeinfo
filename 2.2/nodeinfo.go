package nodeinfo

import (
	"github.com/reiver/go-nodeinfo/internal"
)

type NodeInfo struct {
	Instance Instance
	Software Software
	Protocols []string
	Services Services
	OpenRegistrations bool
	Usage Usage
	MetaData map[string]any
}

func (receiver NodeInfo) MarshalJSON() ([]byte, error) {
	const version string = "2.2"

	return internal.NodeInfo2Dot2MarshalJSON(version, receiver.Instance, receiver.Software, receiver.Protocols, receiver.Services, receiver.OpenRegistrations, receiver.Usage, receiver.MetaData)
}
