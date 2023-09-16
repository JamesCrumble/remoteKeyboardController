package commands

import "fmt"

var commands = map[string]func([]string){
	"press": pressCommand,
}

func GetCommand(commandName string) (func([]string), error) {
	command, ok := commands[commandName]
	if !ok {
		return nil, fmt.Errorf("\"%s\" command doesn't exists in mapping", commandName)
	}
	return command, nil
}
