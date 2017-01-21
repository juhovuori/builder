package repository

import "testing"

func TestContainer(t *testing.T) {

	c := NewContainer()
	_, err := c.Ensure("nop", "http://example.com")
	if err != nil {
		t.Error(err)
	}
	r1, err := c.Ensure("nop", "http://example2.com")
	if err != nil {
		t.Error(err)
	}
	r2, err := c.Ensure("nop", "http://example2.com")
	if err != nil {
		t.Error(err)
	}
	if r1 != r2 {
		t.Error("Expected to get the same repository")
	}
	_, err = c.Ensure("invalid", "http://example2.com")
	if err != ErrInvalidType {
		t.Error(err)
	}
	r := c.Repositories()
	if len(r) != 2 {
		t.Errorf("Expected 2 builds, got %d\n", len(r))
	}
	_, err = c.Repository("nop", "http://example.com")
	if err != nil {
		t.Error(err)
	}
	_, err = c.Repository(git, "2")
	if err != ErrNotFound {
		t.Error(err)
	}
	err = c.Remove("nop", "http://example.com")
	if err != nil {
		t.Error(err)
	}
	err = c.Remove("nop", "http://example2.com")
	if err != nil {
		t.Error(err)
	}
	err = c.Remove("nop", "http://example2.com")
	if err != ErrNotFound {
		t.Error(err)
	}
	r = c.Repositories()
	if len(r) != 0 {
		t.Errorf("Expected 0 builds, got %d\n", len(r))
	}
}
