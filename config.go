package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Api_Key    string `yaml:"api_key"`
	Homeserver string `yaml:"homeserver"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Layout     string `yaml:"layout"`
}

func (c *Config) getConf() *Config {

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
