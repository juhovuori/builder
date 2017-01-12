package exec

// The nop executor does nothing.
type nopExecutor struct{}

func (f *nopExecutor) Run() (<-chan int, error) {
	c := make(chan int)
	go func() { c <- 0 }()
	return c, nil
}

func (f *nopExecutor) Cleanup() error {
	return nil
}
