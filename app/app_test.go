package app

import (
	"fmt"
	"testing"

	"github.com/juhovuori/builder/build"
	"github.com/juhovuori/builder/project"
	"github.com/juhovuori/builder/repository"
)

func TestNewFromFilename(t *testing.T) {
	cases := []struct {
		filename string
		err      string
	}{
		{"testdata/builder.hcl", "<nil>"},
		{"testdata/nothing-here.hcl", "open testdata/nothing-here.hcl: no such file or directory"},
	}
	for _, c := range cases {
		cfg, err := NewFromURL(c.filename)
		if fmt.Sprintf("%v", err) != c.err {
			t.Errorf("Got %v, expected %v\n", err, c.err)
		}
		if c.err == "" && cfg == nil {
			t.Errorf("Expected non-nil-cfg\n")
			continue
		}
	}
}

func TestTriggerBuild(t *testing.T) {
	projectID := "id"
	projectURL := "1"
	builds, err := build.NewContainer("memory")
	if err != nil {
		t.Fatal(err)
	}
	p := project.NewStaticProject(projectID)
	projects := project.NewStaticContainer(p)
	repositories := repository.NewContainer()
	cfg := builderCfg{Projects: []string{projectURL}}
	config := cfgManager{&cfg}
	app := defaultApp{
		projects,
		repositories,
		builds,
		config,
	}
	b, err := app.TriggerBuild(projectID)
	if err != nil {
		t.Fatalf("Unexpected error %v\n", err)
	}
	if b.ProjectID() != projectID {
		t.Errorf("Wrong buildID %v\n", b.ProjectID())
	}
}
