package exec

import (
	"os"
	"testing"
)

func TestForkExecutor(t *testing.T) {
	cases := []struct {
		script string
		status int
		output string
	}{
		{"testdata/success.sh", 0, ""},
		{"testdata/output.sh", 0, "something\n"},
		{"testdata/stderr.sh", 0, "some error\n"},
		{"testdata/fail.sh", 1, ""},
		//TODO: pwd test
	}
	for i, c := range cases {
		dir := tmpFilename()
		e := forkExecutor{dir, c.script, []string{}, []string{}, nil}
		ch, err := e.Run()
		if err != nil {
			t.Fatalf("%d: Run returned error: %v\n", i, err)
		}
		info, err := os.Stat(dir)
		if err != nil {
			t.Fatalf("%d: Error in stat %s: %v\n", i, dir, err)
		}
		if !info.IsDir() {
			t.Errorf("%d: %s is not a directory.\n", i, dir)
		}
		buf := make([]byte, 20)
		n, err := e.stdout.Read(buf)
		if err != nil {
			t.Logf("%d: Error reading output %v\n", i, err)
		}
		buf = buf[:n]
		if string(buf) != c.output {
			t.Errorf("%d: output %s, expected %s\n", i, string(buf), c.output)
		}

		status := <-ch
		if status != c.status {
			t.Errorf("%d: Got wrong exit status %d, expected %d.\n", i, status, c.status)
		}

		if err = e.Cleanup(); err != nil {
			t.Errorf("%d: Cleanup returned error %v.\n", i, err)
		}

		info, err = os.Stat(dir)
		if err == nil {
			t.Errorf("%d: Expected error %+v\n", i, info)
		}
	}
}
