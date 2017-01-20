package repository

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGit(t *testing.T) {

	repo, _ := New(git, "testdata/test.git")
	r := repo.(*gitRepository)
	// TODO: no dir
	if r.dir[:5] != "/tmp/" {
		t.Fatalf("Will only run tests in temp directory, got %s\n", r.dir)
	}
	err := r.Init()
	if err != nil {
		t.Error(err)
	}

	filename := r.dir + "/test"

	f, err := os.Open(filename)
	if err != nil {
		t.Error(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error(err)
	}
	f.Close()
	if string(data) != "ok\n" {
		t.Errorf("Read invalid data %s\n", string(data))
	}

	err = r.Update()
	if err != nil {
		t.Error(err)
	}
	// TODO: how to check

	err = r.Cleanup()
	if err != nil {
		t.Error(err)
	}

	f, err = os.Open(filename)
	if !os.IsNotExist(err) {
		t.Errorf("Expected not exist error, got %v\n", err)
	}
	f.Close()
}
