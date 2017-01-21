package project

import (
	"io/ioutil"
	"testing"
)

func TestNewProject(t *testing.T) {
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
			"testdata/garbage.hcl",
			false,
			"",
		},
	}
	for i, c := range cases {
		data, err := ioutil.ReadFile(c.filename)
		if err != nil {
			t.Error(err)
		}
		p, err := New("", "", repoUUID, c.filename, data)
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
	}
}

func TestAccessors(t *testing.T) {
	cfg := `
	name = "name"
	description = "description"
	script = "script"
	`
	p, err := New("", "", repoUUID, "", []byte(cfg))
	if err != nil {
		t.Errorf("Unexpected error %v\n", err)
	}
	if p.Script() != "script" {
		t.Errorf("Got script %s, expected %s\n", p.Script(), "script")
	}
	if p.URL() != "" {
		t.Errorf("Got URL %s, expected %s\n", p.URL(), "")
	}
	if p.ID() != "ae715c65-c4ad-5e1c-b800-82e0300b740a" {
		t.Errorf("Got ID %s, expected %s\n", p.ID(), "ae715c65-c4ad-5e1c-b800-82e0300b740a")
	}
	if p.Name() != "name" {
		t.Errorf("Got Name %s, expected %s\n", p.Name(), "")
	}
	if p.Description() != "description" {
		t.Errorf("Got Description %s, expected %s\n", p.Description(), "")
	}
}

var repoUUID = "e787cc6c-4840-413e-a3ee-07b9b045e809"
