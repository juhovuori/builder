package exec

import "github.com/juhovuori/builder/build"

// Executor is a object that is capable of executing a build
type Executor interface {
	Run() (<-chan int, error)
	Cleanup() error
}

// New creates a new Executor
func New(b build.Build) (Executor, error) {
	e := forkExecutor{"/tmp/builder", b}
	return &e, nil
}
