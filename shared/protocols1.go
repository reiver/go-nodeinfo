package shared

import (
	"encoding/json"
)

type Protocols1 struct {
	Inbound  []string
	Outbound []string
}

func (receiver Protocols1) MarshalJSON() ([]byte, error) {

	var buffer [256]byte
	var bytes []byte = buffer[0:0]

	bytes = append(bytes, '{')

	{
		const prefix string = `"inbound":`
		var data []string = receiver.Inbound

		bytes = append(bytes, prefix...)

		if len(data) <= 0 {
			bytes = append(bytes, `[]`...)
		} else {
			value, err := json.Marshal(data)
			if nil != err {
				return nil, err
			}

			bytes = append(bytes, value...)
		}
	}

	bytes = append(bytes, ',')

	{
		const prefix string = `"outbound":`
		var data []string = receiver.Outbound

		bytes = append(bytes, prefix...)

		if len(data) <= 0 {
			bytes = append(bytes, `[]`...)
		} else {
			value, err := json.Marshal(data)
			if nil != err {
				return nil, err
			}

			bytes = append(bytes, value...)
		}
	}

	bytes = append(bytes, '}')

	return bytes, nil
}
