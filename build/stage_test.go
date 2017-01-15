package build

import "testing"

func TestStage(t *testing.T) {
	cases := []struct {
		stage       Stage
		predecessor *Stage
		err         error
	}{
		{Stage{Type: STARTED}, nil, nil},
		{Stage{Type: STARTED}, &Stage{Type: PROGRESS}, ErrStageOrder},
		{Stage{Type: PROGRESS}, nil, ErrStageOrder},
		{Stage{Type: PROGRESS}, &Stage{Type: PROGRESS}, nil},
		{Stage{Type: PROGRESS}, &Stage{Type: SUCCESS}, ErrStageOrder},
		{Stage{Type: ABORTED}, nil, ErrStageOrder},
		{Stage{Type: ABORTED}, &Stage{Type: PROGRESS}, nil},
		{Stage{Type: ABORTED}, &Stage{Type: SUCCESS}, ErrStageOrder},
		{Stage{Type: SUCCESS}, nil, ErrStageOrder},
		{Stage{Type: SUCCESS}, &Stage{Type: PROGRESS}, nil},
		{Stage{Type: SUCCESS}, &Stage{Type: SUCCESS}, ErrStageOrder},
		{Stage{Type: FAILURE}, nil, ErrStageOrder},
		{Stage{Type: FAILURE}, &Stage{Type: PROGRESS}, nil},
		{Stage{Type: FAILURE}, &Stage{Type: SUCCESS}, ErrStageOrder},

		{Stage{Type: "invalid"}, nil, ErrStageType},
		{Stage{Type: PROGRESS, Timestamp: 1}, &Stage{Type: PROGRESS, Timestamp: 2}, ErrStageOrder},
	}
	for i, c := range cases {
		err := c.stage.ValidateWithPredecessor(c.predecessor)
		if err != c.err {
			t.Errorf("%d: Got %v\nExpected %v for %v after %v\n", i, err, c.err, c.stage, c.predecessor)
		}
	}
}
