package app

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/juhovuori/builder/build"
	"github.com/juhovuori/builder/exec"
	"github.com/juhovuori/builder/project"
	"github.com/juhovuori/builder/repository"
	"github.com/juhovuori/builder/version"
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
	AddStage(buildID string, stage build.Stage) error
	Version() version.Info
	Shutdown() (<-chan bool, error)
}

type defaultApp struct {
	projects     project.Container
	repositories repository.Container
	builds       build.Container
	cfg          Config
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
	env := []string{
		fmt.Sprintf("BUILD_ID=%s", b.ID()),
		fmt.Sprintf("URL=%s", a.cfg.URL),
	}
	e, err := exec.NewWithEnvironment(b, append(os.Environ(), env...))
	if err != nil {
		return nil, err
	}
	if err = a.builds.AddStage(b.ID(), build.StartStage()); err != nil {
		return nil, err
	}

	stdout := make(chan []byte)
	go func() {
		for data := range stdout {
			a.builds.Output(b.ID(), data)
		}
	}()

	go func() {
		err := e.Run(stdout)
		exitStatus := exec.AsUnixStatusCode(err)
		log.Printf("Exit %d\n", exitStatus)
		if !b.Completed() {
			stage := build.SuccessStage()
			if exitStatus != 0 {
				stage = build.FailureStage()
			}
			err := a.builds.AddStage(b.ID(), stage)
			if err != nil {
				log.Printf("Could not add final stage.%v\n", err)
			}
		}
	}()

	return b, nil
}

//AddStage adds a build stage
func (a defaultApp) AddStage(buildID string, stage build.Stage) error {
	stage.Timestamp = time.Now().UnixNano()
	return a.builds.AddStage(buildID, stage)
}

// Shutdown initiates a graceful shutdown
func (a defaultApp) Shutdown() (<-chan bool, error) {
	// TODO: stop creating builds
	// TODO: wait for builds to finnish instead of sleep
	log.Println("Initiating shutdown")
	ch := make(chan bool)
	go func() {
		<-time.After(time.Second * 2)
		ch <- true
	}()
	return ch, nil
}

// Version returns app version information
func (a defaultApp) Version() version.Info {
	return version.Version()
}

func (a defaultApp) addProject(pc projectConfig) {
	repository, err := a.repositories.Ensure(repository.Type(pc.Type), pc.Repository)
	if err != nil {
		log.Printf("Cannot add repository %v\n", err)
	}
	config, err := repository.File(pc.Config)
	if err != nil {
		log.Printf("Cannot read configuration %v\n", err)
	}
	p, err := project.New(repository.ID(), pc.Config, config)
	if err != nil {
		log.Printf("Cannot create project: %v\n", err)
		return
	}
	err = a.projects.Add(p)
	if err != nil {
		log.Printf("Cannot add project to container %v\n", err)
	}
	err = repository.Init()
	if err != nil {
		log.Printf("Cannot initialize repository %v\n", err)
	}
	err = repository.Update()
	if err != nil {
		log.Printf("Error updating repository %v\n", err)
	}
}

// New creates a new App from configuration
func New(cfg Config) (App, error) {
	builds, err := build.NewContainer(cfg.Store)
	if err != nil {
		return nil, err
	}

	repositories := repository.NewContainer()

	projects := project.NewContainer()

	newApp := defaultApp{
		projects,
		repositories,
		builds,
		cfg,
	}

	for _, p := range cfg.Projects {
		// TODO: add concurrently
		newApp.addProject(p)
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
