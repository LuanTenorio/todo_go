package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

const configPath = "config.yml"

type AppConfig struct {
	DBUrl string `yaml:"db_url"`
}

var Config AppConfig

func ReadConfig() {
	f, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&Config)

	if err != nil {
		panic(err)
	}
}
