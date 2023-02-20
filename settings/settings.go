package settings

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed settings.yaml
var settingsFile []byte

type Api struct {
	Url   string `yaml:"url"`
	Token string `yaml:"token"`
}

type Settings struct {
	Port     string `yaml:"port"`
	External Api    `yaml:"api"`
}

func New() (*Settings, error) {
	var s Settings
	err := yaml.Unmarshal(settingsFile, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
