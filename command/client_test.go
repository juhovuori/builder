package command

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handlerFactory(200, "OK")))
	cmd, err := ClientFactory()
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
	os.Setenv("BUILDER_URL", server.URL)
	os.Setenv("BUILDER_BUILDID", "2")
	os.Setenv("BUILDER_TOKEN", "3")
	status := cmd.Run([]string{"shutdown"})
	if status != 0 {
		t.Fatalf("Non-zero exit status %d\n", status)
	}
}

func handlerFactory(code int, body string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(code)
		io.WriteString(w, body)
	}
}
