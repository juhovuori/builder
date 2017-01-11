package exec

import "errors"

var (
	// ErrInvalidExecutor is returned when executor of unknown type is attempted to create
	ErrInvalidExecutor = errors.New("Invalid executor type")
)
