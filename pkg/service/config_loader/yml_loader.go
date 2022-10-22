package config_loader

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

const configPath = "configs/db_config.yml"

type dbYamlConfig struct {
	Port     int    `yaml:"port"`
	SSLMode  string `yaml:"sslmode"`
	TimeZone string `yaml:"TimeZone"`
}

func GetDbYamlConfig() dbYamlConfig {
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	loadedData := make(map[string]dbYamlConfig)
	err = yaml.Unmarshal(file, &loadedData)
	if err != nil {
		log.Fatal(err)
	}
	return loadedData["db"]
}
