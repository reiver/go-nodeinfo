package nodeinfo

import (
	"net/http"

	"github.com/reiver/go-http500"
)

type ResolvingWellKnown struct {
	WellKnown WellKnown
	Scheme string
}

var _ http.Handler = ResolvingWellKnown{}

func (receiver ResolvingWellKnown) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {

	if nil == responseWriter {
		return
	}
	if nil == request {
		http500.InternalServerError(responseWriter, request)
		return
	}

	var host string = request.Host

	const defaultScheme string = "https"
	var scheme string = defaultScheme
	if "" != receiver.Scheme {
		scheme = receiver.Scheme
	}

	var wellKnown WellKnown = resolveWellKnown(receiver.WellKnown, scheme, host)

	wellKnown.ServeHTTP(responseWriter, request)
}
