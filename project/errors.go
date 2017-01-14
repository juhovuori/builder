package project

import "errors"

var (
	// ErrNotFound is returned when queried project does not exist
	ErrNotFound = errors.New("Project not found")

	// ErrDuplicate is returned when adding duplicate project
	ErrDuplicate = errors.New("Duplicate project")

	// ErrFetchProject is returned when project was not fetched
	// correctly based on its URL
	ErrFetchProject = errors.New("Project failed to be fetched")
)
