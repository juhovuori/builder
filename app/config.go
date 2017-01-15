package app

import (
	"fmt"

	"github.com/hashicorp/hcl"
	"github.com/juhovuori/builder/fetcher"
	"github.com/juhovuori/builder/project"
)

// Config is the server configuration container
type Config interface {
	ServerAddress() string
	URL() string
	project.Config
	Store() string
}

type builderCfg struct {
	ServerAddress string   `hcl:"bind_addr"`
	URL           string   `hcl:"url"`
	Projects      []string `hcl:"projects"`
	Store         string   `hcl:"store"`
}

type cfgManager struct {
	cfg *builderCfg
}

func (cm cfgManager) ServerAddress() string {
	return cm.cfg.ServerAddress
}

func (cm cfgManager) URL() string {
	return cm.cfg.URL
}

func (cm cfgManager) Projects() []string {
	return cm.cfg.Projects
}

func (cm cfgManager) Store() string {
	return cm.cfg.Store
}

// FromString creates a Cfg from string.
func FromString(input string) (Config, error) {
	var cfg builderCfg
	if err := hcl.Decode(&cfg, input); err != nil {
		return nil, fmt.Errorf("Failed to parse configuration: %v", err)
	}
	return cfgManager{&cfg}, nil
}

// NewConfig creates a new configuration manager from given URL / filename
func NewConfig(filename string) (Config, error) {
	bytes, err := fetcher.Fetch(filename)
	if err != nil {
		return nil, err
	}
	return FromString(string(bytes))
}
