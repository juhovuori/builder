package exec

import (
	"io"

	"github.com/juhovuori/builder/build"
)

// Executor is a object that is capable of executing a build
type Executor interface {
	Run() (<-chan int, error)
	Stdout() io.Reader
	Cleanup() error
}

// New creates a new Executor
func New(b build.Build) (Executor, error) {
	return NewWithEnvironment(b, []string{})
}

// NewWithEnvironment creates a new Executor with environment
func NewWithEnvironment(b build.Build, env []string) (Executor, error) {
	switch b.ExecutorType() {
	case "fork":
		e := forkExecutor{
			Dir:       tmpFilenameByID(b.ID()),
			ScriptURL: b.Script(),
			Args:      []string{},
			Env:       env,
		}
		return &e, nil
	case "nop":
		e := nopExecutor{}
		return &e, nil
	default:
		return nil, ErrInvalidExecutor
	}
}
