package exec

// The fork executor is a simple executor that just runs the build script in a
// a temporary directory in a forked process.

type forkExecutor struct {
	dir string
}

func (f *forkExecutor) Run(URL string) (<-chan int, error) {
	var err error
	if err = f.createDir(); err != nil {
		return nil, err
	}
	defer f.cleanup()
	if err = f.copyScript(); err != nil {
		return nil, err
	}
	if err = f.run(); err != nil {
		return nil, err
	}
	return nil, nil
}

func (f *forkExecutor) createDir() error {
	return nil
}

func (f *forkExecutor) copyScript() error {
	return nil
}

func (f *forkExecutor) run() error {
	return nil
}

func (f *forkExecutor) cleanup() error {
	return nil
}
