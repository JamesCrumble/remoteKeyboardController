package helpers

import (
	"fmt"
	"strings"
	"time"
)

const PANIC_DELAY = time.Second * 5

func Apply[T comparable](iterable *[]T, key func(T) string) []string {
	var arr []string

	for _, element := range *iterable {
		arr = append(arr, key(element))
	}

	return arr
}

func Contains[T comparable](arr *[]T, value T) bool {
	for _, element := range *arr {
		if element != value {
			continue
		}
		return true
	}
	return false
}

func IpAddressFromNetInterfaceData(netInterfaceData *[]string) string {
	var ipAddress string = (*netInterfaceData)[1]
	var index = strings.Index(ipAddress, "/")

	if index == -1 {
		return ipAddress
	}
	return ipAddress[:index]
}

func DelayedPanic(err any) {
	fmt.Println(err)
	time.Sleep(PANIC_DELAY)
	panic(err)
}
