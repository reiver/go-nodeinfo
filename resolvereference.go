package nodeinfo

import (
	liburl "net/url"

	"github.com/reiver/go-erorr"
)

func resolveReference(ref string, scheme string, host string) (string, error) {

	var base = liburl.URL{
		Scheme: scheme,
		Host: host,
	}

	reference, err := liburl.Parse(ref)
	if nil != err {
		var nada string
		return nada, erorr.Errorf("nodeinfo: problem parsing %q as a URI: %w", ref, err)
	}

	resolvedRef := base.ResolveReference(reference)
	if nil == resolvedRef {
		var nada string
		return nada, erorr.Errorf("nodeinfo: nil resolved-reference — resoling %q against base %q", ref, base.String())
	}

	return resolvedRef.String(), nil
}
