package cfg

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/hcl"
)

// Cfg is the server configuration container
type Cfg interface {
	ServerAddress() string
	Projects() []string
	StateStore() stateStoreCfg
}

type stateStoreCfg struct {
	Type      string `hcl:"type"`
	Directory string `hcl:"directory"`
}

type builderCfg struct {
	ServerAddress string        `hcl:"bind_addr"`
	Projects      []string      `hcl:"projects"`
	StateStore    stateStoreCfg `hcl:"state_store"`
}

type cfgManager struct {
	cfg builderCfg
}

func (cm *cfgManager) ServerAddress() string {
	return cm.cfg.ServerAddress
}

func (cm *cfgManager) Projects() []string {
	return cm.cfg.Projects
}

func (cm *cfgManager) StateStore() stateStoreCfg {
	return cm.cfg.StateStore
}

// CfgFromString creates a Cfg from string.
func CfgFromString(input string) (Cfg, error) {
	var cfg builderCfg
	if err := hcl.Decode(&cfg, input); err != nil {
		return nil, fmt.Errorf("Failed to parse Configuration: %v", err)
	}
	fmt.Printf("%+v\n", cfg)
	return &cfgManager{cfg}, nil
}

// New creates a new configuration manager from given URL / filename
func New(filename string) (Cfg, error) {
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
	return CfgFromString(string(bytes))
}
