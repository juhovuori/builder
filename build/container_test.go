package build

import (
	"testing"
	"time"
)

func TestContainer(t *testing.T) {
	cases := []struct {
		t   string
		err error
	}{
		{"memory", nil},
		{"sqlite", nil},
		{"invalid", ErrContainerType},
	}
	for i, c := range cases {
		container, err := NewContainer(c.t)
		if err != c.err {
			t.Errorf("%d: Unexpected error %v, expected %v\n", i, err, c.err)
		}
		if c.err != nil {
			continue
		}

		if err = container.Purge(); err != nil {
			t.Error("Cannot purge container")
			continue
		}

		if err = container.Init(); err != nil {
			t.Error("Cannot initialize container")
			continue
		}
		if container == nil {
			t.Fatalf("%d: Nil container", i)
		}
		b, err := New(&mock{"1", "2"})
		if err != nil {
			t.Fatal(err)
		}

		bs := container.Builds(nil)
		if len(bs) != 0 {
			t.Errorf("%d: Expected empty container, got %d builds\n", i, len(bs))
		}
		before := time.Now().UnixNano()
		newB, err := container.New(b)
		if err != nil {
			t.Errorf("%d: Unexpected error %v\n", i, err)
		}
		after := time.Now().UnixNano()

		newB, err = container.Build(newB.ID())
		if err != nil {
			t.Errorf("%d: Unexpected error %v\n", i, err)
		}
		if len(newB.Stages()) != 0 {
			t.Errorf("%d: Expected %d stages, got %d\n", i, 0, len(newB.Stages()))
		}
		if newB.Created() < before || newB.Created() > after {
			t.Errorf("%d: before %d, created %d, after %d\n", i, before, newB.Created(), after)
		}

		container.Builds(nil)
		bs = container.Builds(nil)
		if len(bs) != 1 {
			t.Errorf("%d: Expected 1 build, got %d builds\n", i, len(bs))
		}
		pid := newB.ProjectID()
		bs = container.Builds(&pid)
		if len(bs) != 1 {
			t.Errorf("%d: Expected 1 build, got %d builds\n", i, len(bs))
		}
		pid = "invalid"
		bs = container.Builds(&pid)
		if len(bs) != 0 {
			t.Errorf("%d: Expected 0 builds, got %d builds\n", i, len(bs))
		}

		stages := []Stage{
			Stage{Type: STARTED},
			Stage{Type: PROGRESS},
			Stage{Type: PROGRESS},
			Stage{Type: SUCCESS},
		}
		for j, s := range stages {
			err = container.AddStage(newB.ID(), s)
			if err != nil {
				t.Errorf("%d, %d: Unexpected error %v\n", i, j, err)
				break
			}
			newB, err = container.Build(newB.ID())
			if err != nil {
				t.Errorf("%d, %d: Unexpected error %v\n", i, j, err)
			}
			if len(newB.Stages()) != j+1 {
				t.Errorf("%d, %d: Expected %d stages, got %d\n", i, j+1, j, len(newB.Stages()))
			}
		}
		err = container.Output(newB.ID(), []byte("foo"))
		if err != nil {
			t.Errorf("%d: Unexpected error %v\n", i, err)
		}
		err = container.Output(newB.ID(), []byte("bar"))
		if err != nil {
			t.Errorf("%d: Unexpected error %v\n", i, err)
		}
		newB, err = container.Build(newB.ID())
		if err != nil {
			t.Errorf("%d: Unexpected error %v\n", i, err)
		}
		if string(newB.Stdout()) != "foobar" {
			t.Errorf("%d: Invalid output %s, expected foobar\n", i, string(newB.Stdout()))
		}
		err = container.Close()
		if err != nil {
			t.Errorf("%d: Unexpected error %v\n", i, err)
		}
	}
}
