package commands

import (
	"client/helpers"
	"client/keyboardController"
	"client/settings"
	"errors"
	"fmt"
	"unicode/utf8"
)

func validateCommandParameters(params []string) (string, error) {
	if len(params) != 1 {
		return "", errors.New("\"press\" command expects only one argument which as rune")
	}
	if utf8.RuneCountInString(params[0]) != 1 {
		return "", fmt.Errorf("\"press\" command expect single char string but get \"%s\"", params[0])
	}
	return params[0], nil
}

func pressCommand(params []string) {
	char, err := validateCommandParameters(params)
	if err != nil {
		fmt.Println(err)
	}

	char = keyboardController.ToEng(char)
	keycode, err := keyboardController.CharToKeycode(char)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("GET \"%d\" KEYCODE FOR \"%s\" CHAR\n", keycode, char)

	if !helpers.Contains(&settings.Settings().AcceptableButtons.Buttons, char) {
		return
	}
	keyboardController.Press(keycode)
}
