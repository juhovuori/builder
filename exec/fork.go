package exec

import (
	"io"
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

func (f *forkExecutor) Run(stdout chan<- []byte) error {
	var err error
	if err = f.createDir(); err != nil {
		close(stdout)
		return err
	}
	if err = f.copyScript(); err != nil {
		close(stdout)
		return err
	}
	if err = f.run(stdout); err != nil {
		close(stdout)
		return err
	}
	return nil
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

func (f *forkExecutor) run(ch chan<- []byte) error {
	var err error
	filename := path.Join(f.Dir, scriptfilename)
	cmd := exec.Command(filename, f.Args...)
	cmd.Dir = f.Dir
	cmd.Stdin = nil
	cmd.Env = append(os.Environ(), f.Env...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	cmd.Stderr = cmd.Stdout
	if err := cmd.Start(); err != nil {
		return err
	}

	buf := make([]byte, 1024)
	for {
		n, err := stdout.Read(buf)
		if n != 0 {
			ch <- buf[:n]
		}
		if err != nil {
			if err != io.EOF {
				log.Printf("Error reading stdout: %v\n", err)
			}
			break
		}
	}

	return cmd.Wait()
}

// AsUnixStatusCode takes an error and resolves Unix status code from it
func AsUnixStatusCode(err error) int {
	var exitErr *exec.ExitError
	var status syscall.WaitStatus
	var ok bool
	if err == nil {
		return 0
	}
	if exitErr, ok = err.(*exec.ExitError); !ok {
		log.Printf("Got unexpected error %v\n", err)
		return -1
	}
	if status, ok = exitErr.Sys().(syscall.WaitStatus); !ok {
		log.Println("ExitError was not a WaitStatus. Is this Unix?")
		return -2
	}
	return status.ExitStatus()
}
