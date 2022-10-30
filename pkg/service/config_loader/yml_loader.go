package config_loader

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

const configPath = "configs/db_config.yml"
const jwtConfigPath = "configs/jwt_config.yml"

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

type jwtYamlConfig struct {
	TimeZone string `yaml:"TimeZone"`
}

func GetJwtYamlConfig() jwtYamlConfig {
	file, err := ioutil.ReadFile(jwtConfigPath)
	if err != nil {
		log.Fatal(err)
	}
	loadedData := make(map[string]jwtYamlConfig)
	err = yaml.Unmarshal(file, &loadedData)
	if err != nil {
		log.Fatal(err)
	}
	return loadedData["jwt"]
}
