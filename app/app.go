package app

import (
	"github.com/juhovuori/builder/build"
	"github.com/juhovuori/builder/exec"
	"github.com/juhovuori/builder/project"
)

// App is the container for the whole builder application. This is used by
// frontends such as HTTP server or command line interface
type App interface {
	Config() Config
	Projects() []project.Project
	Project(project string) (project.Project, error)
	TriggerBuild(projectID string) (build.Build, error)
}

type defaultApp struct {
	projects project.Projects
	cfg      Config
}

func (a defaultApp) Config() Config {
	return a.cfg
}

func (a defaultApp) Projects() []project.Project {
	return a.projects.Projects()
}

func (a defaultApp) Project(project string) (project.Project, error) {
	return a.projects.Project(project)
}

func (a defaultApp) TriggerBuild(projectID string) (build.Build, error) {
	p, err := a.Project(projectID)
	if err != nil {
		return nil, err
	}
	b, err := build.New(p)
	if err != nil {
		return nil, err
	}
	e, err := exec.New(b)
	if err != nil {
		return nil, err
	}
	_, err = e.Run()
	if err != nil {
		return nil, err
	}
	return b, nil
}

// New creates a new App from configuration
func New(cfg Config) (App, error) {
	projects, err := project.New(cfg)
	if err != nil {
		return nil, err
	}
	newApp := defaultApp{
		projects,
		cfg,
	}
	return newApp, nil
}

// NewFromURL creates a new App from configuration filename
func NewFromURL(filename string) (App, error) {
	cfg, err := NewConfig(filename)
	if err != nil {
		return nil, err
	}
	return New(cfg)
}
