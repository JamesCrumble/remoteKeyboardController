package settings

import "client/helpers"

type Server struct {
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

type SettingsStruct struct {
	Server               *Server             `yaml:"server"`
	RegistrationServer   *RegistrationServer `yaml:"registrationServer"`
	AcceptableButtons    *AcceptableButtons  `yaml:"acceptableButtons"`
	ClientLocalIpAddress string              `yaml:"clientLocalIpAddress"`
	InterfaceName        string              `yaml:"interfaceName"`
}

func (settings *SettingsStruct) Validate() {
	if settings.InterfaceName == "" && settings.ClientLocalIpAddress == "" {
		helpers.DelayedPanic("One of this parameters should not be empty: clientLocalIpAddress, interfaceName")
	}
}
