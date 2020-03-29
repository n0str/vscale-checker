package cmd

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	ApiToken  string  `yaml:"apiToken"`
	Threshold float32 `yaml:"threshold"`
	Schedule  struct {
		Enabled bool   `yaml:"enabled"`
		Cron    string `yaml:"cron"`
	} `yaml:"schedule"`
	Notify struct {
		NotifyCodex `yaml:"codex"`
	} `yaml:"notify"`
}

type NotifyCodex struct {
	Enabled bool   `yaml:"enabled"`
	URL     string `yaml:"url"`
}

func LoadConfig(filename string) *Config {
	var config Config
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &config
}
