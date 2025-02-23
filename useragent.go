package nodeinfo

import (
	"github.com/reiver/go-nodeinfo/shared"
)

func SetUserAgent(value string) {
	shared.UserAgent = value
}

func UserAgent() string {
	return shared.UserAgent
}
