package main

import (
	"fmt"
	hook "github.com/robotn/gohook"
	"listener/globals"
	"listener/server"
	s "listener/settings"
	"net"
	"time"
)

const ACTION_TIMEOUT = time.Second * 5

var settings s.SettingsStruct = *s.Settings()
var actionsNetDialer = net.Dialer{Timeout: ACTION_TIMEOUT}

func sendPressAction(client *globals.Client, char rune) {
	ip, port := (*client).Ip, (*client).Port

	conn, err := actionsNetDialer.Dial("tcp4", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		fmt.Println(err)
		globals.DeleteClient(client)
		return
	}
	defer conn.Close()

	action := fmt.Sprintf("%s;%c", "press", char)
	if _, err := conn.Write([]byte(action)); err != nil {
		fmt.Println(err)
	}
}

func sendPressActions(char rune) {
	clientsArray := *globals.Clients()

	for i := 0; i < len(clientsArray); i++ {
		if (clientsArray[i] == globals.Client{}) {
			continue
		}
		go sendPressAction(&clientsArray[i], char)
	}
}

func main() {
	buffer := server.NewBuffer(128)
	serverListener := server.CreateListener(settings.Server.Host, uint16(settings.Server.Port))
	fmt.Printf("Listeting socket on %s:%d\n", settings.Server.Host, settings.Server.Port)

	go server.InfiniteListening(&serverListener, &buffer)

	eventsChannel := hook.Start()
	defer hook.End()

	for event := range eventsChannel {
		if event.Kind != hook.KeyDown {
			continue
		}
		if settings.LogPressedChars {
			fmt.Printf("Pressed: char \"%q\", Keycode \"%v\"\n", event.Keychar, event.Rawcode)
		}

		sendPressActions(event.Keychar)
	}
}
