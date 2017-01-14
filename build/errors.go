package build

import "errors"

var (
	// ErrNotFound is returned when queried build does not exist
	ErrNotFound = errors.New("Build not found")

	// ErrNilProject is returned when a build is created with no project
	ErrNilProject = errors.New("Project is nil")

	// ErrContainerType is returned when trying to create an unknwon container type
	ErrContainerType = errors.New("Invalid container type")

	// ErrStageType is returned when trying to create an unknwon stage type
	ErrStageType = errors.New("Invalid stage type")

	// ErrStageOrder is returned when stage is added in an invalid order.
	ErrStageOrder = errors.New("Invalid stage order")
)
