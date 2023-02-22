package main

import (
	server "controllerListener/registerServer"
	"fmt"
	"net"
)

/*
SEND ACTION PROTOCOL

COMMAND;PARAM_1:PARAM_2 etc.

COMMANDS:
	- press;char
*/

func sendPressAction(client *server.Client, char rune) {
	conn, err := net.Dial("tcp4", fmt.Sprintf("%s:%d", (*client).Ip, (*client).Port))
	if err != nil {
		fmt.Println(err.Error())
		server.DeleteClient(client)
		return
	}
	if _, err := conn.Write([]byte(fmt.Sprintf("%s;%s", "press", string(char)))); err != nil {
		fmt.Println(err.Error())
	} else {
		conn.Close()
	}
}

func sendPressActions(char rune) {
	for i := 0; i < len(*server.ClientsArray()); i++ {
		if ((*server.ClientsArray())[i] == server.Client{}) {
			continue
		}
		sendPressAction(&(*server.ClientsArray())[i], char)
	}
}
