package main

import (
	"os"
	"os/exec"
	"testing"
)

// TestMainSuccess tests main function in a separate process using a trick from
// https://talks.golang.org/2014/testing.slide#1
func TestMainSuccess(t *testing.T) {
	if os.Getenv("TEST_MAIN_FUNC") == "1" {
		os.Args = []string{"", "nop"}
		main()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestMainSuccess")
	cmd.Env = append(os.Environ(), "TEST_MAIN_FUNC=1")
	err := cmd.Run()
	if err == nil {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 0", err)
}

// TestMainFailure tests main function in a separate process using a trick from
// https://talks.golang.org/2014/testing.slide#1
func TestMainFailure(t *testing.T) {
	if os.Getenv("TEST_MAIN_FUNC") == "1" {
		os.Args = []string{"", "fail"}
		main()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestMainFailure")
	cmd.Env = append(os.Environ(), "TEST_MAIN_FUNC=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}
