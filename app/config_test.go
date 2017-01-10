package app

import (
	"fmt"
	"testing"
)

func TestNewConfig(t *testing.T) {
	cases := []struct {
		filename string
		err      string
	}{
		{"testdata/builder.hcl", "<nil>"},
		{"testdata/nothing-here.hcl", "open testdata/nothing-here.hcl: no such file or directory"},
	}
	for _, c := range cases {
		cfg, err := NewConfig(c.filename)
		if fmt.Sprintf("%v", err) != c.err {
			t.Errorf("Got %v, expected %v\n", err, c.err)
		}
		if c.err == "" && cfg == nil {
			t.Errorf("Expected non-nil-cfg\n")
			continue
		}
	}
}
