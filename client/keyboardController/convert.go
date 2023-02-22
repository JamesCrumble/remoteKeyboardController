package keyboardController

import (
	"fmt"

	keycode "github.com/vcaesar/keycode"
)

func CharToKeycode(char string) (uint16, error) {
	keycode, ok := keycode.Keycode[char]
	if !ok {
		return 0, fmt.Errorf("UNKNOWN \"%s\" CHAR", char)
	}
	return keycode, nil
}
