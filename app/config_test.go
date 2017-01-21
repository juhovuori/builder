package app

import (
	"fmt"
	"testing"
)

func TestNewConfig(t *testing.T) {
	cases := []struct {
		filename string
		addr     string
		err      string
	}{
		{"testdata/builder.hcl", "0.0.0.0:8080", "<nil>"},
		{"testdata/invalid.hcl", "", "Failed to parse configuration: At 1:13: root.bind_addr: unknown type for string *ast.ListType"},
		{"testdata/nothing-here.hcl", "", "open testdata/nothing-here.hcl: no such file or directory"},
	}
	for i, c := range cases {
		cfg, err := NewConfig(c.filename)
		if fmt.Sprintf("%v", err) != c.err {
			t.Errorf("%d: Got %v, expected %v\n", i, err, c.err)
		}
		if c.err != "<nil>" {
			continue
		}
		if cfg.ServerAddress != c.addr {
			t.Errorf("%d: Got %s, expected%s\n", i, cfg.ServerAddress, c.addr)
		}
	}
}
