package shared

var (
	// UserAgent is the User-Agent used by any (internal) HTTP-client in this package when making an HTTP-request to any server.
	// This can be changed to customize the User-Agent sent.
	//
	// For example:
	//
	//	import "github.com/reiver/go-nodeinfo"
	//	
	//	// ...
	//	
	//	nodeinfo.SetUserAgentText("ExampleAppt/2.71 (+http://example.com/example-app)")
	UserAgentText = "reiver-nodeinfo/0.0 (+https://github.com/reiver/go-nodeinfo)"
)
