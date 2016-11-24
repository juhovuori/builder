package projects

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/hcl"
)

// ProjectManager represents the manager of a single project
type ProjectManager interface {
	URL() string
	MD5() string
	Err() error
	Watched() bool
}

// ProjectConfiguration represents the configuration of a single project
type ProjectConfiguration interface {
	Name() string
	Description() string
}

// Project represents a single managed project
type Project interface {
	ProjectManager
	ProjectConfiguration
}

// projectCfg represents the configurationof a single project
type projectCfg struct {
	Name        string `hcl:"name"`
	Description string `hcl:"description"`
}

// project is the main implementation of Project
type project struct {
	url string
	md5 string
	err error
	cfg projectCfg
}

func (p *project) URL() string {
	return p.url
}

func (p *project) MD5() string {
	return p.md5
}

func (p *project) Err() error {
	return p.err
}

func (p *project) Name() string {
	return p.cfg.Name
}

func (p *project) Description() string {
	return p.cfg.Description
}

func (p *project) Watched() bool {
	return false
}

func projectCfgFromString(input string) (projectCfg, error) {
	var cfg projectCfg
	if err := hcl.Decode(&cfg, input); err != nil {
		return projectCfg{}, fmt.Errorf("Failed to parse project configuration: %v", err)
	}
	return cfg, nil
}

func fetchProjectCfg(filename string) (projectCfg, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return projectCfg{}, err
	}
	return projectCfgFromString(string(bytes))
}

func newProject(URL string) Project {
	MD5 := md5.Sum([]byte(URL))
	projectCfg, err := fetchProjectCfg(URL)
	p := project{
		url: URL,
		md5: fmt.Sprintf("%x", MD5),
		err: err,
		cfg: projectCfg,
	}
	return &p
}
