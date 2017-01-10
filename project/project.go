package project

import (
	"crypto/md5"
	"fmt"
)

// Project represents a single managed project
type Project interface {
	Manager
	Configuration
}

// project is the main implementation of Project
type project struct {
	url string
	md5 string
	err error
	cfg config
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

func newProject(URL string) Project {
	MD5 := md5.Sum([]byte(URL))
	config, err := fetchConfig(URL)
	p := project{
		url: URL,
		md5: fmt.Sprintf("%x", MD5),
		err: err,
		cfg: config,
	}
	return &p
}
