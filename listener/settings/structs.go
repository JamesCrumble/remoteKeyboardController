package settings

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type SettingsStruct struct {
	Server          Server `yaml:"server"`
	ClientsPool     int    `yaml:"clientsPool"`
	LogPressedChars bool   `yaml:"logPressedChars"`
}
