package exec

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"syscall"

	"github.com/juhovuori/builder/fetcher"
)

// The fork executor is a simple executor that just runs the build script in a
// a temporary directory in a forked process.

type forkExecutor struct {
	Dir       string
	ScriptURL string
	Args      []string
	Env       []string
}

const scriptfilename = "script"
const stdoutfilename = "stdout"
const stderrfilename = "stderr"

func (f *forkExecutor) Run() (<-chan int, error) {
	var err error
	ch := make(chan int)
	if err = f.createDir(); err != nil {
		return nil, err
	}
	if err = f.copyScript(); err != nil {
		return nil, err
	}
	if err = f.run(ch); err != nil {
		return nil, err
	}
	return ch, nil
}

func (f *forkExecutor) Cleanup() error {
	return os.RemoveAll(f.Dir)
}

func (f *forkExecutor) createDir() error {
	return os.Mkdir(f.Dir, 0755)
}

func (f *forkExecutor) copyScript() error {
	data, err := fetcher.Fetch(f.ScriptURL)
	if err != nil {
		return err
	}
	filename := path.Join(f.Dir, scriptfilename)
	err = ioutil.WriteFile(filename, data, 0755)
	return err
}

func (f *forkExecutor) run(ch chan<- int) error {
	var err error
	filename := path.Join(f.Dir, scriptfilename)
	stdout := path.Join(f.Dir, stdoutfilename)
	stderr := path.Join(f.Dir, stderrfilename)
	cmd := exec.Command(filename, f.Args...)
	cmd.Dir = f.Dir
	cmd.Stdin = nil
	cmd.Env = append(os.Environ(), f.Env...)

	if cmd.Stdout, err = os.Create(stdout); err != nil {
		return err
	}
	if cmd.Stderr, err = os.Create(stderr); err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	go f.monitor(ch, cmd)
	return nil
}

func (f *forkExecutor) monitor(ch chan<- int, cmd *exec.Cmd) {
	if err := cmd.Wait(); err != nil {
		var exitErr *exec.ExitError
		var status syscall.WaitStatus
		var ok bool
		if exitErr, ok = err.(*exec.ExitError); !ok {
			log.Printf("Got unexpected error while waiting for child process %v\n", err)
			// TODO: communicate unexpected failure
			return
		}
		if status, ok = exitErr.Sys().(syscall.WaitStatus); !ok {
			log.Println("ExitError was not a WaitStatus. Is this Unix?")
			// TODO: communicate unexpected failure. Might happen on non-unix
			return
		}
		ch <- status.ExitStatus()
	}
	ch <- 0
}
