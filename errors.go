package nodeinfo

import (
	"github.com/reiver/go-erorr"
)

const (
	errNilBytes    = erorr.Error("nodeinfo: nil bytes")
	errNilReceiver = erorr.Error("nodeinfo: nil receiver")
)
