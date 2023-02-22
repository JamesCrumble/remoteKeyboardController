package acceptingServer

import (
	keyboardController "client/keyboardController"
	settings "client/settings"
	"errors"
	"fmt"
)

/*
COMMANDS:
  - press;char
*/

func pressCommand(char string) {
	if settings.Settings().AcceptableButtons.Contains(char) {
		keyboardController.Press(char)
	}

	keycode, err := keyboardController.CharToKeycode(char)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("GET \"%d\" KEYCODE FOR \"%s\" CHAR\n", keycode, char)
}

func defineAndProcessCommand(command *string, params *[]string) error {
	if *command == "press" {
		if len(*params) != 1 {
			return errors.New("\"press\" COMMAND EXPECTED 1 ARGUMENT ((string)char)")
		}
		if len((*params)[0]) != 1 {
			return fmt.Errorf("\"press\" COMMAND EXPECT SINGLE CHAR STRING BUT GET => \"%s\"", (*params)[0])
		}
		pressCommand((*params)[0])
	}
	return nil
}
