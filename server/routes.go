package server

import (
	"net/http"

	"github.com/juhovuori/builder/build"
	"github.com/juhovuori/builder/version"
	"github.com/labstack/echo"
)

func (s *echoServer) hRoot(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}

func (s *echoServer) hHealth(c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, "Not implemented.")
}

func (s *echoServer) hVersion(c echo.Context) error {
	v := version.Version()
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
	return c.JSON(http.StatusInternalServerError, "Not implemented.")
}

func (s *echoServer) hListBuilds(c echo.Context) error {
	builds := s.app.Builds()
	return c.JSON(http.StatusOK, builds)
}

func (s *echoServer) hGetBuild(c echo.Context) error {
	buildID := c.Param("id")
	b, err := s.app.Build(buildID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, b)
}

func (s *echoServer) hGetStdout(c echo.Context) error {
	buildID := c.Param("id")
	b, err := s.app.Build(buildID)
	if err != nil {
		return err
	}
	stdout := string(b.Stdout())
	return c.String(http.StatusOK, stdout)
}

func (s *echoServer) hAddStage(c echo.Context) error {
	buildID := c.Param("id")
	stage := build.Stage{
		Type: build.StageType(c.QueryParam("type")),
		Name: c.QueryParam("name"),
	}
	err := s.app.AddStage(buildID, stage)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "OK")
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
