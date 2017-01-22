package server

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/juhovuori/builder/build"
	"github.com/labstack/echo"
)

func (s *echoServer) hRoot(c echo.Context) error {
	return c.JSON(http.StatusOK, root)
}

func (s *echoServer) hHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

func (s *echoServer) hVersion(c echo.Context) error {
	v := s.app.Version()
	return c.JSON(http.StatusOK, v)
}

func (s *echoServer) hTriggerBuild(c echo.Context) error {
	projectID := c.Param("id")
	b, err := s.app.TriggerBuild(projectID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, restfulBuild(b))
}

func (s *echoServer) hCreateBuild(c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, "Not implemented.")
}

func (s *echoServer) hListBuilds(c echo.Context) error {
	builds := s.app.Builds(nil)
	index, page, err := paging(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, restfulItems(builds, "/v1/builds", index, page))
}

func (s *echoServer) hListProjectBuilds(c echo.Context) error {
	projectID := c.Param("id")
	builds := s.app.Builds(&projectID)
	index, page, err := paging(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, restfulItems(builds, "/v1/builds", index, page))
}

func (s *echoServer) hGetBuild(c echo.Context) error {
	buildID := c.Param("id")
	b, err := s.app.Build(buildID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, restfulBuild(b))
}

func (s *echoServer) hGetStageData(c echo.Context) error {
	buildID := c.Param("id")
	stage := c.Param("stage")
	b, err := s.app.StageData(buildID, stage)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, string(b))
}

func (s *echoServer) hGetStdout(c echo.Context) error {
	buildID := c.Param("id")
	stdout, err := s.app.Stdout(buildID)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, string(stdout))
}

func (s *echoServer) hAddStage(c echo.Context) error {
	buildID := c.Param("id")
	data, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	stage := build.Stage{
		Type: build.StageType(c.QueryParam("type")),
		Name: c.QueryParam("name"),
		Data: data,
	}
	if err = s.app.AddStage(buildID, stage); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "OK")
}

func (s *echoServer) hListProjects(c echo.Context) error {
	projects := s.app.Projects()
	index, page, err := paging(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, restfulItems(projects, "/v1/projects", index, page))
}

func (s *echoServer) hGetProject(c echo.Context) error {
	projectID := c.Param("id")
	p, err := s.app.Project(projectID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, restfulProject(p))
}

func (s *echoServer) hShutdown(c echo.Context) error {
	if s.token == nil {
		return c.JSON(http.StatusForbidden, "Shutdown disabled")
	}
	token, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	if string(token) != *s.token {
		return c.JSON(http.StatusForbidden, "Invalid token")
	}
	ch, err := s.app.Shutdown()
	if err != nil {
		return err
	}
	go func() {
		<-ch
		os.Exit(0)
	}()

	return c.JSON(http.StatusOK, "OK")
}
