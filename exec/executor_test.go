package exec

import (
	"reflect"
	"testing"

	"github.com/juhovuori/builder/build"
)

func TestNew(t *testing.T) {
	cases := []struct {
		e   string
		t   string
		err error
	}{
		{"fork", "*exec.forkExecutor", nil},
		{"nop", "*exec.nopExecutor", nil},
		{"", "", ErrInvalidExecutor},
	}
	for _, c := range cases {
		b, err := build.NewWithExecutorType(mock{}, c.e)
		if err != nil {
			t.Errorf("Unexpected error %v\n", err)
		}
		e, err := New(b)
		if err != c.err {
			t.Errorf("Got %v, expect %+v\n", err, c.err)
		}
		if c.err != nil {
			continue
		}
		tname := reflect.TypeOf(e).String()
		if tname != c.t {
			t.Errorf("Invalid executor type %s, expected %s\n", tname, c.t)
		}

	}
}

type mock struct{}

func (m mock) Script() string { return "" }
func (m mock) ID() string     { return "" }
