package app

import "github.com/juhovuori/builder/project"

// App is the container for the whole builder application. This is used by
// frontends such as HTTP server or command line interface
type App interface {
	Config() Config
	Projects() []project.Project
}

type app struct {
	projects project.Projects
	cfg      Config
}

func (app app) Config() Config {
	return app.cfg
}

func (app app) Projects() []project.Project {
	return app.projects.Projects()
}

// New creates a new App from configuration
func New(cfg Config) (App, error) {
	projects, err := project.New(cfg)
	if err != nil {
		return nil, err
	}
	newApp := app{
		projects,
		cfg,
	}
	return newApp, nil
}

// NewFromFilename creates a new App from configuration filename
func NewFromFilename(filename string) (App, error) {
	cfg, err := NewConfig(filename)
	if err != nil {
		return nil, err
	}
	return New(cfg)
}
