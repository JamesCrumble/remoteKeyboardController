package settings

import (
	"client/helpers"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

const SETTINGS_FILENAME string = "client_settings"

var clientSettings SettingsStruct

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

	if err := yaml.Unmarshal(settingsContent, &clientSettings); err != nil {
		helpers.DelayedPanic(err)
	}

}

func Settings() *SettingsStruct {
	if (clientSettings == SettingsStruct{}) {
		initSettings()
		clientSettings.Validate()
	}
	return &clientSettings
}
