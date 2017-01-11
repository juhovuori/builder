package exec

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/juhovuori/builder/project"
)

func TestForkExecutor(t *testing.T) {
	dir := tmpFilename()
	b := mockBuild{}
	e := forkExecutor{dir, &b}
	e.Run()

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
	dir := fmt.Sprintf("%s/builder-%019d", os.TempDir(), n)
	return dir
}

type mockBuild struct{}

func (b *mockBuild) ID() string               { return "" }
func (b *mockBuild) Project() project.Project { return nil }
func (b *mockBuild) Completed() bool          { return false }
func (b *mockBuild) Error() error             { return nil }
