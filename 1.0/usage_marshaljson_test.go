package nodeinfo_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-opt"

	"github.com/reiver/go-nodeinfo/1.0"
)

func TestUsage_MarshalJSON(t *testing.T) {

	tests := []struct{
		Usage nodeinfo.Usage
		Expected []byte
	}{
		{
			Usage: nodeinfo.Usage{
				Users: nodeinfo.Users{
					Total: opt.Something("1234512345123451234512345"),
					ActiveHalfYear: opt.Something("1111122222333334444455555"),
					ActiveMonth: opt.Something("1111222233334444"),
				},
			},
			Expected: []byte(`{"users":{"total":1234512345123451234512345,"activeHalfyear":1111122222333334444455555,"activeMonth":1111222233334444}}`),
		},



		{
			Usage: nodeinfo.Usage{
				Users: nodeinfo.Users{
					Total: opt.Something("1234512345123451234512345"),
					ActiveHalfYear: opt.Something("1111122222333334444455555"),
					ActiveMonth: opt.Something("1111222233334444"),
				},
				LocalPosts: opt.Something("90817263544536271809"),
			},
			Expected: []byte(`{`+
				`"users":{"total":1234512345123451234512345,"activeHalfyear":1111122222333334444455555,"activeMonth":1111222233334444}`+
				`,`+
				`"localPosts":90817263544536271809`+
			`}`),
		},



		{
			Usage: nodeinfo.Usage{
				Users: nodeinfo.Users{
					Total: opt.Something("1234512345123451234512345"),
					ActiveHalfYear: opt.Something("1111122222333334444455555"),
					ActiveMonth: opt.Something("1111222233334444"),
				},
				LocalComments: opt.Something("123234345456567678789"),
			},
			Expected: []byte(`{`+
				`"users":{"total":1234512345123451234512345,"activeHalfyear":1111122222333334444455555,"activeMonth":1111222233334444}`+
				`,`+
				`"localComments":123234345456567678789`+
			`}`),
		},



		{
			Usage: nodeinfo.Usage{
				Users: nodeinfo.Users{
					Total: opt.Something("1234512345123451234512345"),
					ActiveHalfYear: opt.Something("1111122222333334444455555"),
					ActiveMonth: opt.Something("1111222233334444"),
				},
				LocalPosts: opt.Something("90817263544536271809"),
				LocalComments: opt.Something("123234345456567678789"),
			},
			Expected: []byte(`{`+
				`"users":{"total":1234512345123451234512345,"activeHalfyear":1111122222333334444455555,"activeMonth":1111222233334444}`+
				`,`+
				`"localPosts":90817263544536271809`+
				`,`+
				`"localComments":123234345456567678789`+
			`}`),
		},
	}

	for testNumber, test := range tests {

		actual, err := test.Usage.MarshalJSON()

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		{
			expected := test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual marshaled-users value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				continue
			}
		}
	}
}
