package exec

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/juhovuori/builder/fetcher"
)

// The fork executor is a simple executor that just runs the build script in a
// a temporary directory in a forked process.

type forkExecutor struct {
	Dir       string
	ScriptURL string
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
	return os.Mkdir(f.Dir, 0755)
}

func (f *forkExecutor) copyScript() error {
	data, err := fetcher.Fetch(f.ScriptURL)
	if err != nil {
		return err
	}
	filename := path.Join(f.Dir, "script")
	err = ioutil.WriteFile(filename, data, 0755)
	return err
}

func (f *forkExecutor) run() error {
	return nil
}

func (f *forkExecutor) Cleanup() error {
	return os.RemoveAll(f.Dir)
}
