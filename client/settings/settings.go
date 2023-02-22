package settings

import (
	"fmt"
	yaml "gopkg.in/yaml.v3"
	"os"
	path "path/filepath"
	"time"
)

type AcceptingServer struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type RegistrationServer struct {
	Ip   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

type AcceptableButtons struct {
	Buttons []string `yaml:"buttons"`
}

func (struct_ *AcceptableButtons) Contains(char string) bool {
	for _, key := range struct_.Buttons {
		if key == char {
			return true
		}
	}
	return false
}

type SettingsStruct struct {
	AcceptingServer      *AcceptingServer    `yaml:"acceptingServer"`
	RegistrationServer   *RegistrationServer `yaml:"registrationServer"`
	AcceptableButtons    *AcceptableButtons  `yaml:"acceptableButtons"`
	ClientLocalIpAddress string              `yaml:"clientLocalIpAddress"`
}

const fileName string = "client_settings"

var clientSettings SettingsStruct

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

	if err := yaml.Unmarshal(settingsContent, &clientSettings); err != nil {
		delayPanic(err)
	}
}

func Settings() *SettingsStruct {
	if (clientSettings == SettingsStruct{}) {
		initSettings()
	}
	return &clientSettings
}
