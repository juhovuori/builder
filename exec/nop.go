package exec

// The nop executor does nothing.
type nopExecutor struct{}

func (f *nopExecutor) Run() (<-chan int, error) {
	c := make(chan int)
	return c, nil
}

func (f *nopExecutor) Cleanup() error {
	return nil
}
