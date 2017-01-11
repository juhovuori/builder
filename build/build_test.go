package build

import "testing"

func TestBuild(t *testing.T) {
	cases := []struct {
		project Buildable
		exec    string
		err     error
	}{
		{&mock{"1", "sh"}, "fork", nil},
		{&mock{"2", ""}, "nop", nil},
		{nil, "fork", ErrNilProject},
	}
	for _, c := range cases {
		b, err := New(c.project)
		if err != c.err {
			t.Fatalf("Got error %v\n", err)
		}
		if c.err != nil {
			continue
		}
		if b.ID() != "" {
			t.Errorf("Wrong ID %v\n", b.ID())
		}
		if b.ExecutorType() != c.exec {
			t.Errorf("Wrong ExecutorType %v\n", b.ExecutorType())
		}
		if b.ProjectID() != c.project.ID() {
			t.Errorf("Wrong ProjectID %v\n", b.ProjectID())
		}
		if b.Script() != c.project.Script() {
			t.Errorf("Wrong Script %v\n", b.Script())
		}
		if b.Completed() != false {
			t.Errorf("Wrong Completed %v\n", b.Completed())
		}
		if b.Error() != nil {
			t.Errorf("Wrong Error %v\n", b.Error())
		}
	}
}

type mock struct {
	id     string
	script string
}

func (m *mock) Script() string {
	return m.script
}

func (m *mock) ID() string {
	return m.id
}
