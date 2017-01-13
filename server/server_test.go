package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
)

func TestSlash(t *testing.T) {
	Server, err := New(nil)
	server := Server.(echoServer)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("GET", "/", strings.NewReader(""))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := server.echo.NewContext(req, rec)
	err = server.hRoot(c)
	if err != nil {
		t.Fatal(err)
	}
	if rec.Code != 200 {
		t.Error("Invalid code", 200)
	}
	if rec.Body.String() != "\"Hello, World!\"" {
		t.Error("Invalid body", rec.Body.String())
	}
}
