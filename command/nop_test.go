package command

import "testing"

func TestNop(t *testing.T) {
	cmd, err := NopFactory()
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
	status := cmd.Run([]string{})
	if status != 0 {
		t.Fatalf("Non-zero exit status %d\n", status)
	}
}
