package nodeinfo

import (
	"testing"

	"github.com/reiver/go-opt"
)

func TestResolveWellKnown(t *testing.T) {

	tests := []struct{
		WellKnown WellKnown
		Scheme string
		Host string
		Expected WellKnown
	}{
		{

		},



		{
			Scheme: "pow",
		},
		{
			Scheme: "http",
		},
		{
			Scheme: "https",
		},



		{
			Host: "banana.example",
		},



		{
			Scheme: "pow",
			Host: "banana.example",
		},
		{
			Scheme: "http",
			Host: "banana.example",
		},
		{
			Scheme: "https",
			Host: "banana.example",
		},



		{
			WellKnown: WellKnown{
				NodeInfo2: opt.Something("https://mastodon.social/nodeinfo/2.0"),
			},
			Scheme: "pow",
			Host: "banana.example",
			Expected: WellKnown{
				NodeInfo2: opt.Something("https://mastodon.social/nodeinfo/2.0"),
			},
		},
		{
			WellKnown: WellKnown{
				NodeInfo2: opt.Something("https://mastodon.social/nodeinfo/2.0"),
			},
			Scheme: "http",
			Host: "banana.example",
			Expected: WellKnown{
				NodeInfo2: opt.Something("https://mastodon.social/nodeinfo/2.0"),
			},
		},
		{
			WellKnown: WellKnown{
				NodeInfo2: opt.Something("https://mastodon.social/nodeinfo/2.0"),
			},
			Scheme: "https",
			Host: "banana.example",
			Expected: WellKnown{
				NodeInfo2: opt.Something("https://mastodon.social/nodeinfo/2.0"),
			},
		},



		{
			WellKnown: WellKnown{
				NodeInfo2: opt.Something("//mastodon.social/nodeinfo/2.0"),
			},
			Scheme: "pow",
			Host: "banana.example",
			Expected: WellKnown{
				NodeInfo2: opt.Something("pow://mastodon.social/nodeinfo/2.0"),
			},
		},
		{
			WellKnown: WellKnown{
				NodeInfo2: opt.Something("//mastodon.social/nodeinfo/2.0"),
			},
			Scheme: "http",
			Host: "banana.example",
			Expected: WellKnown{
				NodeInfo2: opt.Something("http://mastodon.social/nodeinfo/2.0"),
			},
		},
		{
			WellKnown: WellKnown{
				NodeInfo2: opt.Something("//mastodon.social/nodeinfo/2.0"),
			},
			Scheme: "https",
			Host: "banana.example",
			Expected: WellKnown{
				NodeInfo2: opt.Something("https://mastodon.social/nodeinfo/2.0"),
			},
		},



/*
		{
			WellKnown: WellKnown{
				NodeInfo2: opt.Something("https:///nodeinfo/2.0"),
			},
			Scheme: "pow",
			Host: "banana.example",
			Expected: WellKnown{
				NodeInfo2: opt.Something("https://banana.example/nodeinfo/2.0"),
			},
		},
		{
			WellKnown: WellKnown{
				NodeInfo2: opt.Something("https:///nodeinfo/2.0"),
			},
			Scheme: "http",
			Host: "banana.example",
			Expected: WellKnown{
				NodeInfo2: opt.Something("https://banana.example/nodeinfo/2.0"),
			},
		},
		{
			WellKnown: WellKnown{
				NodeInfo2: opt.Something("https:///nodeinfo/2.0"),
			},
			Scheme: "https",
			Host: "banana.example",
			Expected: WellKnown{
				NodeInfo2: opt.Something("https://banana.example/nodeinfo/2.0"),
			},
		},
*/




		{
			WellKnown: WellKnown{
				NodeInfo2: opt.Something("/nodeinfo/2.0"),
			},
			Scheme: "pow",
			Host: "banana.example",
			Expected: WellKnown{
				NodeInfo2: opt.Something("pow://banana.example/nodeinfo/2.0"),
			},
		},
		{
			WellKnown: WellKnown{
				NodeInfo2: opt.Something("/nodeinfo/2.0"),
			},
			Scheme: "http",
			Host: "banana.example",
			Expected: WellKnown{
				NodeInfo2: opt.Something("http://banana.example/nodeinfo/2.0"),
			},
		},
		{
			WellKnown: WellKnown{
				NodeInfo2: opt.Something("/nodeinfo/2.0"),
			},
			Scheme: "https",
			Host: "banana.example",
			Expected: WellKnown{
				NodeInfo2: opt.Something("https://banana.example/nodeinfo/2.0"),
			},
		},



		{
			WellKnown: WellKnown{
				NodeInfo2Dot2: opt.Something("ftp://host.example/nodeinfo/2.2.json"),
				NodeInfo2Dot1: opt.Something("http://host.example/nodeinfo/2.1"),
				NodeInfo2:     opt.Something("https://host.example/nodeinfo/2.0"),
				NodeInfo1Dot1: opt.Something("/nodeinfo/2.1"),
				NodeInfo1:     opt.Something("nodeinfo/2.1"),
			},
			Scheme: "pow",
			Host: "example.com",
			Expected: WellKnown{
				NodeInfo2Dot2: opt.Something("ftp://host.example/nodeinfo/2.2.json"),
				NodeInfo2Dot1: opt.Something("http://host.example/nodeinfo/2.1"),
				NodeInfo2:     opt.Something("https://host.example/nodeinfo/2.0"),
				NodeInfo1Dot1: opt.Something("pow://example.com/nodeinfo/2.1"),
				NodeInfo1:     opt.Something("pow://example.com/nodeinfo/2.1"),
			},
		},
		{
			WellKnown: WellKnown{
				NodeInfo2Dot2: opt.Something("ftp://host.example/nodeinfo/2.2.json"),
				NodeInfo2Dot1: opt.Something("http://host.example/nodeinfo/2.1"),
				NodeInfo2:     opt.Something("https://host.example/nodeinfo/2.0"),
				NodeInfo1Dot1: opt.Something("/nodeinfo/2.1"),
				NodeInfo1:     opt.Something("nodeinfo/2.1"),
			},
			Scheme: "http",
			Host: "example.com",
			Expected: WellKnown{
				NodeInfo2Dot2: opt.Something("ftp://host.example/nodeinfo/2.2.json"),
				NodeInfo2Dot1: opt.Something("http://host.example/nodeinfo/2.1"),
				NodeInfo2:     opt.Something("https://host.example/nodeinfo/2.0"),
				NodeInfo1Dot1: opt.Something("http://example.com/nodeinfo/2.1"),
				NodeInfo1:     opt.Something("http://example.com/nodeinfo/2.1"),
			},
		},
		{
			WellKnown: WellKnown{
				NodeInfo2Dot2: opt.Something("ftp://host.example/nodeinfo/2.2.json"),
				NodeInfo2Dot1: opt.Something("http://host.example/nodeinfo/2.1"),
				NodeInfo2:     opt.Something("https://host.example/nodeinfo/2.0"),
				NodeInfo1Dot1: opt.Something("/nodeinfo/2.1"),
				NodeInfo1:     opt.Something("nodeinfo/2.1"),
			},
			Scheme: "https",
			Host: "example.com",
			Expected: WellKnown{
				NodeInfo2Dot2: opt.Something("ftp://host.example/nodeinfo/2.2.json"),
				NodeInfo2Dot1: opt.Something("http://host.example/nodeinfo/2.1"),
				NodeInfo2:     opt.Something("https://host.example/nodeinfo/2.0"),
				NodeInfo1Dot1: opt.Something("https://example.com/nodeinfo/2.1"),
				NodeInfo1:     opt.Something("https://example.com/nodeinfo/2.1"),
			},
		},



		{
			WellKnown: WellKnown{
				NodeInfo2Dot2: opt.Something("ftp://host.example/nodeinfo/2.2.json"),
				NodeInfo2Dot1: opt.Something("http://host.example/nodeinfo/2.1"),
				NodeInfo2:     opt.Something("https://host.example/nodeinfo/2.0"),
				NodeInfo1Dot1: opt.Something("/nodeinfo/2.1"),
				NodeInfo1:     opt.Something("nodeinfo/2.1"),
			},
			Scheme: "pow",
			Host: "mastodon.social",
			Expected: WellKnown{
				NodeInfo2Dot2: opt.Something("ftp://host.example/nodeinfo/2.2.json"),
				NodeInfo2Dot1: opt.Something("http://host.example/nodeinfo/2.1"),
				NodeInfo2:     opt.Something("https://host.example/nodeinfo/2.0"),
				NodeInfo1Dot1: opt.Something("pow://mastodon.social/nodeinfo/2.1"),
				NodeInfo1:     opt.Something("pow://mastodon.social/nodeinfo/2.1"),
			},
		},
		{
			WellKnown: WellKnown{
				NodeInfo2Dot2: opt.Something("ftp://host.example/nodeinfo/2.2.json"),
				NodeInfo2Dot1: opt.Something("http://host.example/nodeinfo/2.1"),
				NodeInfo2:     opt.Something("https://host.example/nodeinfo/2.0"),
				NodeInfo1Dot1: opt.Something("/nodeinfo/2.1"),
				NodeInfo1:     opt.Something("nodeinfo/2.1"),
			},
			Scheme: "http",
			Host: "mastodon.social",
			Expected: WellKnown{
				NodeInfo2Dot2: opt.Something("ftp://host.example/nodeinfo/2.2.json"),
				NodeInfo2Dot1: opt.Something("http://host.example/nodeinfo/2.1"),
				NodeInfo2:     opt.Something("https://host.example/nodeinfo/2.0"),
				NodeInfo1Dot1: opt.Something("http://mastodon.social/nodeinfo/2.1"),
				NodeInfo1:     opt.Something("http://mastodon.social/nodeinfo/2.1"),
			},
		},
		{
			WellKnown: WellKnown{
				NodeInfo2Dot2: opt.Something("ftp://host.example/nodeinfo/2.2.json"),
				NodeInfo2Dot1: opt.Something("http://host.example/nodeinfo/2.1"),
				NodeInfo2:     opt.Something("https://host.example/nodeinfo/2.0"),
				NodeInfo1Dot1: opt.Something("/nodeinfo/2.1"),
				NodeInfo1:     opt.Something("nodeinfo/2.1"),
			},
			Scheme: "https",
			Host: "mastodon.social",
			Expected: WellKnown{
				NodeInfo2Dot2: opt.Something("ftp://host.example/nodeinfo/2.2.json"),
				NodeInfo2Dot1: opt.Something("http://host.example/nodeinfo/2.1"),
				NodeInfo2:     opt.Something("https://host.example/nodeinfo/2.0"),
				NodeInfo1Dot1: opt.Something("https://mastodon.social/nodeinfo/2.1"),
				NodeInfo1:     opt.Something("https://mastodon.social/nodeinfo/2.1"),
			},
		},
	}

	for testNumber, test := range tests {

		actual := resolveWellKnown(test.WellKnown, test.Scheme, test.Host)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual resolved-reference is not what was expected.", testNumber)
			t.Logf("EXPECTED:\n%#v", expected)
			t.Logf("ACTUAL:\n%#v", actual)
			t.Logf("WELL-KNOWN: %#v", test.WellKnown)
			t.Logf("SCHEME: %q", test.Host)
			t.Logf("HOST:   %q", test.Scheme)
			continue
		}
	}
}
