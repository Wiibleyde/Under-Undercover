package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Webserver struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"webserver"`
	Webhooks struct {
		LogsWebhook string `yaml:"logswebhook"`
	} `yaml:"webhooks"`
}

var config *Config

func InitConfig() bool {
	config = &Config{}
	configFile, err := os.Open("config.yaml")
	if err != nil {
		CreateFile()
		return true
	}
	defer configFile.Close()

	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(config)
	if err != nil {
		fmt.Println(err)
	}
	return false
}

func CreateFile() {
	configFile, err := os.Create("config.yaml")
	if err != nil {
		fmt.Println(err)
	}
	defer configFile.Close()

	encoder := yaml.NewEncoder(configFile)
	err = encoder.Encode(config)
	if err != nil {
		fmt.Println(err)
	}
}

func GetConfig() *Config {
	return config
}
