package main

import (
	server "client/acceptingServer"
	s "client/settings"
	"fmt"
	"net"
	"time"
)

func registerClient() {
	settings := s.Settings()
	address := fmt.Sprintf("%s:%d", settings.RegistrationServer.Ip, settings.RegistrationServer.Port)
	writable := []byte(fmt.Sprintf("%s;%s:%d", "register", settings.ClientLocalIpAddress, settings.AcceptingServer.Port))

	fmt.Printf("TRYING TO REGISTER CLIENT ON \"%s\"\n", address)

	for {
		conn, err := net.Dial("tcp4", address)
		if err != nil {
			fmt.Println(err.Error())
			time.Sleep(time.Second * 1)
			continue
		}
		if _, err := conn.Write(writable); err != nil {
			fmt.Printf("CANNOT SEND REGISTRATION COMMAND WITH ERROR => %s", err.Error())
		} else {
			conn.Close()
			break
		}
	}
}

func main() {
	settings := s.Settings()

	registerClient()

	buffer := server.NewBuffer(32)
	serverListener := server.CreateListener(settings.AcceptingServer.Host, uint16(settings.AcceptingServer.Port))
	fmt.Printf("SOCKET LISTENING ON %s:%d\n", settings.AcceptingServer.Host, settings.AcceptingServer.Port)
	server.Run(&serverListener, &buffer)
}
