package nodeinfo

import (
	"github.com/reiver/go-erorr"
	"github.com/reiver/go-jsonint"
	"github.com/reiver/go-opt"
)

type Usage struct {
	Users Users
	LocalPosts opt.Optional[string]
	LocalComments opt.Optional[string]
}

func (receiver Usage) MarshalJSON() ([]byte, error) {
	var buffer [256]byte
	var bytes []byte = buffer[0:0]

	bytes = append(bytes, '{')

	{
		const prefix string = `"users":`
		bytes = append(bytes, prefix...)

		p, err := receiver.Users.MarshalJSON()
		if nil != err {
			return nil, err
		}

		bytes = append(bytes, p...)
	}

	{
		const prefix string = `"localPosts":`

		err := func() error {
			value, found :=  receiver.LocalPosts.Get()
			if !found {
				return nil
			}

			if !jsonint.IsNumericString(value) {
				return erorr.Errorf("nodeinfo: usage.localPosts value %q is not an integer", value)
			}

			bytes = append(bytes, ',')

			bytes = append(bytes, prefix...)
			bytes = append(bytes, value...)

			return nil
		}()
		if nil != err {
			return nil, err
		}

	}

	{
		const prefix string = `"localComments":`

		err := func() error {
			value, found :=  receiver.LocalComments.Get()
			if !found {
				return nil
			}

			if !jsonint.IsNumericString(value) {
				return erorr.Errorf("nodeinfo: usage.localComments value %q is not an integer", value)
			}

			bytes = append(bytes, ',')

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
