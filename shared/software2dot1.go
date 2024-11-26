package shared

import (
	"encoding/json"

	"github.com/reiver/go-opt"
)

type Software2Dot1 struct {
	Name       string
	Version    string
	Repository opt.Optional[string]
	HomePage   opt.Optional[string]
}

func (receiver Software2Dot1) MarshalJSON() ([]byte, error) {
	var buffer [256]byte
	var bytes []byte = buffer[0:0]

	bytes = append(bytes, '{')

	{
		const prefix string = `"name":`
		var data string = receiver.Name

		bytes = append(bytes, prefix...)

		{
			value, err := json.Marshal(data)
			if nil != err {
				return nil, err
			}

			bytes = append(bytes, value...)
		}
	}

	bytes = append(bytes, ',')

	{
		const prefix string = `"version":`
		var data string = receiver.Version

		bytes = append(bytes, prefix...)

		{
			value, err := json.Marshal(data)
			if nil != err {
				return nil, err
			}

			bytes = append(bytes, value...)
		}
	}

	{
		const prefix string = `"repository":`

		err := func() error {
			data, found :=  receiver.Repository.Get()
			if !found {
                                return nil
			}

			bytes = append(bytes, ',')

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
		const prefix string = `"homepage":`

		err := func() error {
			data, found :=  receiver.HomePage.Get()
			if !found {
                                return nil
			}

			bytes = append(bytes, ',')

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
