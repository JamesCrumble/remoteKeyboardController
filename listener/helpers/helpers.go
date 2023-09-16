package helpers

import (
	"fmt"
	"time"
)

const PANIC_DELAY = time.Second * 5

func DelayedPanic(err any) {
	fmt.Println(err)
	time.Sleep(PANIC_DELAY)
	panic(err)
}
