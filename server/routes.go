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
	str := ""
	projects := s.app.Projects()
	for _, p := range projects {
		if len(str) != 0 {
			str += ", "
		}
		str += p.MD5() + "/" + p.URL() + "/" + p.Name()
		if p.Err() != nil {
			str += "/" + p.Err().Error()
		}
	}
	return c.String(http.StatusOK, str)
}

func (s *echoServer) hGetProject(c echo.Context) error {
	return c.String(http.StatusInternalServerError, "Not implemented.")
}
