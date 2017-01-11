package exec

import (
	"reflect"
	"testing"

	"github.com/juhovuori/builder/build"
)

func TestNew(t *testing.T) {
	cases := []struct {
		b   build.Build
		t   string
		err error
	}{
		{mock{"fork"}, "*exec.forkExecutor", nil},
		{mock{"nop"}, "*exec.nopExecutor", nil},
		{mock{""}, "", ErrInvalidExecutor},
	}
	for _, c := range cases {
		e, err := New(c.b)
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

// Build describes a single build
type mock struct {
	t string
}

func (m mock) ID() string           { return "" }
func (m mock) ExecutorType() string { return m.t }
func (m mock) ProjectID() string    { return "" }
func (m mock) Completed() bool      { return false }
func (m mock) Error() error         { return nil }
func (m mock) Script() string       { return "" }