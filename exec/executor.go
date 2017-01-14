package exec

import (
	"fmt"

	"github.com/juhovuori/builder/build"
)

// Executor is a object that is capable of executing a build
type Executor interface {
	Run() (<-chan int, error)
	Cleanup() error
}

// New creates a new Executor
func New(b build.Build) (Executor, error) {
	switch b.ExecutorType() {
	case "fork":
		e := forkExecutor{
			Dir:       tmpFilename(),
			ScriptURL: b.Script(),
			Args:      []string{},
			Env:       []string{fmt.Sprintf("BUILD_ID=%s", b.ID())},
		}
		return &e, nil
	case "nop":
		e := nopExecutor{}
		return &e, nil
	default:
		return nil, ErrInvalidExecutor
	}
}
