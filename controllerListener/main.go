package main

import (
	server "controllerListener/registerServer"
	s "controllerListener/settings"
	"fmt"
	hook "github.com/robotn/gohook"
)

func main() {
	settings := s.Settings()
	buffer := server.NewBuffer(128)
	serverListener := server.CreateListener(settings.RegisterServer.Host, uint16(settings.RegisterServer.Port))
	fmt.Printf("SOCKET LISTENING ON %s:%d\n", settings.RegisterServer.Host, settings.RegisterServer.Port)

	go server.Run(&serverListener, &buffer)

	eventsChannel := hook.Start()
	defer hook.End()

	for event := range eventsChannel {

		if event.Kind == hook.KeyDown {
			if settings.LogPressedChars {
				fmt.Printf("PRESSED: CHAR \"%q\", KEYCODE \"%v\"\n", event.Keychar, event.Rawcode)
			}
			sendPressActions(event.Keychar)
		}
	}
}
