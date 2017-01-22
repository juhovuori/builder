package app

import (
	"log"
	"strconv"
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
	Builds(projectID *string) []string
	Build(build string) (build.Build, error)
	Stdout(build string) ([]byte, error)
	StageData(build string, stage string) ([]byte, error)
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

func (a defaultApp) Builds(projectID *string) []string {
	return a.builds.Builds(projectID)
}

func (a defaultApp) Build(build string) (build.Build, error) {
	return a.builds.Build(build)
}

func (a defaultApp) Stdout(buildID string) ([]byte, error) {
	b, err := a.Build(buildID)
	if err != nil {
		return nil, err
	}
	return b.Stdout(), err
}

func (a defaultApp) StageData(buildID string, stage string) ([]byte, error) {
	b, err := a.Build(buildID)
	if err != nil {
		return nil, err
	}
	i, err := strconv.Atoi(stage)
	if err != nil {
		return nil, err
	}
	return b.Stages()[i].Data, err
}

func (a defaultApp) TriggerBuild(projectID string) (build.Build, error) {
	p, err := a.Project(projectID)
	if err != nil {
		return nil, err
	}
	repository, err := a.repositories.Repository(repository.Type(p.VCS()), p.URL())
	if err != nil {
		return nil, err
	}
	script, err := repository.File(p.Script())
	if err != nil {
		return nil, err
	}
	b, err := a.builds.New(p)
	if err != nil {
		return nil, err
	}
	env := createEnv(a.cfg.URL, b.ID())
	e, err := exec.NewWithEnvironment(b, env)
	if err != nil {
		return nil, err
	}
	if err = e.SaveFile("script", script); err != nil {
		return nil, err
	}
	if err = a.builds.AddStage(b.ID(), build.StartStage()); err != nil {
		return nil, err
	}

	stdout := make(chan []byte)
	go func() {
		for data := range stdout {
			if a.Config().Verbose {
				log.Printf(">%s", string(data))
			}
			a.builds.Output(b.ID(), data)
		}
	}()
	go func() {

		err := e.Run("script", stdout)
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
	if err = repository.Init(); err != nil {
		log.Printf("Cannot initialize repository %v\n", err)
	}
	if err = repository.Update(); err != nil {
		log.Printf("Error updating repository %v\n", err)
	}
	config, err := repository.File(pc.Config)
	if err != nil {
		log.Printf("Cannot read configuration %v\n", err)
	}
	p, err := project.New(string(repository.Type()), repository.URL(), repository.ID(), pc.Config, config)
	if err != nil {
		log.Printf("Cannot create project: %v\n", err)
		return
	}
	if err = a.projects.Add(p); err != nil {
		log.Printf("Cannot add project to container %v\n", err)
	}
	log.Printf("Added project: %s - %s - %s\n", pc.Type, pc.Repository, pc.Config)
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
