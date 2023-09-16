package settings

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"listener/helpers"
	"os"
	"path/filepath"
)

const SETTINGS_FILENAME string = "server_settings"

var serverSettings SettingsStruct

func settingsFilePath() string {
	workingDir, err := os.Getwd()
	if err != nil {
		helpers.DelayedPanic(err)
	}

	return filepath.Join(workingDir, fmt.Sprintf("%s.yaml", SETTINGS_FILENAME))
}

func initSettings() {
	settingsContent, err := os.ReadFile(settingsFilePath())
	if err != nil {
		helpers.DelayedPanic(err)
	}

	if err := yaml.Unmarshal(settingsContent, &serverSettings); err != nil {
		helpers.DelayedPanic(err)
	}
}

func Settings() *SettingsStruct {
	if (serverSettings == SettingsStruct{}) {
		initSettings()
	}
	return &serverSettings
}
