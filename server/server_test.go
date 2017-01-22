package server

import (
	"testing"

	"github.com/juhovuori/builder/app"
)

func TestServer(t *testing.T) {
	cfg := app.Config{Store: "memory"}
	a, err := app.New(cfg)
	if err != nil {
		t.Error(err)
	}
	server, err := New(a)
	if err != nil {
		t.Error(err)
	}
	_ = server
	// TODO: start, 200, 404 and echo.Shutdown()
}
