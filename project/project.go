package project

import (
	"crypto/md5"
	"fmt"

	"github.com/hashicorp/hcl"
	"github.com/juhovuori/builder/fetcher"
)

// Project represents a single managed project
type Project interface {
	Manager
	ProjectConfig
}

// Manager represents the manager of a single project
type Manager interface {
	URL() string
	ID() string
	Err() error
}

// ProjectConfig represents the configuration of a single project
type ProjectConfig interface {
	Name() string
	Description() string
	Script() string
}

// defaultProject is the main implementation of Project
type defaultProject struct {
	Purl         string `json:"url"`
	Pmd5         string `json:"id"`
	Perr         error  `json:"error"`
	Pname        string `hcl:"name"`
	Pdescription string `hcl:"description"`
	Pscript      string `hcl:"script"`
}

func (p *defaultProject) URL() string {
	return p.Purl
}

func (p *defaultProject) ID() string {
	return p.Pmd5
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
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

func newProject(config string, URL string, err error) (*defaultProject, error) {
	var p defaultProject
	if err == nil {
		if err = hcl.Decode(&p, config); err != nil {
			return &p, fmt.Errorf("Failed to parse project configuration: %v", err)
		}
	}
	MD5 := md5.Sum([]byte(URL))
	p.Purl = URL
	p.Pmd5 = fmt.Sprintf("%x", MD5)
	p.Perr = err
	return &p, err
}

// NewProject creates a new project
func NewProject(URL string) (Project, error) {
	config, err := fetchConfig(URL)
	return newProject(config, URL, err)
}

// NewFromString creates a new project from configuration string
func NewFromString(config string) (Project, error) {
	return newProject(config, "", nil)
}
