package client

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShutdown(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handlerFactory(200, "OK")))

	buildID := "1"
	client, err := New(server.URL, buildID)
	if err != nil {
		t.Error(err)
	}
	client.Shutdown("t")
}

func handlerFactory(code int, body string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(code)
		io.WriteString(w, body)
	}
}
