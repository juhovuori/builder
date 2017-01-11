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
	project.ProjectsConfig
	Store() storeCfg
}

type storeCfg struct {
	Type      string `hcl:"type"`
	Directory string `hcl:"directory"`
}

type builderCfg struct {
	ServerAddress string   `hcl:"bind_addr"`
	Projects      []string `hcl:"projects"`
	Store         storeCfg `hcl:"state_store"`
}

type cfgManager struct {
	cfg *builderCfg
}

func (cm cfgManager) ServerAddress() string {
	return cm.cfg.ServerAddress
}

func (cm cfgManager) Projects() []string {
	return cm.cfg.Projects
}

func (cm cfgManager) Store() storeCfg {
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
