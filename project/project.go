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
	Purl string `json:"url"`
	Pmd5 string `json:"id"`
	Perr error  `json:"error"`
	cfg  config
}

func (p *project) URL() string {
	return p.Purl
}

func (p *project) ID() string {
	return p.Pmd5
}

func (p *project) Err() error {
	return p.Perr
}

func (p *project) Name() string {
	return p.cfg.Name
}

func (p *project) Description() string {
	return p.cfg.Description
}

func (p *project) Script() string {
	return p.cfg.Script
}

func newProject(URL string) Project {
	MD5 := md5.Sum([]byte(URL))
	config, err := fetchConfig(URL)
	p := project{
		Purl: URL,
		Pmd5: fmt.Sprintf("%x", MD5),
		Perr: err,
		cfg:  config,
	}
	return &p
}
