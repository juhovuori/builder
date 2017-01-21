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
	Purl         string `json:"url"`
	Pid          string `json:"id"`
	Perr         error  `json:"error"`
	Pname        string `json:"name" hcl:"name"`
	Pdescription string `json:"description" hcl:"description"`
	Pscript      string `json:"script" hcl:"script"`
}

func (p *defaultProject) URL() string {
	return p.Purl
}

func (p *defaultProject) Config() string {
	return "project.hcl"
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
func New(repoID string, filename string, config []byte) (Project, error) {
	var p defaultProject
	err := hcl.Decode(&p, string(config))
	repoUUID, err2 := uuid.FromString(repoID)
	if err2 != nil {
		err = err2
	}
	p.Pid = uuid.NewV5(repoUUID, filename).String()
	p.Perr = err
	return &p, err
}
