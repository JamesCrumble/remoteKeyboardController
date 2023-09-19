package main

import (
	"client/helpers"
	"client/server"
	s "client/settings"
	"fmt"
	"net"
	"time"
)

const REGISTRATION_TIMEOUT = time.Second * 5
const REGISTRATION_RETRY_SLEEP = time.Second * 1

var settings s.SettingsStruct = *s.Settings()

func defineRegistrationIpAddress() string {
	if settings.ClientLocalIpAddress != "" {
		return settings.ClientLocalIpAddress
	}

	interfaces, err := net.Interfaces()
	if err != nil {
		helpers.DelayedPanic(err)
	}

	for _, netInterface := range interfaces {
		if netInterface.Name != settings.InterfaceName {
			continue
		}

		interfaceAddrs, err := netInterface.Addrs()
		if err != nil {
			helpers.DelayedPanic(err)
		}
		netInterfaceData := helpers.Apply(&interfaceAddrs, func(addr net.Addr) string { return addr.String() })
		ipAddress := helpers.IpAddressFromNetInterfaceData(&netInterfaceData)
		fmt.Printf("Found \"%s\" ip address for \"%s\" net interface name \n", ipAddress, settings.InterfaceName)
		return ipAddress
	}

	helpers.DelayedPanic(fmt.Sprintf("Cannot define address by \"%s\" net interface name\n", settings.InterfaceName))
	return ""
}

func registerClient() {
	registrationIpAddress := defineRegistrationIpAddress()
	registrationServerAddress := fmt.Sprintf("%s:%d", settings.RegistrationServer.Ip, settings.RegistrationServer.Port)
	registrationCommandMessage := []byte(fmt.Sprintf("%s;%s:%d", "register", registrationIpAddress, settings.Server.Port))

	fmt.Printf("Trying to inf register client on \"%s\" ip address\n", registrationServerAddress)

	dialer := net.Dialer{Timeout: REGISTRATION_TIMEOUT}
	for {
		conn, err := dialer.Dial("tcp4", registrationServerAddress)
		if err != nil {
			fmt.Printf("Cannot connect to registration server. New try after %d seconds\n", int(REGISTRATION_RETRY_SLEEP.Seconds()))
			time.Sleep(REGISTRATION_RETRY_SLEEP)
			continue
		}
		if _, err := conn.Write(registrationCommandMessage); err != nil {
			fmt.Printf("Cannot write to the socket => %v. New try after %d seconds\n", err, int(REGISTRATION_RETRY_SLEEP.Seconds()))
			time.Sleep(REGISTRATION_RETRY_SLEEP)
			continue
		}

		conn.Close()
		break
	}
}

func main() {
	registerClient()
	fmt.Printf("Listen socket on %s:%d\n", settings.Server.Host, settings.Server.Port)
	server.InfiniteListening()
}
