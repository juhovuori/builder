package exec

// Executor is a object that is capable of executing a build
type Executor interface {
	Run(URL string) (<-chan int, error)
}

// New creates a new Executor
func New() (Executor, error) {
	e := forkExecutor{}
	return &e, nil
}
