package repository

import (
	"os"
	"os/exec"
)

type gitRepository struct {
	id  string
	url string
	dir string
}

func (r *gitRepository) Type() Type {
	return git
}

func (r *gitRepository) URL() string {
	return r.url
}

func (r *gitRepository) ID() string {
	return r.id
}

func (r *gitRepository) Init() error {
	_, err := os.Stat(r.dir)
	if err == nil {
		return err
	}
	cmd := exec.Command("git", "clone", r.url, r.dir)
	return cmd.Run()
}

func (r *gitRepository) Cleanup() error {
	return os.RemoveAll(r.dir)
}

func (r *gitRepository) Update() error {
	cmd := exec.Command("git", "pull", "-f")
	cmd.Dir = r.dir
	return cmd.Run()
}
