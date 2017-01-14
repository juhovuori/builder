package build

import "testing"

func TestNewContainer(t *testing.T) {
	cases := []struct {
		t   string
		err error
	}{
		{"memory", nil},
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
		b, err := New(&mock{"1", "2"})
		if err != nil {
			t.Fatal(err)
		}
		container.New(b)
		container.Builds()
		container.Build("")
		container.AddStage("", "")
	}
}
