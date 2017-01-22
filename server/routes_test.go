package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
)

func TestRoutes(t *testing.T) {
	Server, err := New(nil)
	server := Server.(echoServer)

	cases := []struct {
		handler func(c echo.Context) error
		code    int
		body    string
	}{
		{server.hRoot, 200, "\"Hello, World!\""},
		{server.hHealth, 500, "\"Not implemented.\""},
		//	{server.hVersion, 200, "\"xxx\""},
	}

	for i, c := range cases {
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest("GET", "", strings.NewReader(""))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := server.echo.NewContext(req, rec)
		err = c.handler(ctx)
		if err != nil {
			t.Fatal(err)
		}
		if rec.Code != c.code {
			t.Errorf("%d: Got code %d, expected %d", i, rec.Code, c.code)
		}
		if rec.Body.String() != c.body {
			t.Errorf("%d: Got body %s, expected %s", i, rec.Body.String(), c.body)
		}
	}
}
