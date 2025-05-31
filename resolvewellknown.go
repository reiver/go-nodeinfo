package nodeinfo

import (
	"github.com/reiver/go-opt"
)

func resolveWellKnown(wellKnown WellKnown, scheme string, host string) WellKnown {

	var resolvedWellKnown WellKnown = wellKnown
	{
		resolvedWellKnown.NodeInfo2Dot2.WhenSomething(func(reference string){
			resolvedReference, err := resolveReference(reference, scheme, host)
			if nil == err {
				resolvedWellKnown.NodeInfo2Dot2 = opt.Something(resolvedReference)
			}
		})
		resolvedWellKnown.NodeInfo2Dot1.WhenSomething(func(reference string){
			resolvedReference, err := resolveReference(reference, scheme, host)
			if nil == err {
				resolvedWellKnown.NodeInfo2Dot1 = opt.Something(resolvedReference)
			}
		})
		resolvedWellKnown.NodeInfo2.WhenSomething(func(reference string){
			resolvedReference, err := resolveReference(reference, scheme, host)
			if nil == err {
				resolvedWellKnown.NodeInfo2 = opt.Something(resolvedReference)
			}
		})
		resolvedWellKnown.NodeInfo1Dot1.WhenSomething(func(reference string){
			resolvedReference, err := resolveReference(reference, scheme, host)
			if nil == err {
				resolvedWellKnown.NodeInfo1Dot1 = opt.Something(resolvedReference)
			}
		})
		resolvedWellKnown.NodeInfo1.WhenSomething(func(reference string){
			resolvedReference, err := resolveReference(reference, scheme, host)
			if nil == err {
				resolvedWellKnown.NodeInfo1 = opt.Something(resolvedReference)
			}
		})
	}

	return resolvedWellKnown
}
