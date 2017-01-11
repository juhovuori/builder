package exec

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"sync"
	"testing"
	"time"
)

func TestForkExecutor(t *testing.T) {
	dir := tmpFilename()
	script := "testdata/success.sh"
	e := forkExecutor{dir, script}
	_, err := e.Run()
	if err != nil {
		t.Fatalf("Run returned error: %v\n", err)
	}
	info, err := os.Stat(dir)
	if err != nil {
		t.Fatalf("Error in stat %s: %v\n", dir, err)
	}
	if !info.IsDir() {
		t.Errorf("%s is not a directory.\n", dir)
	}

	err = e.Cleanup()

	if err != nil {
		t.Errorf("Cleanup returned error %v.\n", err)
	}

	info, err = os.Stat(dir)
	if err == nil {
		t.Errorf("Expected error %+v\n", info)
	}
}

var once sync.Once

func tmpFilename() string {
	once.Do(func() { rand.Seed(time.Now().UnixNano()) })
	n := rand.Int63()
	dir := path.Join(os.TempDir(), fmt.Sprintf("builder-%019d", n))
	return dir
}
