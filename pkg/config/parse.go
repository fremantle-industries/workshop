package config

import (
	"log"

	"github.com/fremantle-industries/workshop/pkg/inbound_connectors"
	"github.com/fremantle-industries/workshop/pkg/outbound_connectors"
	"gopkg.in/yaml.v3"
	// "github.com/caarlos0/env/v6"
)

var defaultUiPort = "8080"
var defaultUiEnabled = false
var defaultAPIPort = "8081"
var defaultAPIEnabled = false

func New(data []byte) *Config {
	var c Config

	err := yaml.Unmarshal([]byte(data), &c)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	return &c
}

func NewFromYAMLFile(path string) *Config {
	var data []byte
	return New(data)
}

type Config struct {
	UI                 UIConfig                   `yaml:"ui"`
	API                APIConfig                  `yaml:"api"`
	InboundConnectors  inbound_connectors.Config  `yaml:"outbound_connectors"`
	OutboundConnectors outbound_connectors.Config `yaml:"inbound_connectors"`
}

type UIConfig struct {
	Enabled bool   `yaml:"enabled"`
	Port    string `yaml:"port"`
}

type APIConfig struct {
	Enabled bool   `yaml:"enabled"`
	Port    string `yaml:"port"`
}
