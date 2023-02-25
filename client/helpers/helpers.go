package helpers

import (
	"net"
)

func Apply(iterable []any, key func(any) any) []any {

	var applyedArray []any = make([]any, 8)

	for _, element := range iterable {
		applyedArray = append(applyedArray, key(element))
	}

	return applyedArray
}

func ApplyAddrToString(iterable []net.Addr, key func(net.Addr) string) []string {
	var applyedArray []string

	for _, element := range iterable {
		applyedArray = append(applyedArray, key(element))
	}

	return applyedArray
}
