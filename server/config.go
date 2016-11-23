package server

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/hcl"
)

// Config is the server configuration container
type Config *builderConfig

type stateStoreConfig struct {
	Type      string `hcl:"type"`
	Directory string `hcl:"directory"`
}

type builderConfig struct {
	ServerAddress string           `hcl:"bind_addr"`
	Projects      []string         `hcl:"projects"`
	StateStore    stateStoreConfig `hcl:"state_store"`
}

// ConfigFromString creates a config from string.
func ConfigFromString(input string) (Config, error) {
	var cfg builderConfig
	if err := hcl.Decode(&cfg, input); err != nil {
		return nil, fmt.Errorf("Failed to parse Configuration: %v", err)
	}
	fmt.Printf("%+v\n", cfg)
	return &cfg, nil
}

// DefaultConfig reads in a file and parses it to
// create a Config.
func DefaultConfig(filename string) (Config, error) {
	if filename == "" {
		filename = os.Getenv("BUILDER_CONFIG")
	}
	if filename == "" {
		wd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		filename = wd + "/builder.hcl"
	}
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return ConfigFromString(string(bytes))
}
