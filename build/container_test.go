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

		stages := []Stage{
			Stage{},
			Stage{Type: CREATED},
			Stage{Type: STARTED},
			Stage{Type: PROGRESS},
			Stage{Type: PROGRESS},
			Stage{Type: SUCCESS},
		}
		for j, s := range stages {
			if j != 0 {
				err = container.AddStage(newB.ID(), s)
				if err != nil {
					t.Errorf("%d, %d: Unexpected error %v\n", i, j, err)
				}
			}
			newB, err = container.Build(newB.ID())
			if err != nil {
				t.Errorf("%d, %d: Unexpected error %v\n", i, j, err)
			}
			if len(newB.Stages()) != j {
				t.Errorf("%d, %d: Expected %d stages, got %d\n", i, j, j, len(newB.Stages()))
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
	}
}
