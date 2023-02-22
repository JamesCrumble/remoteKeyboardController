package settings

import (
	"fmt"
	yaml "gopkg.in/yaml.v3"
	"os"
	path "path/filepath"
	"time"
)

type RegisterServer struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type SettingsStruct struct {
	RegisterServer  RegisterServer `yaml:"registerServer"`
	ClientsPool     int            `yaml:"clientsPool"`
	LogPressedChars bool           `yaml:"logPressedChars"`
}

const fileName string = "server_settings"

var settingsStruct = SettingsStruct{}

func delayPanic(err error) {
	fmt.Println(err.Error())
	time.Sleep(time.Second * 5)
	panic(err)
}

func getPathToSettingsFile() string {
	workingDir, err := os.Getwd()
	if err != nil {
		delayPanic(err)
	}

	return path.Join(workingDir, fmt.Sprintf("%s.yaml", fileName))
}

func initSettings() {
	settingsContent, err := os.ReadFile(getPathToSettingsFile())
	if err != nil {
		delayPanic(err)
	}

	if err := yaml.Unmarshal(settingsContent, &settingsStruct); err != nil {
		delayPanic(err)
	}
}

func Settings() *SettingsStruct {
	if (settingsStruct == SettingsStruct{}) {
		initSettings()
	}
	return &settingsStruct
}
