package nodeinfo

import (
	"github.com/reiver/go-nodeinfo/shared"
)

// SetUserAgent sets the User-Agent used by any (internal) HTTP-client in this package when making an HTTP-request to any server.
//
// For example:
//
//	import "github.com/reiver/go-nodeinfo"
//	
//	// ...
//	
//	nodeinfo.SetUserAgent("ExampleAppt/2.71 (+http://example.com/example-app)")
//
// You can get the value of the User-Agent by calling [UserAgent].
func SetUserAgent(value string) {
	shared.UserAgent = value
}

// UserAgent returns the User-Agent used by any (internal) HTTP-client in this package when making an HTTP-request to any server.
//
// For example:
//
//	import "github.com/reiver/go-nodeinfo"
//	
//	// ...
//	
//	userAgent := nodeinfo.UserAgent()
//
// This can be changed to customize the User-Agent sent using [SetUserAgent].
func UserAgent() string {
	return shared.UserAgent
}
