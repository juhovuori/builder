package exec

import "github.com/juhovuori/builder/build"

// The fork executor is a simple executor that just runs the build script in a
// a temporary directory in a forked process.

type forkExecutor struct {
	dir string
	b   build.Build
}

func (f *forkExecutor) Run() (<-chan int, error) {
	var err error
	if err = f.createDir(); err != nil {
		return nil, err
	}
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

func (f *forkExecutor) Cleanup() error {
	return nil
}
