package registerServer

import (
	s "controllerListener/settings"
	"fmt"
)

type Buffer struct {
	arr []byte
}

func (buffer *Buffer) ResetBuffer() {
	for i := 0; i < len(buffer.arr); i++ {
		buffer.arr[i] = 0
	}
}

func NewBuffer(buf_size int) Buffer {
	return Buffer{arr: make([]byte, buf_size)}
}

type Client struct {
	Ip   string
	Port uint16
}

// better to implement Clients struct with (new, append, contains, iter(channel which return non-empty Client structs) etc.) methods
// but doesn't matter actually for this type of program

var clientsArray []Client = make([]Client, s.Settings().ClientsPool)

func ClientsArray() *[]Client {
	return &clientsArray
}

func DeleteClient(client *Client) {
	for i := 0; i < len(clientsArray); i++ {
		if clientsArray[i].Ip == client.Ip && clientsArray[i].Port == client.Port {
			clientsArray[i] = Client{}
			return
		}
	}
	fmt.Printf("CANNOT DELETE CLIENT WITH \"%s\" IP AND \"%d\" PORT COUSE IT'S DOES NOT EXISTS\n", client.Ip, client.Port)
}
