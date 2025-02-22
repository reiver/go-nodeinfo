package nodeinfo

import (
	"encoding/json"
	"net/http"

	"github.com/reiver/go-erorr"
	"github.com/reiver/go-opt"

	"github.com/reiver/go-nodeinfo/internal"
)

type WellKnown struct {
	NodeInfo2Dot2 opt.Optional[string]
	NodeInfo2Dot1 opt.Optional[string]
	NodeInfo2     opt.Optional[string]
	NodeInfo1Dot1 opt.Optional[string]
	NodeInfo1     opt.Optional[string]
}

var _ json.Marshaler = WellKnown{}
var _ http.Handler = WellKnown{}

func (receiver WellKnown) MarshalJSON() ([]byte, error) {
	var buffer [256]byte
	var bytes []byte = buffer[0:0]

	bytes = append(bytes, `{"links":[`...)

	var comma bool = false

	{
		err := func() error {
			const rel string = "http://nodeinfo.diaspora.software/ns/schema/2.2"

			href, found := receiver.NodeInfo2Dot2.Get()
			if !found {
				return nil
			}

			if comma {
				bytes = append(bytes, ',')
			}

			{
				var link = internalLink{
					Rel:  rel,
					HRef: href,
				}

				value, err := json.Marshal(link)
				if nil != err {
					return err
				}

				bytes = append(bytes, value...)
			}

			comma = true
			return nil
		}()
		if nil != err {
			return nil, err
		}
	}

	{
		const rel string = "http://nodeinfo.diaspora.software/ns/schema/2.1"

		err := func() error {
			href, found := receiver.NodeInfo2Dot1.Get()
			if !found {
				return nil
			}

			if comma {
				bytes = append(bytes, ',')
			}

			{
				var link = internalLink{
					Rel:  rel,
					HRef: href,
				}

				value, err := json.Marshal(link)
				if nil != err {
					return err
				}

				bytes = append(bytes, value...)
			}

			comma = true
			return nil
		}()
		if nil != err {
			return nil, err
		}
	}

	{
		const rel string = "http://nodeinfo.diaspora.software/ns/schema/2.0"

		err := func() error {
			href, found := receiver.NodeInfo2.Get()
			if !found {
				return nil
			}

			if comma {
				bytes = append(bytes, ',')
			}

			{
				var link = internalLink{
					Rel:  rel,
					HRef: href,
				}

				value, err := json.Marshal(link)
				if nil != err {
					return err
				}

				bytes = append(bytes, value...)
			}

			comma = true
			return nil
		}()
		if nil != err {
			return nil, err
		}
	}

	{
		const rel string = "http://nodeinfo.diaspora.software/ns/schema/1.1"

		err := func() error {
			href, found := receiver.NodeInfo1Dot1.Get()
			if !found {
				return nil
			}

			if comma {
				bytes = append(bytes, ',')
			}

			{
				var link = internalLink{
					Rel:  rel,
					HRef: href,
				}

				value, err := json.Marshal(link)
				if nil != err {
					return err
				}

				bytes = append(bytes, value...)
			}

			comma = true
			return nil
		}()
		if nil != err {
			return nil, err
		}
	}

	{
		const rel string = "http://nodeinfo.diaspora.software/ns/schema/1.0"

		err := func() error {
			href, found := receiver.NodeInfo1.Get()
			if !found {
				return nil
			}

			if comma {
				bytes = append(bytes, ',')
			}

			{
				var link = internalLink{
					Rel:  rel,
					HRef: href,
				}

				value, err := json.Marshal(link)
				if nil != err {
					return err
				}

				bytes = append(bytes, value...)
			}

			comma = true
			return nil
		}()
		if nil != err {
			return nil, err
		}
	}

	bytes = append(bytes, `]}`...)

	return bytes, nil
}

func (receiver *WellKnown) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return errNilReceiver
	}
	if nil == bytes {
		return errNilBytes
	}

	type internalWellKnown struct {
		Links []internalLink `json:"links"`
	}

	var data internalWellKnown
	{
		err := json.Unmarshal(bytes, &data)
		if nil != err {
			return erorr.Errorf("nodeinfo: problem json-unmarshaling data from NodeInfo well-known API end-point: %w", err)
		}
	}

	{
		var nada WellKnown
		*receiver = nada

		for _, link := range data.Links {
			switch link.Rel {
			case "http://nodeinfo.diaspora.software/ns/schema/1.0":
				if "" != link.HRef {
					receiver.NodeInfo1 = opt.Something(link.HRef)
				}
			case "http://nodeinfo.diaspora.software/ns/schema/1.1":
				if "" != link.HRef {
					receiver.NodeInfo1Dot1 = opt.Something(link.HRef)
				}
			case "http://nodeinfo.diaspora.software/ns/schema/2.0":
				if "" != link.HRef {
					receiver.NodeInfo2 = opt.Something(link.HRef)
				}
			case "http://nodeinfo.diaspora.software/ns/schema/2.1":
				if "" != link.HRef {
					receiver.NodeInfo2Dot1 = opt.Something(link.HRef)
				}
			case "http://nodeinfo.diaspora.software/ns/schema/2.2":
				if "" != link.HRef {
					receiver.NodeInfo2Dot2 = opt.Something(link.HRef)
				}
			}
		}
	}

	return nil
}

func (receiver WellKnown) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	internal.JSONServeHTTP(responseWriter, request, receiver)
}
