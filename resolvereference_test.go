package nodeinfo

import (
	"testing"
)

func TestResolveReference(t *testing.T) {

	tests := []struct{
		Reference string
		Scheme string
		Host string
		Expected string
	}{
		{
			Reference: "/.well-known/nodeinfo/2.0",
			Host: "example.com",
			Expected: "//example.com/.well-known/nodeinfo/2.0",
		},
		{
			Reference: "/.well-known/nodeinfo/2.0",
			Scheme: "http",
			Host: "example.com",
			Expected: "http://example.com/.well-known/nodeinfo/2.0",
		},
		{
			Reference: "/.well-known/nodeinfo/2.0",
			Scheme: "https",
			Host: "example.com",
			Expected: "https://example.com/.well-known/nodeinfo/2.0",
		},



		{
			Reference: "/nodeinfo/2.0",
			Host: "example.com",
			Expected: "//example.com/nodeinfo/2.0",
		},
		{
			Reference: "/nodeinfo/2.0",
			Scheme: "http",
			Host: "example.com",
			Expected: "http://example.com/nodeinfo/2.0",
		},
		{
			Reference: "/nodeinfo/2.0",
			Scheme: "https",
			Host: "example.com",
			Expected: "https://example.com/nodeinfo/2.0",
		},



		{
			Reference: "nodeinfo/2.0",
			Host: "example.com",
			Expected: "//example.com/nodeinfo/2.0",
		},
		{
			Reference: "nodeinfo/2.0",
			Scheme: "http",
			Host: "example.com",
			Expected: "http://example.com/nodeinfo/2.0",
		},
		{
			Reference: "nodeinfo/2.0",
			Scheme: "https",
			Host: "example.com",
			Expected: "https://example.com/nodeinfo/2.0",
		},



		{
			Reference: "http://mastodon.social/nodeinfo/2.0",
			Host: "example.com",
			Expected:  "http://mastodon.social/nodeinfo/2.0",
		},
		{
			Reference: "http://mastodon.social/nodeinfo/2.0",
			Scheme: "http",
			Host: "example.com",
			Expected:  "http://mastodon.social/nodeinfo/2.0",
		},
		{
			Reference: "http://mastodon.social/nodeinfo/2.0",
			Scheme: "https",
			Host: "example.com",
			Expected:  "http://mastodon.social/nodeinfo/2.0",
		},



		{
			Reference: "https://mastodon.social/nodeinfo/2.0",
			Host: "example.com",
			Expected:  "https://mastodon.social/nodeinfo/2.0",
		},
		{
			Reference: "https://mastodon.social/nodeinfo/2.0",
			Scheme: "http",
			Host: "example.com",
			Expected:  "https://mastodon.social/nodeinfo/2.0",
		},
		{
			Reference: "https://mastodon.social/nodeinfo/2.0",
			Scheme: "https",
			Host: "example.com",
			Expected:  "https://mastodon.social/nodeinfo/2.0",
		},
	}

	for testNumber, test := range tests {

		actual, err := resolveReference(test.Reference, test.Scheme, test.Host)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("REFERENCE: %q", test.Reference)
			t.Logf("SCHEME: %q", test.Host)
			t.Logf("HOST:   %q", test.Scheme)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual resolved-reference is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("REFERENCE: %q", test.Reference)
			t.Logf("SCHEME: %q", test.Host)
			t.Logf("HOST:   %q", test.Scheme)
			continue
		}
	}
}
