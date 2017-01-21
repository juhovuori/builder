package exec

import "testing"

func TestNopExecutor(t *testing.T) {
	e := nopExecutor{}
	err := e.Run(nil, nil)
	if err != nil {
		t.Fatalf("Run returned error: %v\n", err)
	}
	if err = e.Cleanup(); err != nil {
		t.Errorf("Cleanup returned error %v.\n", err)
	}
}
