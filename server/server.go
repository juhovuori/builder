package server

import (
	"net/http"

	"github.com/juhovuori/builder/app"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

func (s echoServer) errorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := http.StatusText(code)
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	}
	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD { // Issue #608
			c.NoContent(code)
		} else {
			err := serverError{code, msg}
			c.JSON(code, err)
		}
	}
	s.echo.Logger.Error(err)
}

// New creates a new echo Server instance
func New(app app.App) (Server, error) {
	server := echoServer{
		echo: echo.New(),
		app:  app,
	}
	server.echo.HTTPErrorHandler = server.errorHandler
	server.echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	server.setupRouteHandlers()
	setupVersion()
	return server, nil
}

type serverError struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
