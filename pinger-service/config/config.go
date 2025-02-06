package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Response struct {
		Address       string `yaml:"address"`
		RetryCount    int    `yaml:"retryCount"`
		RetryInterval int    `yaml:"retryInterval"`
	} `yaml:"Response"`
	Worker struct {
		Count int `yaml:"count"`
	} `yaml:"Worker"`
	Channel struct {
		Size int `yaml:"size"`
	} `yaml:"Channel"`
}

func LoadConfig(path string) *Config {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		log.Fatalf("Error decoding YAML: %v", err)
	}

	return &cfg
}
