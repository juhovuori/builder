package build

import "github.com/juhovuori/builder/project"

// Build describes a single build
type Build interface {
	ID() string
	Project() project.Project
	Completed() bool
	Error() error
}

type defaultBuild struct {
	id        string
	project   project.Project
	completed bool
	err       error
}

func (b *defaultBuild) ID() string {
	return b.id
}

func (b *defaultBuild) Project() project.Project {
	return b.project
}

func (b *defaultBuild) Completed() bool {
	return b.completed
}

func (b *defaultBuild) Error() error {
	return b.err
}

// New creates a new build
func New(project project.Project) (Build, error) {
	return &defaultBuild{"", project, false, nil}, nil
}
