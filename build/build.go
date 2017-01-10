package build

import "github.com/juhovuori/builder/project"

// Build describes a single build
type Build interface {
	ID() string
	Project() project.Project
	Completed() bool
	Error() error
}

// New creates a new build
func New(project project.Project) (Build, error) {
	return nil, nil
}
