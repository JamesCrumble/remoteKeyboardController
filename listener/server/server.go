package server

import (
	"bytes"
	"errors"
	"fmt"
	"listener/commands"
	"net"
	"strings"
)

/*
SERVER PROTOCOL
COMMAND;PARAM_1:PARAM_2 etc.
*/

func parseRequest(requestBuffer *[]byte) (string, []string, error) {
	requestData := string(bytes.Trim(*requestBuffer, "\x00"))
	commandIndex := strings.Index(requestData, ";")
	if commandIndex == -1 {
		return "", make([]string, 0), errors.New("CANNOT PARSE REQUEST DATA")
	}

	return requestData[:commandIndex], strings.Split(requestData[commandIndex+1:], ":"), nil
}

func reactToRequest(requestBuffer *[]byte) error {
	commandName, params, err := parseRequest(requestBuffer)
	if err != nil {
		return err
	}
	fmt.Printf("get \"%s\" command name with %#v parameters\n", commandName, params)

	command, err := commands.GetCommand(commandName)
	if err != nil {
		return err
	}

	go command(params)
	return nil
}

func InfiniteListening(listener *net.Listener, buffer *Buffer) {
	for {
		conn, err := (*listener).Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		if _, err := conn.Read(buffer.arr); err != nil {
			panic(err)
		}
		if err := reactToRequest(&buffer.arr); err != nil {
			fmt.Println(err)
		}

		buffer.ResetBuffer()
		conn.Close()
	}
}

func CreateListener(host string, port uint16) net.Listener {
	serverListener, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		panic(err)
	}
	return serverListener
}
