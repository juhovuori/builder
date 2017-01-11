package server

import (
	"github.com/juhovuori/builder/app"
	"github.com/labstack/echo"
)

// Server is the HTTP server that implements Builder API
type Server interface {
	Run() error
}

type echoServer struct {
	echo *echo.Echo
	app  app.App
}

func (s echoServer) Run() error {
	return s.echo.Start(s.app.Config().ServerAddress())
}

func (s echoServer) setupRouteHandlers() error {
	s.echo.GET("/", s.hRoot)
	s.echo.GET("/health", s.hHealth)
	s.echo.POST("/v1/builds", s.hCreateBuild)
	s.echo.GET("/v1/builds", s.hListBuilds)
	s.echo.GET("/v1/builds/:id", s.hGetBuild)
	s.echo.POST("/v1/builds/:id", s.hAddStage)
	s.echo.GET("/v1/projects", s.hListProjects)
	s.echo.GET("/v1/projects/:id", s.hGetProject)
	s.echo.POST("/v1/projects/:id/trigger", s.hTriggerBuild)
	s.echo.GET("/v1/version", s.hVersion)
	return nil
}

// New creates a new echo Server instance
func New(app app.App) (Server, error) {
	echo := echo.New()
	server := echoServer{
		echo,
		app,
	}
	server.setupRouteHandlers()
	return server, nil
}
