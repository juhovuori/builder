package project

import (
	"fmt"

	"github.com/hashicorp/hcl"
	"github.com/juhovuori/builder/fetcher"
)

// ProjectsConfig is the configurationfor projects object.
type ProjectsConfig interface {
	Projects() []string
}

// Configuration represents the configuration of a single project
type Configuration interface {
	Name() string
	Description() string
	Script() string
}

// config represents the configurationof a single project
type config struct {
	Name        string `hcl:"name"`
	Description string `hcl:"description"`
	Script      string `hcl:"script"`
}

func configFromString(input string) (config, error) {
	var cfg config
	if err := hcl.Decode(&cfg, input); err != nil {
		return config{}, fmt.Errorf("Failed to parse project configuration: %v", err)
	}
	return cfg, nil
}

func fetchConfig(filename string) (config, error) {
	bytes, err := fetcher.Fetch(filename)
	if err != nil {
		return config{}, err
	}
	return configFromString(string(bytes))
}
