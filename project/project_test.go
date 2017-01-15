package project

import "testing"

func TestNewProject(t *testing.T) {
	cases := []struct {
		url     string
		success bool
		script  string
	}{
		{
			"testdata/project.hcl",
			true,
			"https://raw.githubusercontent.com/juhovuori/builder/master/scripts/build.sh",
		},
		{
			"testdata/garbage.hcl",
			false,
			"",
		},
		{
			"testdata/nothing-here.hcl",
			false,
			"",
		},
	}
	for i, c := range cases {
		p, err := NewProject(c.url)
		if c.success != (err == nil) {
			t.Errorf("%d: Got unexpected error %v\n", i, err)
		}
		if !c.success {
			if p.Err() == nil {
				t.Errorf("%d: Expected error, got nil\n", i)
			}
			continue
		}
		if p.Script() != c.script {
			t.Errorf("%d: Got script %s, expected %s\n", i, p.Script(), c.script)
		}
		if p.URL() != c.url {
			t.Errorf("%d: Got URL %s, expected %s\n", i, p.URL(), c.url)
		}
	}
}

func TestNewFromStringAndAccessors(t *testing.T) {
	cfg := `
	name = "name"
	description = "description"
	script = "script"
	`
	p, err := New(cfg)
	if err != nil {
		t.Errorf("Unexpected error %v\n", err)
	}
	if p.Script() != "script" {
		t.Errorf("Got script %s, expected %s\n", p.Script(), "script")
	}
	if p.URL() != "" {
		t.Errorf("Got URL %s, expected %s\n", p.URL(), "")
	}
	if p.ID() != "b18cfe21-db37-5154-951b-4359cbefd080" {
		t.Errorf("Got ID %s, expected %s\n", p.ID(), "b18cfe21-db37-5154-951b-4359cbefd080")
	}
	if p.Name() != "name" {
		t.Errorf("Got Name %s, expected %s\n", p.Name(), "")
	}
	if p.Description() != "description" {
		t.Errorf("Got Description %s, expected %s\n", p.Description(), "")
	}
}
