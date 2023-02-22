package registerServer

import (
	"errors"
	"fmt"
	"strconv"
)

/*

COMMANDS:
  - register; ip : port
  - delete; ip : port

*/

func registerCommand(ip string, port uint16) {
	for i := 0; i < len(clientsArray); i++ {
		if (clientsArray[i] != Client{}) {
			continue
		} else if clientsArray[i].Ip == ip {
			fmt.Printf("CLIENT WITH \"%s\" IP AND \"%d\" PORT ALREADY EXISTS\n", ip, port)
			return
		}

		clientsArray[i] = Client{ip, port}
		fmt.Printf("REGISTER NEW CLIENT WITH \"%s\" IP AND \"%d\" PORT\n", ip, port)
		return
	}
	fmt.Printf("CANNOT REGISTER NEW CLIENT WITH \"%s\" IP AND \"%d\" PORT\n", ip, port)
}

func defineAndProcessCommand(command *string, params *[]string) error {
	if *command == "register" {
		if len(*params) != 2 {
			return errors.New("\"register\" COMMAND EXPECTED 2 ARGUMENTS ((string)ip, (uint16)port)")
		}
		port, err := strconv.Atoi((*params)[1])
		if err != nil {
			return errors.New("\"register\" COMMAND EXPECTED 2 ARGUMENTS ((string)ip, (uint16)port)")
		}
		registerCommand((*params)[0], uint16(port))
	}
	return nil
}
