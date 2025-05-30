package shared

var (
	// ServerText is the Server (name) used by any (internal) HTTP-server in this package when making an HTTP-response to any client.
	// This can be changed to customize the User-Agent sent.
	//
	// For example:
	//
	//	import "github.com/reiver/go-nodeinfo"
	//	
	//	// ...
	//	
	//	nodeinfo.SetServerText("Example/1.0")
	ServerText = "(reiver-nodeinfo)"
)
