package command

import "testing"

func TestShowConfig(t *testing.T) {
	cases := []struct {
		args   []string
		status int
	}{
		{[]string{}, 1},
		{[]string{"does-not-exist.hcl"}, 1},
		{[]string{"../app/testdata/builder.hcl"}, 0},
	}
	for _, c := range cases {
		cmd, err := ShowConfigFactory()
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
		status := cmd.Run(c.args)
		if status != c.status {
			t.Fatalf("Exit status %d, expected %d\n", status, c.status)
		}
	}
}
