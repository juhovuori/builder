package app

import (
	"fmt"
	"testing"

	"github.com/juhovuori/builder/project"
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
		cfg, err := NewFromFilename(c.filename)
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
	projects := mockP{projectID}
	cfg := builderCfg{Projects: []string{projectURL}}
	config := cfgManager{&cfg}
	app := defaultApp{
		projects,
		config,
	}
	b, err := app.TriggerBuild(projectID)
	if err != nil {
		t.Fatalf("Unexpected error %v\n", err)
	}
	if b.Project().ID() != projectID {
		t.Errorf("Wrong buildID %v\n", b.Project().ID())
	}
}

type mockP struct {
	id string
}

func (p mockP) Name() string        { return "" }
func (p mockP) Description() string { return "" }
func (p mockP) Script() string      { return "" }
func (p mockP) URL() string         { return "" }
func (p mockP) ID() string          { return p.id }
func (p mockP) Err() error          { return nil }

func (p mockP) Configure([]string) {}
func (p mockP) Projects() []project.Project {
	return nil
}
func (p mockP) Project(id string) (project.Project, error) {
	if id == p.id {
		return p, nil
	}
	return nil, project.ErrNotFound
}
