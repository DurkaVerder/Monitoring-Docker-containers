package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"Server"`
	Routes struct {
		AddPing    string `yaml:"addPing"`
		GetAllPing string `yaml:"getAllPing"`
	} `yaml:"Routes"`
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
