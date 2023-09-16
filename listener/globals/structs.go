package globals

import s "listener/settings"

var clients []Client = make([]Client, s.Settings().ClientsPool)

type Client struct {
	Ip   string
	Port uint16
}

func Clients() *[]Client {
	return &clients
}

func DeleteClient(client *Client) {
	for i := 0; i < len(clients); i++ {
		if clients[i].Ip != client.Ip || clients[i].Port != client.Port {
			continue
		}
		clients[i] = Client{}
		return
	}
}
