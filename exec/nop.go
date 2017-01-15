package exec

import "io"

// The nop executor does nothing.
type nopExecutor struct{}

func (f *nopExecutor) Run() (<-chan int, error) {
	c := make(chan int)
	go func() { c <- 0 }()
	return c, nil
}

func (f *nopExecutor) Stdout() io.Reader {
	return nil
}

func (f *nopExecutor) Cleanup() error {
	return nil
}
