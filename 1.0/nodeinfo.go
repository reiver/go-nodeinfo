package nodeinfo

import (
	"encoding/json"
)

type NodeInfo struct {
	Software Software
	Protocols Protocols
	Services Services
	OpenRegistrations bool
	Usage Usage
	MetaData map[string]any
}

func (receiver NodeInfo) MarshalJSON() ([]byte, error) {
	var buffer [256]byte
	var bytes []byte = buffer[0:0]

	bytes = append(bytes, '{')

	{
		const field string = `"version":"1.0"`

		bytes = append(bytes, field...)
	}

	bytes = append(bytes, ',')

	{
		const prefix string = `"software":`

		bytes = append(bytes, prefix...)

		value, err := json.Marshal(receiver.Software)
		if nil != err {
			return nil, err
		}

		bytes = append(bytes, value...)
	}

	bytes = append(bytes, ',')

	{
		const prefix string = `"protocols":`

		bytes = append(bytes, prefix...)

		value, err := json.Marshal(receiver.Protocols)
		if nil != err {
			return nil, err
		}

		bytes = append(bytes, value...)
	}

	bytes = append(bytes, ',')

	{
		const prefix string = `"services":`

		bytes = append(bytes, prefix...)

		value, err := json.Marshal(receiver.Services)
		if nil != err {
			return nil, err
		}

		bytes = append(bytes, value...)
	}

	bytes = append(bytes, ',')

	{
		const prefix string = `"openRegistrations":`
		var data bool = receiver.OpenRegistrations

		bytes = append(bytes, prefix...)

		if data {
			bytes = append(bytes, `true`...)
		} else {
			bytes = append(bytes, `false`...)
		}

	}

	bytes = append(bytes, ',')

	{
		const prefix string = `"usage":`

		bytes = append(bytes, prefix...)

		value, err := json.Marshal(receiver.Usage)
		if nil != err {
			return nil, err
		}

		bytes = append(bytes, value...)
	}

	bytes = append(bytes, ',')

	{
		const prefix string = `"metadata":`
		var data map[string]any = receiver.MetaData

		bytes = append(bytes, prefix...)

		if len(data) <= 0 {
			bytes = append(bytes, `{}`...)
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
