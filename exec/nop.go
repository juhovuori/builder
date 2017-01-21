package exec

// The nop executor does nothing.
type nopExecutor struct{}

func (n *nopExecutor) SaveFile(relative string, data []byte) error {
	return nil
}

func (n *nopExecutor) Run(filename string, stdout chan<- []byte) error {
	if stdout != nil {
		close(stdout)
	}
	return nil
}

func (n *nopExecutor) Cleanup() error {
	return nil
}
