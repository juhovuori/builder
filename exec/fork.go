package exec

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"syscall"
)

// The fork executor is a simple executor that just runs the build script in a
// a temporary directory in a forked process.

type forkExecutor struct {
	Dir       string
	ScriptURL string
	Args      []string
	Env       []string
}

func (f *forkExecutor) SaveFile(relative string, data []byte) error {
	// TODO: ensure no leaks with ..
	filename := path.Join(f.Dir, relative)
	dir, _ := filepath.Split(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0755)
}

func (f *forkExecutor) Run(relative string, stdout chan<- []byte) error {
	var err error
	filename := path.Join(f.Dir, relative)
	cmd := exec.Command(filename, f.Args...)
	cmd.Dir = f.Dir
	cmd.Stdin = nil
	cmd.Env = append(os.Environ(), f.Env...)
	reader, err := cmd.StdoutPipe()
	if err != nil {
		close(stdout)
		return err
	}
	cmd.Stderr = cmd.Stdout
	if err := cmd.Start(); err != nil {
		close(stdout)
		return err
	}

	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if n != 0 {
			stdout <- buf[:n]
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

func (f *forkExecutor) Directory() string {
	return f.Dir
}

func (f *forkExecutor) Cleanup() error {
	return os.RemoveAll(f.Dir)
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
