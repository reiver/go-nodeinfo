package internal

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/reiver/go-nodeinfo/shared"
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

	var digest [sha256.Size]byte
	{
		digest = sha256.Sum256(bytes)
	}

	var cacheDigest string
	{
		cacheDigest = fmt.Sprintf("sha-256=:%s:", base64.StdEncoding.EncodeToString(digest[:]))
	}

	var eTag string
	{
		var format string = fmt.Sprintf("sha256=0x%%0%dX", sha256.Size*2)
		eTag = fmt.Sprintf(format, digest[:])
	}

	{
		var header http.Header = responseWriter.Header()
		if nil == header {
			var code int = http.StatusInternalServerError
			var text string = http.StatusText(code)

			http.Error(responseWriter, text, code)
			return
		}

		header.Add("Access-Control-Allow-Origin", "*")
		header.Add("Cache-Control", "max-age=907")
		header.Add("Content-Digest", cacheDigest)
		header.Add("Content-Type", "application/json; charset=utf-8")
		header.Add("ETag", `"`+eTag+`"`)
		if servertext := shared.ServerText; "" != servertext {
			header.Add("Server", servertext)
		}
	}

	{
		var header http.Header = request.Header
		if nil != header {
			var ifNoneMatch string = header.Get("If-None-Match")
			if eTag == ifNoneMatch {
				responseWriter.WriteHeader(http.StatusNotModified)
				return
			}
		}
	}

	{
		responseWriter.WriteHeader(http.StatusOK)

		responseWriter.Write(bytes)
	}
}
