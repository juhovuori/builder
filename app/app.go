package app

import (
	"fmt"
	"io"
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
	started := build.Stage{
		Type:      build.STARTED,
		Name:      "started",
		Timestamp: time.Now().UnixNano(),
	}
	err = a.builds.AddStage(b.ID(), started)
	if err != nil {
		return nil, err
	}
	ch, err := e.Run()
	if err != nil {
		return nil, err
	}
	go a.pipeOutput(b.ID(), e.Stdout())
	go a.monitorExit(b.ID(), ch)
	return b, nil
}

func (a defaultApp) pipeOutput(buildID string, stdout io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := stdout.Read(buf)
		if n != 0 {
			a.builds.Output(buildID, buf[:n])
		}
		if err == nil {
			continue
		}
		if err != io.EOF {
			log.Printf("Error reading stdout: %v\n", err)
		}
		break

	}
}

func (a defaultApp) monitorExit(buildID string, ch <-chan int) {
	exitStatus := <-ch
	log.Printf("Exit %d\n", exitStatus)
	if b, _ := a.builds.Build(buildID); !b.Completed() {
		t := build.SUCCESS
		if exitStatus != 0 {
			t = build.FAILURE
		}
		lastStage := build.Stage{
			Type:      t,
			Name:      "end-of-script",
			Timestamp: time.Now().UnixNano(),
		}
		err := a.builds.AddStage(buildID, lastStage)
		if err != nil {
			log.Printf("Could not add final stage.%v\n", err)
		}
	}
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

// New creates a new App from configuration
func New(cfg Config) (App, error) {
	builds, err := build.NewContainer(cfg.Store)
	if err != nil {
		return nil, err
	}

	repositories := repository.NewContainer()

	projects := project.NewContainer()
	for _, p := range cfg.Projects {
		project, err := project.NewProject(p.Repository)
		if err != nil {
			log.Printf("Could not create project %s: %v\n", p.Repository, err)
		}
		projects.Add(project)
	}

	newApp := defaultApp{
		projects,
		repositories,
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
