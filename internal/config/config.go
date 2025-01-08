package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

const configPath = "config.yml"

type AppConfig struct {
	DBUrl string `yaml:"db_url"`
}

var Env AppConfig

func ReadConfig() {
	f, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&Env)

	if err != nil {
		panic(err)
	}
}
