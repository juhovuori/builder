package exec

import (
	"os"
	"testing"
)

func TestForkExecutor(t *testing.T) {
	cases := []struct {
		script string
		status int
	}{
		{"testdata/success.sh", 0},
		{"testdata/fail.sh", 1},
		//TODO: pwd test
		//TODO test stdout/stderr redirection
	}
	for _, c := range cases {
		dir := tmpFilename()
		e := forkExecutor{dir, c.script}
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
		if status != c.status {
			t.Errorf("Got wrong exit status %d, expected %d.\n", status, c.status)
		}

		if err = e.Cleanup(); err != nil {
			t.Errorf("Cleanup returned error %v.\n", err)
		}

		info, err = os.Stat(dir)
		if err == nil {
			t.Errorf("Expected error %+v\n", info)
		}
	}
}
