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

func (s *echoServer) hVersion(c echo.Context) error {
	v := version()
	return c.JSON(http.StatusOK, v)
}

func (s *echoServer) hTriggerBuild(c echo.Context) error {
	projectID := c.Param("id")
	b, err := s.app.TriggerBuild(projectID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, b)
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

func (s *echoServer) hAddStage(c echo.Context) error {
	return c.String(http.StatusInternalServerError, "Not implemented.")
}

func (s *echoServer) hListProjects(c echo.Context) error {
	projects := s.app.Projects()
	return c.JSON(http.StatusOK, projects)
}

func (s *echoServer) hGetProject(c echo.Context) error {
	projectID := c.Param("id")
	p, err := s.app.Project(projectID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, p)
}
