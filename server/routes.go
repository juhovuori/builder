package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func (s *echoServer) hRoot(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func (s *echoServer) hHealth(c echo.Context) error {
	return c.String(http.StatusInternalServerError, "Not implemented.")
}

func (s *echoServer) hCreateBuild(c echo.Context) error {
	return c.String(http.StatusInternalServerError, "Not implemented.")
}

func (s *echoServer) hListBuilds(c echo.Context) error {
	return c.String(http.StatusInternalServerError, "Not implemented.")
}

func (s *echoServer) hGetBuild(c echo.Context) error {
	return c.String(http.StatusInternalServerError, "Not implemented.")
}

func (s *echoServer) hRouteStage(c echo.Context) error {
	return c.String(http.StatusInternalServerError, "Not implemented.")
}

func (s *echoServer) hListProjects(c echo.Context) error {
	return c.String(http.StatusInternalServerError, "Not implemented.")
}

func (s *echoServer) hGetProject(c echo.Context) error {
	return c.String(http.StatusInternalServerError, "Not implemented.")
}