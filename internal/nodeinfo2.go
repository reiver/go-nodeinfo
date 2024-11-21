package internal

import (
	"encoding/json"

	"github.com/reiver/go-erorr"

	"github.com/reiver/go-nodeinfo/shared"
)

func NodeInfo2MarshalJSON(version string, software shared.Software1, protocols []string, services shared.Services, openRegistrations bool, usage shared.Usage, metadata map[string]any) ([]byte, error) {
	var buffer [256]byte
	var bytes []byte = buffer[0:0]

	bytes = append(bytes, '{')

	{
		const field string = `"version":`
		bytes = append(bytes, field...)


		encoded, err := json.Marshal(version)
		if nil != err {
			return nil, erorr.Errorf("nodeinfo: problem encoding nodeinfo version: %w", err)
		}

		bytes = append(bytes, encoded...)
	}

	bytes = append(bytes, ',')

	{
		const prefix string = `"software":`

		bytes = append(bytes, prefix...)

		value, err := json.Marshal(software)
		if nil != err {
			return nil, err
		}

		bytes = append(bytes, value...)
	}

	bytes = append(bytes, ',')

	{
		const prefix string = `"protocols":`
		bytes = append(bytes, prefix...)

		if len(protocols) <= 0 {
			bytes = append(bytes, `[]`...)
		} else {
			value, err := json.Marshal(protocols)
			if nil != err {
				return nil, err
			}

			bytes = append(bytes, value...)
		}
	}

	bytes = append(bytes, ',')

	{
		const prefix string = `"services":`

		bytes = append(bytes, prefix...)

		value, err := json.Marshal(services)
		if nil != err {
			return nil, err
		}

		bytes = append(bytes, value...)
	}

	bytes = append(bytes, ',')

	{
		const prefix string = `"openRegistrations":`
		var data bool = openRegistrations

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

		value, err := json.Marshal(usage)
		if nil != err {
			return nil, err
		}

		bytes = append(bytes, value...)
	}

	bytes = append(bytes, ',')

	{
		const prefix string = `"metadata":`
		var data map[string]any = metadata

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
