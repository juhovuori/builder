package exec

// The nop executor does nothing.
type nopExecutor struct{}

func (f *nopExecutor) Run(stdout chan<- []byte) error {
	if stdout != nil {
		close(stdout)
	}
	return nil
}

func (f *nopExecutor) Cleanup() error {
	return nil
}
