package commands

import (
	"errors"
	"fmt"
	"listener/globals"
	"strconv"
)

const validationErrorMsg = "\"register\" command expected 2 arguments ip(string), port(int)"

func validateCommandParameters(params []string) (string, uint16, error) {
	if len(params) != 2 {
		return "", 0, errors.New(validationErrorMsg)
	}
	port, err := strconv.Atoi(params[1])
	if err != nil {
		return "", 0, errors.New(validationErrorMsg)
	}
	return params[0], uint16(port), nil
}

func registerCommand(params []string) {
	ip, port, err := validateCommandParameters(params)
	if err != nil {
		fmt.Println(err)
	}

	clients := *globals.Clients()

	for i := 0; i < len(clients); i++ {
		if (clients[i] != globals.Client{}) {
			continue
		}
		if clients[i].Ip == ip && clients[i].Port == port {
			fmt.Printf("Client with \"%s:%d\" address already exists\n", ip, port)
			return
		}

		clients[i] = globals.Client{Ip: ip, Port: port}
		fmt.Printf("New client with \"%s:%d\" address registered successfully\n", ip, port)
		return
	}
}
