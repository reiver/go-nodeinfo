package shared

import (
	"encoding/json"

	"github.com/reiver/go-opt"
)

type Instance2Dot2 struct {
	Name        opt.Optional[string]
	Description opt.Optional[string]
}

func (receiver Instance2Dot2) MarshalJSON() ([]byte, error) {
	var buffer [256]byte
	var bytes []byte = buffer[0:0]

	var comma bool = false

	bytes = append(bytes, '{')

	{
		const prefix string =           `"name":`

		err := func() error {
			data, found :=  receiver.Name.Get()
			if !found {
				return nil
			}

			comma = true

			bytes = append(bytes, prefix...)

			{
				value, err := json.Marshal(data)
				if nil != err {
					return err
				}

				bytes = append(bytes, value...)
			}

			return nil
		}()
		if nil != err {
			return nil, err
		}
	}

	{
		const prefix string =           `"description":`

		err := func() error {
			data, found :=  receiver.Description.Get()
			if !found {
				return nil
			}

			if comma {
				bytes = append(bytes, ',')
			}
			comma = true

			bytes = append(bytes, prefix...)

			{
				value, err := json.Marshal(data)
				if nil != err {
					return err
				}

				bytes = append(bytes, value...)
			}

			return nil
		}()
		if nil != err {
			return nil, err
		}
	}

	bytes = append(bytes, '}')

	return bytes, nil
}
