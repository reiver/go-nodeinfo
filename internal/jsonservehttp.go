package internal

import (
	"encoding/json"
	"net/http"
)

func JSONServeHTTP(responseWriter http.ResponseWriter, request *http.Request, jsonMarshaler json.Marshaler) {
	if nil == responseWriter {
		return
	}
	if nil == request {
		var code int = http.StatusInternalServerError
		var text string = http.StatusText(code)

		http.Error(responseWriter, text, code)
		return
	}
	if nil == jsonMarshaler {
		var code int = http.StatusInternalServerError
		var text string = http.StatusText(code)

		http.Error(responseWriter, text, code)
		return
	}

	var bytes []byte
	{
		var err error

		bytes, err = jsonMarshaler.MarshalJSON()
		if nil != err {
			var code int = http.StatusInternalServerError
			var text string = http.StatusText(code)

			http.Error(responseWriter, text, code)
			return
		}
	}

	{
		responseWriter.Write(bytes)
	}
}
