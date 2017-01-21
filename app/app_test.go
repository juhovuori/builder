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
	repoID := "152d085f-de7d-4503-911a-cc10995b0551"
	projectURL := "1"
	builds, err := build.NewContainer("memory")
	if err != nil {
		t.Fatal(err)
	}
	p, err := project.New("nop", projectURL, repoID, "testconfig.hcl", []byte{})
	if err != nil {
		t.Fatal(err)
	}
	projects := project.NewContainer()
	if err = projects.Add(p); err != nil {
		t.Fatal(err)
	}
	repositories := repository.NewContainer()
	repositories.Ensure("nop", projectURL)
	config := Config{Projects: []projectConfig{{"nop", projectURL, "x"}}}
	app := defaultApp{
		projects,
		repositories,
		builds,
		config,
	}
	b, err := app.TriggerBuild(p.ID())
	if err != nil {
		t.Logf("%+v\n", repositories)
		t.Logf("%+v %+v %+v\n", p.ID(), p.VCS(), p.URL())
		t.Fatalf("Unexpected error %v\n", err)
	}
	if b.ProjectID() != p.ID() {
		t.Errorf("Wrong buildID %v\n", b.ProjectID())
	}
}
