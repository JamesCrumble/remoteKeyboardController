package keyboardController

import (
	"fmt"

	keycode "github.com/vcaesar/keycode"
)

var ruToEng = map[string]string{
	"й": "q",
	"ц": "w",
	"у": "e",
	"к": "r",
	"е": "t",
	"н": "y",
	"г": "u",
	"ш": "i",
	"щ": "o",
	"з": "p",
	"х": "[",
	"ъ": "]",
	"ф": "a",
	"ы": "s",
	"в": "d",
	"а": "f",
	"п": "g",
	"р": "h",
	"о": "j",
	"л": "k",
	"д": "l",
	"ж": ";",
	"э": "'",
	"я": "z",
	"ч": "x",
	"с": "c",
	"м": "v",
	"и": "b",
	"т": "n",
	"ь": "m",
	"б": ",",
	"ю": ".",
}

func ToEng(char string) string {
	engKey, ok := ruToEng[char]
	if !ok {
		return char
	}
	return engKey
}

func CharToKeycode(char string) (uint16, error) {
	keycode, ok := keycode.Keycode[char]
	if !ok {
		return 0, fmt.Errorf("UNKNOWN \"%s\" CHAR", char)
	}
	return keycode, nil
}
