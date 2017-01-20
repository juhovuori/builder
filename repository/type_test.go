package repository

import "testing"

func TestValidate(t *testing.T) {
	cases := []struct {
		t   Type
		err error
	}{
		{"git", nil},
		{"nop", nil},
		{"unknown", ErrInvalidType},
	}
	for i, c := range cases {
		err := c.t.Validate()
		if err != c.err {
			t.Errorf("%d: Unexpected error %v\n", i, err)
		}
	}
}
