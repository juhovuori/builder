package project

import (
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/hcl"
)

// Configuration represents the configuration of a single project
type Configuration interface {
	Name() string
	Description() string
}

// config represents the configurationof a single project
type config struct {
	Name        string `hcl:"name"`
	Description string `hcl:"description"`
}

func configFromString(input string) (config, error) {
	var cfg config
	if err := hcl.Decode(&cfg, input); err != nil {
		return config{}, fmt.Errorf("Failed to parse project configuration: %v", err)
	}
	return cfg, nil
}

func fetchConfig(filename string) (config, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return config{}, err
	}
	return configFromString(string(bytes))
}
