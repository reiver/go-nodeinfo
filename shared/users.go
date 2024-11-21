package shared

import (
	"github.com/reiver/go-erorr"
	"github.com/reiver/go-jsonint"
	"github.com/reiver/go-opt"
)

type Users struct {
	Total          opt.Optional[string]
	ActiveHalfYear opt.Optional[string]
	ActiveMonth    opt.Optional[string]
}

func (receiver Users) MarshalJSON() ([]byte, error) {
	var buffer [256]byte
	var bytes []byte = buffer[0:0]

	var comma bool = false

	bytes = append(bytes, '{')

	{
		const prefix string =           `"total":`

		err := func() error {
			value, found :=  receiver.Total.Get()
			if !found {
				return nil
			}

			if !jsonint.IsNumericString(value) {
				return erorr.Errorf("nodeinfo: usage.users.total value %q is not an integer", value)
			}

			if comma {
				bytes = append(bytes, ',')
			}
			comma = true

			bytes = append(bytes, prefix...)
			bytes = append(bytes, value...)

			return nil
		}()
		if nil != err {
			return nil, err
		}
	}

	{
		const prefix string =           `"activeHalfyear":`

		err := func() error {
			value, found :=  receiver.ActiveHalfYear.Get()
			if !found {
				return nil
			}

			if !jsonint.IsNumericString(value) {
				return erorr.Errorf("nodeinfo: usage.users.activeHalfyear value %q is not an integer", value)
			}

			if comma {
				bytes = append(bytes, ',')
			}
			comma = true

			bytes = append(bytes, prefix...)
			bytes = append(bytes, value...)

			return nil
		}()
		if nil != err {
			return nil, err
		}
	}

	{
		const prefix string =           `"activeMonth":`

		err := func() error {
			value, found :=  receiver.ActiveMonth.Get()
			if !found {
				return nil
			}

			if !jsonint.IsNumericString(value) {
				return erorr.Errorf("nodeinfo: usage.users.activeMonth value %q is not an integer", value)
			}

			if comma {
				bytes = append(bytes, ',')
			}
			comma = true

			bytes = append(bytes, prefix...)
			bytes = append(bytes, value...)

			return nil
		}()
		if nil != err {
			return nil, err
		}
	}

	bytes = append(bytes, '}')

	return bytes, nil
}
