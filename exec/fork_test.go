package exec

import (
	"os"
	"testing"
)

func TestForkExecutor(t *testing.T) {
	dir := tmpFilename()
	script := "testdata/success.sh"
	e := forkExecutor{dir, script}
	ch, err := e.Run()
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

	status := <-ch
	if status != 0 {
		t.Errorf("Got wrong exit status %d.\n", status)
	}

	if err = e.Cleanup(); err != nil {
		t.Errorf("Cleanup returned error %v.\n", err)
	}

	info, err = os.Stat(dir)
	if err == nil {
		t.Errorf("Expected error %+v\n", info)
	}
}
