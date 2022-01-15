package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type BlackWhiteList struct {
	Rooms  []string `yaml:"rooms",json:"rooms"`
	Active bool     `yaml:"active",json:"active"`
}

type Config struct {
	Api_Key    string         `yaml:"api_key",json:"api_key"`
	Homeserver string         `yaml:"homeserver",json:"homeserver"`
	Username   string         `yaml:"username",json:"username"`
	Password   string         `yaml:"password",json:"password"`
	Layout     string         `yaml:"layout",json:"layout"`
	Blacklist  BlackWhiteList `yaml:"blacklist",json:"blacklist"`
	Whitelist  BlackWhiteList `yaml:"whitelist",json:"whitelist"`
}

func (c *Config) getConf() *Config {

	jsonFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(jsonFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
