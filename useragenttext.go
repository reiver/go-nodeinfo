package nodeinfo

import (
	"github.com/reiver/go-nodeinfo/shared"
)

// SetUserAgentText sets the User-Agent used by any (internal) HTTP-client in this package when making an HTTP-request to any server.
//
// For example:
//
//	import "github.com/reiver/go-nodeinfo"
//	
//	// ...
//	
//	nodeinfo.SetUserAgentText("ExampleAppt/2.71 (+http://example.com/example-app)")
//
// You can get the value of the User-Agent by calling [UserAgentText].
func SetUserAgentText(value string) {
	shared.UserAgentText = value
}

// UserAgentText returns the User-Agent used by any (internal) HTTP-client in this package when making an HTTP-request to any server.
//
// For example:
//
//	import "github.com/reiver/go-nodeinfo"
//	
//	// ...
//	
//	userAgent := nodeinfo.UserAgentText()
//
// This can be changed to customize the User-Agent sent using [SetUserAgentText].
func UserAgentText() string {
	return shared.UserAgentText
}
