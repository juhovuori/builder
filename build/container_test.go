package build

import "testing"

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
		err = container.Init(true)
		if err != nil {
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

		bs := container.Builds()
		if len(bs) != 0 {
			t.Errorf("%d: Expected empty container, got %d builds\n", i, len(bs))
		}
		newB, err := container.New(b)
		if err != nil {
			t.Errorf("%d: Unexpected error %v\n", i, err)
		}
		container.Builds()
		bs = container.Builds()
		if len(bs) != 1 {
			t.Errorf("%d: Expected 1 build, got %d builds\n", i, len(bs))
		}

		retrievedB, err := container.Build(newB.ID())
		if err != nil {
			t.Errorf("%d: Unexpected error %v\n", i, err)
		}

		if len(retrievedB.Stages()) != 0 {
			t.Errorf("%d: Expected 0 stages, got %d\n", i, len(retrievedB.Stages()))
		}
		err = container.AddStage(retrievedB.ID(), Stage{Type: CREATED})
		if err != nil {
			t.Errorf("%d: Unexpected error %v\n", i, err)
		}
		retrievedB, err = container.Build(newB.ID())
		if err != nil {
			t.Errorf("%d: Unexpected error %v\n", i, err)
		}

		if len(retrievedB.Stages()) != 1 {
			t.Errorf("%d: Expected 1 stages, got %d\n", i, len(retrievedB.Stages()))
		}
	}
}
