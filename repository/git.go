package repository

import "errors"

type gitRepository struct {
	id  string
	url string
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
	return errors.New("not implemented")
}

func (r *gitRepository) Cleanup() error {
	return errors.New("not implemented")
}

func (r *gitRepository) Update() error {
	return errors.New("not implemented")
}
