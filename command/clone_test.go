package command

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"sync"
	"testing"
	"time"
)

var once sync.Once

func TestClone(t *testing.T) {
	cmd, err := CloneFactory()
	if err != nil {
		t.Fatalf("Factory returned error %v\n", err)
	}
	s := cmd.Synopsis()
	if len(s) == 0 {
		t.Fatalf("Too brief synopsis %s\n", s)
	}
	h := cmd.Help()
	if len(s) == 0 {
		t.Fatalf("Too brief help %s\n", h)
	}
	os.Setenv("BUILDER_REPOSITORY", "any")
	os.Setenv("BUILDER_REPOSITORY_TYPE", "nop")
	status := cmd.Run([]string{"/tmp"})
	if status != 0 {
		t.Fatalf("Non-zero exit status %d\n", status)
	}
}

func tmpFilename() string {
	once.Do(func() { rand.Seed(time.Now().UnixNano()) })
	n := rand.Int63()
	dir := path.Join(os.TempDir(), fmt.Sprintf("builder-test-%019d", n))
	return dir
}
