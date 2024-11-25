# go-nodeinfo/1.1

Package **nodeinfo/1.1** implements the [NodeInfo](https://github.com/jhass/nodeinfo) 1.1 protocol, for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-nodeinfo

[![GoDoc](https://godoc.org/github.com/reiver/go-nodeinfo?status.svg)](https://godoc.org/github.com/reiver/go-nodeinfo)

## Example

```golang
import (
	"github.com/reiver/go-nodeinfo/1.1"
	"github.com/reiver/go-opt"
}

// ...

var nodeInfo nodeinfo.NodeInfo = NodeInfo: nodeinfo.NodeInfo{
	Software: nodeinfo.Software{
		Name: "diaspora",
		Version: "0.5.0",
	},
	Protocols: nodeinfo.Protocols{
		Inbound: []string{"diaspora"},
		Outbound: []string{"diaspora"},
	},
	Services: nodeinfo.Services{
		Inbound: []string{"gnusocial"},
		Outbound: []string{"facebook", "twitter"},
	},
	OpenRegistrations: true,
	Usage: nodeinfo.Usage{
		Users: nodeinfo.Users{
			Total: opt.Something("123"),
			ActiveHalfYear: opt.Something("42"),
			ActiveMonth: opt.Something("23"),
		},
		LocalPosts:    opt.Something("500"),
		LocalComments: opt.Something("1000"),
	},
	MetaData: map[string]any{
		"chat_enabled": true,
	},
}

// ...

err := json.NewEncoder(responseWriter).Encode(nodeInfo)

// ...
```


## Import

To import package **nodeinfo** use `import` code like the following:
```
import "github.com/reiver/go-nodeinfo/1.1"
```

## Installation

To install package **nodeinfo** do the following:
```
GOPROXY=direct go get github.com/reiver/go-nodeinfo
```

## Author

Package **nodeinfo** was written by [Charles Iliya Krempeaux](http://reiver.link)
