package exec

// Executor is a object that is capable of executing a build
type Executor interface {
	Run() (<-chan int, error)
}

// New creates a new Executor
func New(URL string) (Executor, error) {
	e := forkExecutor{"/tmp/builder", URL}
	return &e, nil
}
