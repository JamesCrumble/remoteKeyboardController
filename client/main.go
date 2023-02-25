package main

import (
	server "client/acceptingServer"
	"client/helpers"
	s "client/settings"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"
)

var settings s.SettingsStruct = *s.Settings()

func defineLocalAddress() string {
	if settings.ClientLocalIpAddress != "" {
		return settings.ClientLocalIpAddress
	}
	if settings.InterfaceName == "" {
		panic(errors.New("CANNOT DEFINE LOCAL ADDRESS. clientLocalIpAddress and interfaceName are empty"))
	}

	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	fmt.Println("FOUNDED INTERFACES => ")
	for _, interface_ := range interfaces {
		interface_Addrs, err := interface_.Addrs()
		if err != nil {
			panic(err)
		}

		var interfaceData []string = helpers.ApplyAddrToString(interface_Addrs, func(addr net.Addr) string { return addr.String() })

		fmt.Printf(
			"INTERFACE \n\tNAME: \"%s\"\n\tADDRESS: \"%#v\"\n\n", interface_.Name, interfaceData,
		)
		if interface_.Name == settings.InterfaceName {

			var localIpAddress string
			var ipAddress string = interfaceData[1]
			var index = strings.Index(ipAddress, "/")

			if index == -1 {
				localIpAddress = ipAddress
			} else {
				localIpAddress = ipAddress[:index]
			}
			fmt.Printf("FOUND LOCAL IP ADDRESS TO CLIENT REGISTRATION => \"%s\"\n", localIpAddress)
			return localIpAddress
		}
	}
	panic(fmt.Errorf("CANNOT DEFINE LOCAL ADDRESS BY \"%s\" interfaceName", settings.InterfaceName))

}

func registerClient() {
	localIpAddress := defineLocalAddress()
	address := fmt.Sprintf("%s:%d", settings.RegistrationServer.Ip, settings.RegistrationServer.Port)
	writable := []byte(fmt.Sprintf("%s;%s:%d", "register", localIpAddress, settings.AcceptingServer.Port))

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
	registerClient()

	buffer := server.NewBuffer(32)
	serverListener := server.CreateListener(settings.AcceptingServer.Host, uint16(settings.AcceptingServer.Port))
	fmt.Printf("SOCKET LISTENING ON %s:%d\n", settings.AcceptingServer.Host, settings.AcceptingServer.Port)
	server.Run(&serverListener, &buffer)
}
