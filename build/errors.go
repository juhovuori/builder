package build

import "errors"

var (
	// ErrNotFound is returned when queried build does not exist
	ErrNotFound = errors.New("Build not found")

	// ErrNilProject is returned when a build is created with no project
	ErrNilProject = errors.New("Project is nil")

	// ErrInvalidContainerType is returned when trying to create an unknwon container type
	ErrInvalidContainerType = errors.New("Invalid container type")

	errNotImplemented = errors.New("Not implemented")
)
