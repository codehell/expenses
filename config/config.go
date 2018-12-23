package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	TokenKey string
}

func (c *Config) GetConfig() *Config {
	confFile := "./config.yaml"
	yamlFile, err := ioutil.ReadFile(confFile)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}
