package project

import "testing"

func TestNewConfig(t *testing.T) {
	cases := []struct {
		filename string
		success  bool
		script   string
	}{
		{
			"testdata/project.hcl",
			true,
			"https://raw.githubusercontent.com/juhovuori/builder/master/scripts/build.sh",
		},
		{
			"testdata/nothing-here.hcl",
			false,
			"",
		},
	}
	for _, c := range cases {
		cfg, err := fetchConfig(c.filename)
		if c.success != (err == nil) {
			t.Errorf("Got unexpected error %v\n", err)
		}
		if !c.success {
			continue
		}
		if cfg.Script != c.script {
			t.Errorf("Got script %s, expected %s\n", cfg.Script, c.script)
		}
	}
}
