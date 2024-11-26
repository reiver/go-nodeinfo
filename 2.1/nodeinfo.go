package nodeinfo

import (
	"encoding/json"

	"github.com/reiver/go-nodeinfo/internal"
)

type NodeInfo struct {
	Software Software
	Protocols []string
	Services Services
	OpenRegistrations bool
	Usage Usage
	MetaData map[string]any
}

var _ json.Marshaler = NodeInfo{}

func (receiver NodeInfo) MarshalJSON() ([]byte, error) {
	const version string = "2.1"

	return internal.NodeInfo2Dot1MarshalJSON(version, receiver.Software, receiver.Protocols, receiver.Services, receiver.OpenRegistrations, receiver.Usage, receiver.MetaData)
}
