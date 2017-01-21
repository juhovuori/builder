package exec

import (
	"io/ioutil"
	"os"
	"strings"
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
		{"testdata/pwd.sh", 0, "<dir>\n"},
		{"testdata/fail.sh", 1, ""},
	}
	for i, c := range cases {
		dir := tmpFilename()
		c.output = strings.Replace(c.output, "<dir>", dir, 1)
		e := forkExecutor{dir, c.script, []string{}, []string{}}
		stdout := make(chan []byte)

		go func() {
			for data := range stdout {
				if string(data) != c.output {
					t.Errorf("%d: output %s, expected %s\n", i, string(data), c.output)
				}
			}
		}()

		data, err := ioutil.ReadFile(c.script)
		if err != nil {
			t.Error(i, err)
		}
		err = e.Run(data, stdout)
		status := AsUnixStatusCode(err)
		if status != c.status {
			t.Errorf("%d: Got status: %v, expected %v\n", i, status, c.status)
		}
		info, err := os.Stat(dir)
		if err != nil {
			t.Fatalf("%d: Error in stat %s: %v\n", i, dir, err)
		}
		if !info.IsDir() {
			t.Errorf("%d: %s is not a directory.\n", i, dir)
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
