package nodeinfo_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-opt"

	"github.com/reiver/go-nodeinfo/1.0"
)

func TestUsers_MarshalJSON(t *testing.T) {

	tests := []struct{
		Users nodeinfo.Users
		Expected []byte
	}{
		{
			Expected: []byte(`{}`),
		},









		{
			Users: nodeinfo.Users{
				Total: opt.Something("0"),
			},
			Expected: []byte(`{"total":0}`),
		},
		{
			Users: nodeinfo.Users{
				Total: opt.Something("1"),
			},
			Expected: []byte(`{"total":1}`),
		},
		{
			Users: nodeinfo.Users{
				Total: opt.Something("2"),
			},
			Expected: []byte(`{"total":2}`),
		},

		{
			Users: nodeinfo.Users{
				Total: opt.Something("12345"),
			},
			Expected: []byte(`{"total":12345}`),
		},

		{
			Users: nodeinfo.Users{
				Total: opt.Something("4294967295"),
			},
			Expected: []byte(`{"total":4294967295}`),
		},

		{
			Users: nodeinfo.Users{
				Total: opt.Something("90817263544536271809"),
			},
			Expected: []byte(`{"total":90817263544536271809}`),
		},









		{
			Users: nodeinfo.Users{
				ActiveHalfYear: opt.Something("0"),
			},
			Expected: []byte(`{"activeHalfyear":0}`),
		},
		{
			Users: nodeinfo.Users{
				ActiveHalfYear: opt.Something("1"),
			},
			Expected: []byte(`{"activeHalfyear":1}`),
		},
		{
			Users: nodeinfo.Users{
				ActiveHalfYear: opt.Something("2"),
			},
			Expected: []byte(`{"activeHalfyear":2}`),
		},
		{
			Users: nodeinfo.Users{
				ActiveHalfYear: opt.Something("12345"),
			},
			Expected: []byte(`{"activeHalfyear":12345}`),
		},

		{
			Users: nodeinfo.Users{
				ActiveHalfYear: opt.Something("4294967295"),
			},
			Expected: []byte(`{"activeHalfyear":4294967295}`),
		},

		{
			Users: nodeinfo.Users{
				ActiveHalfYear: opt.Something("90817263544536271809"),
			},
			Expected: []byte(`{"activeHalfyear":90817263544536271809}`),
		},









		{
			Users: nodeinfo.Users{
				ActiveMonth: opt.Something("0"),
			},
			Expected: []byte(`{"activeMonth":0}`),
		},
		{
			Users: nodeinfo.Users{
				ActiveMonth: opt.Something("1"),
			},
			Expected: []byte(`{"activeMonth":1}`),
		},
		{
			Users: nodeinfo.Users{
				ActiveMonth: opt.Something("2"),
			},
			Expected: []byte(`{"activeMonth":2}`),
		},

		{
			Users: nodeinfo.Users{
				ActiveMonth: opt.Something("12345"),
			},
			Expected: []byte(`{"activeMonth":12345}`),
		},

		{
			Users: nodeinfo.Users{
				ActiveMonth: opt.Something("4294967295"),
			},
			Expected: []byte(`{"activeMonth":4294967295}`),
		},

		{
			Users: nodeinfo.Users{
				ActiveMonth: opt.Something("90817263544536271809"),
			},
			Expected: []byte(`{"activeMonth":90817263544536271809}`),
		},









		{
			Users: nodeinfo.Users{
				Total: opt.Something("1234512345123451234512345"),
				ActiveHalfYear: opt.Something("1111122222333334444455555"),
			},
			Expected: []byte(`{"total":1234512345123451234512345,"activeHalfyear":1111122222333334444455555}`),
		},
		{
			Users: nodeinfo.Users{
				Total: opt.Something("1234512345123451234512345"),
				ActiveMonth: opt.Something("1111222233334444"),
			},
			Expected: []byte(`{"total":1234512345123451234512345,"activeMonth":1111222233334444}`),
		},
		{
			Users: nodeinfo.Users{
				ActiveHalfYear: opt.Something("1111122222333334444455555"),
				ActiveMonth: opt.Something("1111222233334444"),
			},
			Expected: []byte(`{"activeHalfyear":1111122222333334444455555,"activeMonth":1111222233334444}`),
		},









		{
			Users: nodeinfo.Users{
				Total: opt.Something("1234512345123451234512345"),
				ActiveHalfYear: opt.Something("1111122222333334444455555"),
				ActiveMonth: opt.Something("1111222233334444"),
			},
			Expected: []byte(`{"total":1234512345123451234512345,"activeHalfyear":1111122222333334444455555,"activeMonth":1111222233334444}`),
		},
	}

	for testNumber, test := range tests {

		actual, err := test.Users.MarshalJSON()

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
