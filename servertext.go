package nodeinfo

import (
	"github.com/reiver/go-nodeinfo/shared"
)

// SetServerText sets the Server (name) used by any (internal) HTTP-server in this package when making an HTTP-response to any client.
//
// For example:
//
//	import "github.com/reiver/go-nodeinfo"
//	
//	// ...
//	
//	nodeinfo.SetServerText("Example/1.0.1 (hello world)")
//
// You can get the value of the Server (name) by calling [ServerText].
func SetServerText(value string) {
	shared.ServerText = value
}

// ServerText returns the Server (name) used by any (internal) HTTP-server in this package when making an HTTP-response to any client.
//
// For example:
//
//	import "github.com/reiver/go-nodeinfo"
//	
//	// ...
//	
//	Server := nodeinfo.ServerText()
//
// This can be changed to customize the Server (name) sent using [SetServerText].
func ServerText() string {
	return shared.ServerText
}
