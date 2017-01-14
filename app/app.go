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
	Projects() []string
	Project(project string) (project.Project, error)
	Builds() []string
	Build(build string) (build.Build, error)
	TriggerBuild(projectID string) (build.Build, error)
}

type defaultApp struct {
	projects project.Container
	builds   build.Container
	cfg      Config
}

func (a defaultApp) Config() Config {
	return a.cfg
}

func (a defaultApp) Projects() []string {
	return a.projects.Projects()
}

func (a defaultApp) Project(project string) (project.Project, error) {
	return a.projects.Project(project)
}

func (a defaultApp) Builds() []string {
	return a.builds.Builds()
}

func (a defaultApp) Build(build string) (build.Build, error) {
	return a.builds.Build(build)
}

func (a defaultApp) TriggerBuild(projectID string) (build.Build, error) {
	p, err := a.Project(projectID)
	if err != nil {
		return nil, err
	}
	b, err := a.builds.New(p)
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
	projects, err := project.NewContainer(cfg)
	if err != nil {
		return nil, err
	}
	builds, err := build.NewContainer("memory")
	if err != nil {
		return nil, err
	}
	newApp := defaultApp{
		projects,
		builds,
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
