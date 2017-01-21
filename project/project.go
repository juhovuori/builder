package project

import (
	"github.com/hashicorp/hcl"
	"github.com/juhovuori/builder/fetcher"
	"github.com/juhovuori/builder/repository"
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
	Repository() repository.Repository
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

func (p *defaultProject) Repository() repository.Repository {
	return nil
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

func fetchConfig(filename string) (string, error) {
	bytes, err := fetcher.Fetch(filename)
	return string(bytes), err
}

func newProject(config string, URL string, err error) (*defaultProject, error) {
	var p defaultProject
	err2 := hcl.Decode(&p, config)
	if err == nil {
		err = err2
	}
	p.Purl = URL
	p.Pid = uuid.NewV5(namespace, URL).String()
	p.Perr = err
	return &p, err
}

// NewProject creates a new project
func NewProject(URL string) (Project, error) {
	config, err := fetchConfig(URL)
	return newProject(config, URL, err)
}

// New creates a new project from configuration string
func New(config string) (Project, error) {
	return newProject(config, "", nil)
}

// namespace is a global namespace for project ID generation.
var namespace, _ = uuid.FromString("a7cf1c8b-7b5e-4216-85d3-877e16845ebb")
