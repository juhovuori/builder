package app

import (
	"fmt"

	"github.com/hashicorp/hcl"
	"github.com/juhovuori/builder/fetcher"
)

// Config is the builder configuration
type Config struct {
	ServerAddress string          `hcl:"bind_addr"`
	URL           string          `hcl:"url"`
	Projects      []projectConfig `hcl:"projects"`
	Store         string          `hcl:"store"`
}

type projectConfig struct {
	Type       string `hcl:"type"`
	Repository string `hcl:"repository"`
	Config     string `hcl:"config"`
}

// FromString creates a Cfg from string.
func FromString(input string) (Config, error) {
	var cfg Config
	if err := hcl.Decode(&cfg, input); err != nil {
		return cfg, fmt.Errorf("Failed to parse configuration: %v", err)
	}
	return cfg, nil
}

// NewConfig creates a new configuration manager from given URL / filename
func NewConfig(filename string) (Config, error) {
	bytes, err := fetcher.Fetch(filename)
	if err != nil {
		return Config{}, err
	}
	return FromString(string(bytes))
}
