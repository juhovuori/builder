package build

import "errors"

var (
	// ErrNilProject is returned when a build is created with no project
	ErrNilProject = errors.New("Project is nil")
)
