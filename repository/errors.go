package repository

import "errors"

var (
	// ErrInvalidType is returned when repository type is invalid
	ErrInvalidType = errors.New("Invalid repository type")

	// ErrDuplicate is returned when adding duplicate repository
	ErrDuplicate = errors.New("Duplicate repository")

	// ErrNotFound is returned when queried repository does not exist
	ErrNotFound = errors.New("Repository not found")
)
