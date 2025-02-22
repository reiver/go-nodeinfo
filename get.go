package nodeinfo

import (
	"io"
	"net/http"
	gourl "net/url"

	"github.com/reiver/go-erorr"
)

func GetHost(host string) (WellKnown, error) {
	var url gourl.URL
	url.Scheme = "https"
	url.Host = host
	url.Path = DefaultPath

	return getURL(&url)
}

func GetURL(url string) (WellKnown, error) {

	urloc, err := gourl.Parse(url)
	if nil != err {
		var nada WellKnown
		return nada, erorr.Errorf("nodeinfo: problem parsing URL %q: %w", url, err)
	}

	return getURL(urloc)
}

func getURL(url *gourl.URL) (WellKnown, error) {

	var httprequest http.Request
	{
		httprequest.Method = http.MethodGet
		httprequest.URL = url
		if nil == httprequest.Header {
			httprequest.Header = http.Header{}
		}
		httprequest.Header.Add("User-Agent", UserAgent)
	}

	var httpresponse *http.Response
	{
		var err error

		httpresponse, err = http.DefaultClient.Do(&httprequest)
		if nil != err {
			var nada WellKnown
			return nada, erorr.Errorf("nodeinfo: problem making HTTP-request to %q: %w", url.String(), err)
		}
	}

	var bytes []byte
	{
		var err error

		bytes, err = io.ReadAll(httpresponse.Body)
		if nil != err {
			var nada WellKnown
			return nada, erorr.Errorf("nodeinfo: problem get data from body of HTTP-response from %q: %w", url.String(), err)
		}
	}

	var wellknown WellKnown
	{
		err := wellknown.UnmarshalJSON(bytes)
		if nil != err {
			var nada WellKnown
			return nada, erorr.Errorf("nodeinfo: problem json-unmarshaling the data from the body of the HTTP-response from %q: %w", url.String(), err)
		}
	}

	return wellknown, nil
}
