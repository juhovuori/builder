package repository

import "testing"

func TestNew(t *testing.T) {
	cases := []struct {
		t   Type
		url string
		err error
	}{
		{"git", "htttp://x", nil},
		{"nop", "htttp://x", nil},
		{"unknown", "htttp://y", ErrInvalidType},
	}
	for i, c := range cases {
		r, err := New(c.t, c.url)
		if err != c.err {
			t.Errorf("%d: Unexpected error %v\n", i, err)
		}
		if c.err != nil {
			continue
		}
		if r.Type() != c.t {
			t.Errorf("%d: Got type %s, expected %s\n", i, r.Type(), c.t)
		}
		if r.URL() != c.url {
			t.Errorf("%d: Got URL %s, expected %s\n", i, r.URL(), c.url)
		}
	}

}
