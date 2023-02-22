package acceptingServer

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"strings"
)

/*
SERVER PROTOCOL
COMMAND;PARAM_1:PARAM_2 etc.
*/

func parseRequest(receivedData *[]byte) (string, []string, error) {
	var converted string = string(bytes.Trim(*receivedData, "\x00"))

	var commandIndex int = strings.Index(converted, ";")
	if commandIndex == -1 {
		return "", make([]string, 0), errors.New("BAD DATA")
	}

	return converted[:commandIndex], strings.Split(converted[commandIndex+1:], ":"), nil
}

func Run(listener *net.Listener, buffer *Buffer) {
	for {
		conn, err := (*listener).Accept()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if _, err := conn.Read(buffer.arr); err != nil {
			panic(err)
		}

		command, params, err := parseRequest(&buffer.arr)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("COMMAND => \"%s\", PARAMS => %#v\n", command, params)
			defineAndProcessCommand(&command, &params)
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
