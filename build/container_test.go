package build

import "testing"

func TestNewContainer(t *testing.T) {
	cases := []struct {
		t   string
		err error
	}{
		{"memory", nil},
		{"invalid", ErrInvalidContainerType},
	}
	for i, c := range cases {
		container, err := NewContainer(c.t)
		if err != c.err {
			t.Errorf("%d: Unexpected error %v, expected %v\n", i, err, c.err)
		}
		if c.err != nil {
			continue
		}
		container.Builds()
		container.Build("")
		container.AddStage("", "")
		container.New(nil)
	}
}
