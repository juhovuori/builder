package command

import "testing"

// TestServer tests bits of Server command. In go <1.8 ListenAndServe
// is not stoppable so we don't test it properly. Testing in a subprocess
// is more pain than gain.
func TestServer(t *testing.T) {
	cmd, err := ServerFactory()
	if err != nil {
		t.Fatalf("Factory returned error %v\n", err)
	}
	s := cmd.Synopsis()
	if len(s) == 0 {
		t.Fatalf("Too brief synopsis %s\n", s)
	}
	h := cmd.Help()
	if len(s) == 0 {
		t.Fatalf("Too brief help %s\n", h)
	}
	status := cmd.Run([]string{"-invalid"})
	if status == 0 {
		t.Fatalf("Expected non-zero exit status, got %d\n", status)
	}
}
