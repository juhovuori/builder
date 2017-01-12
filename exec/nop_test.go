package exec

import "testing"

func TestNopExecutor(t *testing.T) {
	e := nopExecutor{}
	ch, err := e.Run()
	if err != nil {
		t.Fatalf("Run returned error: %v\n", err)
	}
	status := <-ch
	if status != 0 {
		t.Errorf("Got wrong exit status %d, expected %d.\n", status, 0)
	}

	if err = e.Cleanup(); err != nil {
		t.Errorf("Cleanup returned error %v.\n", err)
	}
}
