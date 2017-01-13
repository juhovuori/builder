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
	for i, c := range cases {
		p, err := NewProject(c.filename)
		if c.success != (err == nil) {
			t.Errorf("%d: Got unexpected error %v\n", i, err)
		}
		if !c.success {
			continue
		}
		if p.Script() != c.script {
			t.Errorf("%d: Got script %s, expected %s\n", i, p.Script(), c.script)
		}
	}
}
