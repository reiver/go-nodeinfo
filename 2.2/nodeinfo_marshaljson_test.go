package nodeinfo_test

import (
	"testing"

	"bytes"
	"encoding/json"

	"github.com/reiver/go-opt"

	"github.com/reiver/go-nodeinfo/2.2"
)

func TestNodeInfo_MarshalJSON(t *testing.T) {

	tests := []struct{
		NodeInfo nodeinfo.NodeInfo
		Expected []byte
	}{
		{
			Expected: []byte(`{`+
				`"version":"2.2"`+
				`,`+
				`"instance":{`+
				`}`+
				`,`+
				`"software":{`+
					`"name":""`+
					`,`+
					`"version":""`+
				`}`+
				`,`+
				`"protocols":[]`+
				`,`+
				`"services":{"inbound":[],"outbound":[]}`+
				`,`+
				`"openRegistrations":false`+
				`,`+
				`"usage":{"users":{}}`+
				`,`+
				`"metadata":{}`+
			`}`),
		},



		{
			NodeInfo: nodeinfo.NodeInfo{
				Instance: nodeinfo.Instance{
					Name: opt.Something("FI𝑓F fediverse"),
					Description: opt.Something("Welcome to the FI𝑓F fediverse server"),
				},
				Software: nodeinfo.Software{
					Name: "diaspora",
					Version: "0.5.0",
					Repository: opt.Something("https://github.com/diaspora/diaspora"),
					HomePage: opt.Something("https://diasporafoundation.org/"),
				},
				Protocols: []string{"diaspora"},
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
						ActiveWeek: opt.Something("22"),
					},
					LocalPosts:    opt.Something("500"),
					LocalComments: opt.Something("1000"),
				},
				MetaData: map[string]any{
					"chat_enabled": true,
				},
			},
			Expected: []byte(`{`+
				`"version":"2.2"`+
				`,`+
				`"instance":{`+
					`"name":"FI𝑓F fediverse"`+
					`,`+
					`"description":"Welcome to the FI𝑓F fediverse server"`+
				`}`+
				`,`+
				`"software":{`+
					`"name":"diaspora"`+
					`,`+
					`"version":"0.5.0"`+
					`,`+
					`"repository":"https://github.com/diaspora/diaspora"`+
					`,`+
					`"homepage":"https://diasporafoundation.org/"`+
				`}`+
				`,`+
				`"protocols":["diaspora"]`+
				`,`+
				`"services":{"inbound":["gnusocial"],"outbound":["facebook","twitter"]}`+
				`,`+
				`"openRegistrations":true`+
				`,`+
				`"usage":{"users":{"total":123,"activeHalfyear":42,"activeMonth":23,"activeWeek":22},"localPosts":500,"localComments":1000}`+
				`,`+
				`"metadata":{"chat_enabled":true}`+
			`}`),
		},
	}

	for testNumber, test := range tests {

		actual, err := json.Marshal(test.NodeInfo)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		{
			expected := test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual json-marshaled nodeinfo is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				continue
			}
		}
	}
}
