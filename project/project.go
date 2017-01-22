package project

import (
	"github.com/hashicorp/hcl"
	uuid "github.com/satori/go.uuid"
)

// Project represents a single managed project
type Project interface {
	Manager
	Attributes
}

// Manager represents the manager of a single project
type Manager interface {
	VCS() string
	URL() string
	Config() string
	ID() string
	Err() error
}

// Attributes represents the configuration of a single project
type Attributes interface {
	Name() string
	Description() string
	Script() string
}

// defaultProject is the main implementation of Project
type defaultProject struct {
	Pid          string
	Perr         error
	Pname        string `hcl:"name"`
	Pdescription string `hcl:"description"`
	Pscript      string `hcl:"script"`
	configfile   string
	url          string
	vcs          string
}

func (p *defaultProject) URL() string {
	return p.url
}

func (p *defaultProject) VCS() string {
	return p.vcs
}

func (p *defaultProject) Config() string {
	return p.configfile
}

func (p *defaultProject) ID() string {
	return p.Pid
}

func (p *defaultProject) Err() error {
	return p.Perr
}

func (p *defaultProject) Name() string {
	return p.Pname
}

func (p *defaultProject) Description() string {
	return p.Pdescription
}

func (p *defaultProject) Script() string {
	return p.Pscript
}

// New creates a new project from configuration string
func New(VCS string, URL string, repoID string, filename string, config []byte) (Project, error) {
	var p defaultProject
	err := hcl.Decode(&p, string(config))
	repoUUID, err2 := uuid.FromString(repoID)
	if err2 != nil {
		err = err2
	}
	p.Pid = uuid.NewV5(repoUUID, filename).String()
	p.configfile = filename
	p.Perr = err
	p.vcs = VCS
	p.url = URL
	return &p, err
}
