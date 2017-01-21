package project

import "errors"

var (
	// ErrNotFound is returned when queried project does not exist
	ErrNotFound = errors.New("Project not found")

	// ErrDuplicate is returned when adding duplicate project
	ErrDuplicate = errors.New("Duplicate project")
)
