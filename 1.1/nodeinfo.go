package nodeinfo

import (
	"github.com/reiver/go-nodeinfo/internal"
)

type NodeInfo struct {
	Software Software
	Protocols Protocols
	Services Services
	OpenRegistrations bool
	Usage Usage
	MetaData map[string]any
}

func (receiver NodeInfo) MarshalJSON() ([]byte, error) {
	const version string = "1.1"

	return internal.NodeInfo1MarshalJSON(version, receiver.Software,  receiver.Protocols,  receiver.Services,  receiver.OpenRegistrations,  receiver.Usage,  receiver.MetaData)
}
